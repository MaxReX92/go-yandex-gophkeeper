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
	binaryListCommandName      = "list"
	binaryListShortDescription = "list of all binaries"
	binaryListFullDescription  = "Command list all stored binaries,"
)

type binaryListCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewBinaryListCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *binaryListCommand {
	command := &binaryListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		binaryListCommandName,
		binaryListShortDescription,
		binaryListFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *binaryListCommand) invoke(args map[string]string) error {
	binaries, err := c.storage.GetAllSecrets(model.Binary)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	for _, modelBinary := range binaries {
		binary := modelBinary.(*secret.BinarySecret)
		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s\t\t%s", binary.GetIdentity(), binary.Name, binary.GetComment()))
	}

	return nil
}
