package commands

import (
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

func NewInitialCommand(stream io.CommandStream) *initialCommand {
	command := &initialCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		initialCommandName,
		initialShortDescription,
		initialFullDescription,
		nil,
		[]cli.Argument{},
		command.invoke,
	)
	return command
}

func (c *initialCommand) invoke() error {

	return nil
}
