package commands

import (
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
	cardGetCommandName      = "get"
	cardGetShortDescription = "get card secret"
	cardGetFullDescription  = "Command get all stored cards,"
)

type cardGetCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewCardGetCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *cardGetCommand {
	command := &cardGetCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		cardGetCommandName,
		cardGetShortDescription,
		cardGetFullDescription,
		children,
		[]cli.Argument{
			newArgument("Identity", true, idFullArgName, idShortArgName),
			newArgument("Reveal secret value", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *cardGetCommand) invoke(args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	cardSecret, err := c.storage.GetSecretById(model.Card, identity)
	if err != nil {
		return logger.WrapError("get card secret", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	card := cardSecret.(*secret.CardSecret)
	value := "***"
	if reveal {
		value = parser.Int32ToString(card.CVV)
	}

	c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s\t\t%s", card.GetIdentity(), card.Number, card.Valid, value, card.GetComment()))

	return nil
}
