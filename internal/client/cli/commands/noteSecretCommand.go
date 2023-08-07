package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	noteCommandName      = "note"
	noteShortDescription = "work with note secrets"
	noteFullDescription  = "" +
		"System provide functionality for storing various note data.\n" +
		"'Note' command helps to work with stored notes list."
)

type noteCommand struct {
	*baseCommand
}

func NewNoteCommand(stream io.CommandStream, children ...cli.Command) *noteCommand {
	command := &noteCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		noteCommandName,
		noteShortDescription,
		noteFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *noteCommand) invoke(context.Context, map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command arguments. See '%s help'.\n", c.FullName()))
	return nil
}
