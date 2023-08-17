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
)

const (
	binaryRemoveCommandName      = "remove"
	binaryRemoveShortDescription = "remove binary from store"
	binaryRemoveFullDescription  = "Remove binary from secure store,"
)

type binaryRemoveCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

// NewBinaryRemoveCommand creates a new instance of remove binary secret command.
func NewBinaryRemoveCommand(
	stream io.CommandStream,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *binaryRemoveCommand {
	command := &binaryRemoveCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		binaryRemoveCommandName,
		binaryRemoveShortDescription,
		binaryRemoveFullDescription,
		children,
		[]cli.Argument{
			newArgument("Remove all binary secrets", false, allFullArgName),
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *binaryRemoveCommand) invoke(ctx context.Context, args map[string]string) error {
	var toRemove []*secret.BinarySecret
	_, removeAll := argValue(args, allFullArgName)
	if removeAll {
		binaries, err := c.storage.GetAllSecrets(ctx, model.Binary)
		if err != nil {
			return logger.WrapError("get all secrets", err)
		}

		for _, binary := range binaries {
			toRemove = append(toRemove, binary.(*secret.BinarySecret))
		}
	} else {
		identity, ok := argValue(args, idShortArgName, idFullArgName)
		if !ok {
			return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
		}

		binary, err := c.storage.GetSecretByID(ctx, model.Binary, identity)
		if err != nil {
			return logger.WrapError("get secret", err)
		}

		toRemove = append(toRemove, binary.(*secret.BinarySecret))
	}

	for _, binary := range toRemove {
		logger.InfoFormat("Remove %s binary", binary.GetIdentity())
		err := c.storage.RemoveSecret(ctx, binary)
		if err != nil {
			return logger.WrapError("remove secret", err)
		}
	}

	return nil
}
