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
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	binaryEditCommandName      = "edit"
	binaryEditShortDescription = "edit binary from store"
	binaryEditFullDescription  = "Edit new binary from secure store,"
)

type binaryEditCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

// NewBinaryEditCommand creates a new instance of edit binary secret command.
func NewBinaryEditCommand(
	stream clientIO.CommandStream,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *binaryEditCommand {
	command := &binaryEditCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		binaryEditCommandName,
		binaryEditShortDescription,
		binaryEditFullDescription,
		children,
		[]cli.Argument{
			newArgument("Secret identity", true, idFullArgName, idShortArgName),
			newArgument("Binary text", true, textFullArgName, textShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *binaryEditCommand) invoke(ctx context.Context, args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	currentBinary, err := c.storage.GetSecretByID(ctx, model.Binary, identity)
	if err != nil {
		return logger.WrapError("get secret", err)
	}

	if currentBinary == nil {
		return logger.WrapError("edit secret", cli.ErrSecretNotFound)
	}

	binary, ok := currentBinary.(*secret.BinarySecret)
	if !ok {
		return logger.WrapError("edit secret", cli.ErrInvalidSecretType)
	}

	name, ok := argValue(args, nameFullArgName, nameShortArgName)
	if ok {
		binary.Name = name
	}

	var reader io.Reader
	filePath, ok := argValue(args, filePathFullArgName, filePathShortArgName)
	if ok {
		if _, err = os.Stat(filePath); err != nil && errors.Is(err, os.ErrNotExist) {
			return logger.WrapError(fmt.Sprintf("open %s file", filePath), cli.ErrFileNotFound)
		}

		reader, err = os.OpenFile(filePath, os.O_RDONLY, fileMode)
		if err != nil {
			return logger.WrapError("open file", err)
		}
	}

	base64String, ok := argValue(args, base64FullArgName, base64ShortArgName)
	if ok {
		if reader != nil {
			return logger.WrapError("add binary: either file path or base64 string is expected", cli.ErrInvalidArguments)
		}

		reader = base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(base64String))
	}

	if reader != nil {
		binary.Reader = reader
	}

	comment, ok := argValue(args, commentFullArgName, commentShortArgName)
	if ok {
		binary.Comment = comment
	}

	logger.Info("Edit binary")
	err = c.storage.ChangeSecret(ctx, binary)
	if err != nil {
		return logger.WrapError("edit secret", err)
	}

	return nil
}
