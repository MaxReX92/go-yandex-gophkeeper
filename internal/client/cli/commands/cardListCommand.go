package commands

import (
	"context"
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
	storage storage.ClientSecretsStorage
}

func NewCardListCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *cardListCommand {
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
		},
		command.invoke,
	)
	return command
}

func (c *cardListCommand) invoke(ctx context.Context, args map[string]string) error {
	cards, err := c.storage.GetAllSecrets(ctx, model.Card)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	for _, modelCard := range cards {
		card := modelCard.(*secret.CardSecret)
		value := hiddenValue
		if reveal {
			value = parser.Int32ToString(card.CVV)
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s\t\t%s\n", card.GetIdentity(), card.Number, card.Valid, value, card.GetComment()))
	}

	return nil
}
