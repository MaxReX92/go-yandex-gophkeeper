package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	noteAddCommandName      = "add"
	noteAddShortDescription = "add note to store"
	noteAddFullDescription  = "Add new note to secure store,"
)

type noteAddCommand struct {
	*baseCommand
	generator generator.Generator
	storage   storage.ClientSecretsStorage
}

func NewNoteAddCommand(
	stream io.CommandStream,
	generator generator.Generator,
	storage storage.ClientSecretsStorage,
	children ...cli.Command,
) *noteAddCommand {
	command := &noteAddCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		noteAddCommandName,
		noteAddShortDescription,
		noteAddFullDescription,
		children,
		[]cli.Argument{
			newArgument("Note text", true, textFullArgName, textShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *noteAddCommand) invoke(ctx context.Context, args map[string]string) error {
	text, ok := argValue(args, textFullArgName, textShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: note texttext is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	comment, _ := argValue(args, commentFullArgName, commentShortArgName)

	note := secret.NewNoteSecret(text, c.generator.GenerateNewIdentity(), comment)
	logger.InfoFormat("Add %s note", note.Text)
	err := c.storage.AddSecret(ctx, note)
	if err != nil {
		return logger.WrapError("add secret", err)
	}

	return nil
}
