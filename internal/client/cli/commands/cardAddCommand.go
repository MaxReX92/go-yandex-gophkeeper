package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/identity"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
)

const (
	cardAddCommandName      = "add"
	cardAddShortDescription = "add card to store"
	cardAddFullDescription  = "Add new card to secure store,"
)

type cardAddCommand struct {
	*baseCommand
	generator identity.Generator
	storage   storage.ClientSecretsStorage
}

func NewCardAddCommand(
	stream io.CommandStream,
	generator identity.Generator,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *cardAddCommand {
	command := &cardAddCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		cardAddCommandName,
		cardAddShortDescription,
		cardAddFullDescription,
		children,
		[]cli.Argument{
			newArgument("Card number", true, numFullArgName, numShortArgName),
			newArgument("CVV", true, cvvFullArgName),
			newArgument("Valid thru date (MM/YY)", true, validThruFullArgName, validThruShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *cardAddCommand) invoke(ctx context.Context, args map[string]string) error {
	number, ok := argValue(args, numFullArgName, numShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: card number is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	validThru, ok := argValue(args, validThruFullArgName, validThruShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: valid thru date is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	valid, err := parser.ToTime(validThru)
	if err != nil {
		if errors.Is(err, parser.ErrInvalidFormat) {
			c.stream.Write("Invalid valid thru value format, see help command for more details")
			return nil
		} else {
			return logger.WrapError("parse valid thru value", err)
		}
	}

	cvvArg, ok := argValue(args, cvvFullArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: cvv value is missed", c.name), cli.ErrRequiredArgNotFound)
	}
	cvv, err := parser.ToInt32(cvvArg)
	if err != nil {
		return logger.WrapError("parse cvv value", err)
	}

	comment, _ := argValue(args, commentFullArgName, commentShortArgName)

	card := secret.NewCardSecret(number, cvv, valid, c.generator.GenerateNewIdentity(), comment)
	logger.InfoFormat("Add %s card", number)
	err = c.storage.AddSecret(ctx, card)
	if err != nil {
		return logger.WrapError("add secret", err)
	}

	return nil
}
