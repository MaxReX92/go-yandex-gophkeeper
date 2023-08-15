package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	initialCommandName      = "gophkeeper"
	initialShortDescription = "gophkeeper is a secure secrets storage"
	initialFullDescription  = "" +
		"gophkeeper is a client to high availability secure secrets storage system, \n" +
		"that provides multiple concurrent connections from different devices."
)

type initialCommand struct {
	*baseCommand
}

// NewInitialCommand creates a new instance of main secret command.
func NewInitialCommand(stream io.CommandStream, children ...cli.Command) *initialCommand {
	command := &initialCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		initialCommandName,
		initialShortDescription,
		initialFullDescription,
		children,
		[]cli.Argument{
			newArgument("Show information about command line tool version", false, versionFullArgName),
		},
		command.invoke,
	)
	return command
}

func (c *initialCommand) invoke(context.Context, map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command. See '%s help'.\n", c.FullName()))
	return nil
}
