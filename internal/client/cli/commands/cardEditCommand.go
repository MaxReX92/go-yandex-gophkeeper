package commands

import (
	"errors"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
)

const (
	cardEditCommandName      = "edit"
	cardEditShortDescription = "edit card from store"
	cardEditFullDescription  = "Edit card from secure store,"
)

type cardEditCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewCardEditCommand(
	stream io.CommandStream,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *cardEditCommand {
	command := &cardEditCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		cardEditCommandName,
		cardEditShortDescription,
		cardEditFullDescription,
		children,
		[]cli.Argument{
			newArgument("Secret identity", true, initialFullDescription, initialShortDescription),
			newArgument("Card number", true, numFullArgName, numShortArgName),
			newArgument("CVV", true, cvvFullArgName),
			newArgument("Valid thru date (MM/YY)", true, validThruFullArgName, validThruShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *cardEditCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentCard, err := c.storage.GetSecretById(nil, model.Card, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentCard == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	card, ok := currentCard.(*secret.CardSecret)
	if !ok {
		return logger.WrapError("edit secret", cli.ErrInvalidSecretType)
	}

	number, ok := argValue(args, numFullArgName, numShortArgName)
	if ok {
		card.Number = number
	}

	validThru, ok := argValue(args, validThruFullArgName, validThruShortArgName)
	if ok {
		valid, err := parser.ToTime(validThru)
		if err != nil {
			if errors.Is(err, parser.ErrInvalidFormat) {
				c.stream.Write("Invalid valid thru value format, see help command for more details")
				return nil
			} else {
				return logger.WrapError("parse valid thru value", err)
			}
		}

		card.Valid = valid
	}

	cvvArg, ok := argValue(args, cvvFullArgName)
	if ok {
		cvv, err := parser.ToInt32(cvvArg)
		if err != nil {
			return logger.WrapError("parse cvv value", err)
		}
		card.CVV = cvv
	}

	comment, ok := argValue(args, commentFullArgName, commentShortArgName)
	if ok {
		card.Comment = comment
	}

	logger.InfoFormat("Edit %s card", number)
	err = c.storage.ChangeSecret(nil, card)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
