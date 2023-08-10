package commands

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	clientIO "github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	binaryAddCommandName                  = "add"
	binaryAddShortDescription             = "add binary to store"
	binaryAddFullDescription              = "Add new binary to secure store,"
	fileMode                  os.FileMode = 0o644
)

type binaryAddCommand struct {
	*baseCommand
	generator identity.Generator
	storage   storage.ClientSecretsStorage
}

func NewBinaryAddCommand(
	stream clientIO.CommandStream,
	generator identity.Generator,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *binaryAddCommand {
	command := &binaryAddCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		binaryAddCommandName,
		binaryAddShortDescription,
		binaryAddFullDescription,
		children,
		[]cli.Argument{
			newArgument("Name", true, nameFullArgName, nameShortArgName),
			newArgument("Base64", true, base64FullArgName, base64ShortArgName),
			newArgument("File path", true, filePathFullArgName, filePathShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *binaryAddCommand) invoke(ctx context.Context, args map[string]string) error {
	name, ok := argValue(args, nameFullArgName, nameShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: name is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	var reader io.Reader
	filePath, ok := argValue(args, filePathFullArgName, filePathShortArgName)
	if ok {
		if _, err := os.Stat(filePath); err != nil && errors.Is(err, os.ErrNotExist) {
			return logger.WrapError(fmt.Sprintf("open %s file", filePath), cli.ErrFileNotFound)
		}

		fileStream, err := os.OpenFile(filePath, os.O_RDONLY, fileMode)
		if err != nil {
			return logger.WrapError("open file", err)
		}
		reader = fileStream
	}

	base64String, ok := argValue(args, base64FullArgName, base64ShortArgName)
	if ok {
		if reader != nil {
			return logger.WrapError("add binary: either file path or base64 string is expected", cli.ErrInvalidArguments)
		}

		reader = base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(base64String))
	}

	if reader == nil {
		return logger.WrapError("add binary: either file path or base64 string is expected", cli.ErrInvalidArguments)
	}

	comment, _ := argValue(args, commentFullArgName, commentShortArgName)

	binary := secret.NewBinarySecret(name, reader, c.generator.GenerateNewIdentity(), comment)
	logger.Info("Add binary")
	err := c.storage.AddSecret(ctx, binary)
	if err != nil {
		return logger.WrapError("add secret", err)
	}

	return nil
}
