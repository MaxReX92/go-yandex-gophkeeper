package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/generator"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

const (
	notesAddCommandName      = "add"
	notesAddShortDescription = "add notes to store"
	notesAddFullDescription  = "Add new notes to secure store,"
)

type notesAddCommand struct {
	*baseCommand
	generator generator.Generator
	storage   storage.LocalSecretsStorage
}

func NewNotesAddCommand(
	stream io.CommandStream,
	generator generator.Generator,
	storage storage.LocalSecretsStorage,
	children ...cli.Command,
) *notesAddCommand {
	command := &notesAddCommand{
		generator: generator,
		storage:   storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		notesAddCommandName,
		notesAddShortDescription,
		notesAddFullDescription,
		children,
		[]cli.Argument{
			newArgument("Note text", true, textFullArgName, textShortArgName),
			newArgument("Comment", true, commentFullArgName, commentShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *notesAddCommand) invoke(args map[string]string) error {
	text, ok := argValue(args, textFullArgName, textShortArgName)
	if !ok {
		return logger.WrapError(fmt.Sprintf("invoke %s command: note text is missed", c.name), cli.ErrRequiredArgNotFound)
	}

	comment, _ := argValue(args, commentFullArgName, commentShortArgName)

	notes := secret.NewNotesSecret(text, c.generator.GenerateNewIdentity(), comment)
	logger.InfoFormat("Add %s notes", notes.Text)
	err := c.storage.AddSecret(notes)
	if err != nil {
		return logger.WrapError("add secret", err)
	}

	return nil
}
