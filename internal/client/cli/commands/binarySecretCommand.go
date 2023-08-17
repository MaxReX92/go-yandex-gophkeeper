package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	binaryCommandName      = "binary"
	binaryShortDescription = "work with binary secrets"
	binaryFullDescription  = "" +
		"System provide functionality for storing various binary data.\n" +
		"'Binary' command helps to work with stored binary secrets."
)

type binaryCommand struct {
	*baseCommand
}

// NewBinaryCommand creates a new instance of main binary secret command.
func NewBinaryCommand(stream io.CommandStream, children ...cli.Command) *binaryCommand {
	command := &binaryCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		binaryCommandName,
		binaryShortDescription,
		binaryFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *binaryCommand) invoke(context.Context, map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command arguments. See '%s help'.\n", c.FullName()))
	return nil
}
