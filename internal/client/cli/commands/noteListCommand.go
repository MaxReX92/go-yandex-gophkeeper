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
	noteListCommandName      = "list"
	noteListShortDescription = "list of all note"
	noteListFullDescription  = "Command list all stored note."
)

type noteListCommand struct {
	*baseCommand
	storage storage.ClientSecretsStorage
}

// NewNoteListCommand creates a new instance of list note secret command.
func NewNoteListCommand(stream io.CommandStream, storage storage.ClientSecretsStorage, children ...cli.Command) *noteListCommand {
	command := &noteListCommand{
		storage: storage,
	}
	command.baseCommand = newBaseCommand(
		stream,
		noteListCommandName,
		noteListShortDescription,
		noteListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *noteListCommand) invoke(ctx context.Context, args map[string]string) error {
	notes, err := c.storage.GetAllSecrets(ctx, model.Note)
	if err != nil {
		return logger.WrapError("get secrets", err)
	}

	_, reveal := argValue(args, revealFullArgName, revealShortArgName)
	for _, modelNote := range notes {
		note := modelNote.(*secret.NoteSecret)
		value := hiddenValue
		if reveal {
			value = note.Text
		}

		c.stream.Write(fmt.Sprintf("\t%s\t\t%s\t\t%s\n", note.GetIdentity(), value, note.GetComment()))
	}

	return nil
}
