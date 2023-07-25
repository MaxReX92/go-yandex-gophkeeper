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
	cardListCommandName      = "list"
	cardListShortDescription = "list of all cards"
	cardListFullDescription  = "Command list all stored cards,"
)

type cardListCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCardListCommand(stream io.CommandStream, storage storage.LocalSecretsStorage, children ...cli.Command) *cardListCommand {
	command := &cardListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		cardListCommandName,
		cardListShortDescription,
		cardListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
			newArgument("Show full secret info", false, verboseFullArgName, verboseShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *cardListCommand) invoke(args map[string]string) error {
	_, reveal := argValue(args, revealFullArgName, revealShortArgName)

	cards, err := c.storage.GetAllSecrets(model.Card)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	for _, modelCard := range cards {
		card := modelCard.(*secret.CardSecret)
		value := "***"
		if reveal {
			value = parser.Int32ToString(card.CVV)
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s\t\t%s", card.GetIdentity(), card.Number, card.Valid, value, card.GetComment()))
	}

	return nil
}
