package commands

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	statusCommandName      = "status"
	statucShortDescription = "provide information about client status"
	statusFullDescription  = "'Status' command provide complex information about current gophkeeper client state."
)

type statusCommand struct {
	*baseCommand
}

func NewStatusCommand(stream io.CommandStream, children ...cli.Command) *statusCommand {
	command := &statusCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		statusCommandName,
		statucShortDescription,
		statusFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *statusCommand) invoke(map[string]string) error {
	c.stream.Write("OK")
	return nil
}
