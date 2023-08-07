package commands

import (
	"context"
	"encoding/base64"
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
	binaryGetCommandName      = "get"
	binaryGetShortDescription = "get of all binaries"
	binaryGetFullDescription  = "Command get all stored binaries,"
)

type binaryGetCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

func NewBinaryGetCommand(stream clientIO.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *binaryGetCommand {
	command := &binaryGetCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		binaryGetCommandName,
		binaryGetShortDescription,
		binaryGetFullDescription,
		children,
		[]cli.Argument{
			newArgument("Identity", true, idFullArgName, idShortArgName),
			newArgument("Output file path", true, filePathFullArgName, filePathShortArgName),
			newArgument("To bae64", false, base64FullArgName, base64ShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *binaryGetCommand) invoke(ctx context.Context, args map[string]string) error {
	identity, ok := argValue(args, idFullArgName, idShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: secret identity is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	binarySecret, err := c.storage.GetSecretById(ctx, model.Binary, identity)
	if err != nil {
		return logger.WrapError("get binary secret", err)
	}

	binary := binarySecret.(*secret.BinarySecret)

	value := ""
	_, toBase64 := argValue(args, base64FullArgName, base64ShortArgName)
	if toBase64 {
		bytes, err := io.ReadAll(binary.Reader)
		if err != nil {
			return logger.WrapError("read binary secret data", err)
		}

		value = base64.StdEncoding.EncodeToString(bytes)
	}

	filePath, ok := argValue(args, filePathFullArgName, filePathShortArgName)
	if ok {
		fileStream, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, fileMode)
		if err != nil {
			return logger.WrapError("open file", err)
		}
		defer func(fileStream *os.File) {
			err = fileStream.Close()
			if err != nil {
				logger.ErrorFormat("failed to close file: %v", err)
			}
		}(fileStream)

		if value != "" {
			_, err = fileStream.WriteString(value)
			if err != nil {
				return logger.WrapError("write secret data", err)
			}
		} else {
			_, err = io.Copy(fileStream, binary.Reader)
			if err != nil {
				return logger.WrapError("write secret data", err)
			}
			value = fileStream.Name()
		}
	}

	c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\t\t%s", binary.GetIdentity(), binary.Name, value, binary.GetComment()))

	return nil
}
