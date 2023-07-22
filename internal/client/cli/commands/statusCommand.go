package commands

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	statusCommandName      = "status"
	statucShortDescription = "provide information about client status"
	statusFullDescription  = "Status command provide complex information about current gophkeeper client state."
)

type statusCommand struct {
	*baseCommand
}

func NewStatusCommand(parent cli.Command, stream io.CommandStream) *statusCommand {
	command := &statusCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		statusCommandName,
		statucShortDescription,
		statusFullDescription,
		parent,
		[]cli.Argument{},
		command.invoke,
	)
	return command
}

func (c *statusCommand) invoke() error {

	return nil
}
