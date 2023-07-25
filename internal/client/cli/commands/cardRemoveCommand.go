package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	cardRemoveCommandName      = "remove"
	cardRemoveShortDescription = "remove card from store"
	cardRemoveFullDescription  = "Remove card from secure store,"
)

type cardRemoveCommand struct {
	*baseCommand
	storage storage.LocalSecretsStorage
}

func NewCardRemoveCommand(
	stream io.CommandStream,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *cardRemoveCommand {
	command := &cardRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		cardRemoveCommandName,
		cardRemoveShortDescription,
		cardRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all card secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *cardRemoveCommand) invoke(args map[string]string) error {
	var toRemove []*secret.CardSecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		cards, err := c.storage.GetAllSecrets(model.Card)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, card := range cards {
			toRemove = append(toRemove, card.(*secret.CardSecret))
		}

	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		card, err := c.storage.GetSecretById(model.Card, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, card.(*secret.CardSecret))
	}

	for _, card := range toRemove {
		logger.InfoFormat("Remove %s %s card", card.GetIdentity(), card.Number)
		err := c.storage.RemoveSecret(card)
		if err != nil {
			return logger.WrapError("get secret", err)
		}
	}

	return nil
}