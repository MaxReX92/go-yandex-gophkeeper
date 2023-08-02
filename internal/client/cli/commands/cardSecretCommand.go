package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	cardCommandName      = "card"
	cardShortDescription = "work with card secrets"
	cardFullDescription  = "" +
		"System provide functionality for storing various card data,\n" +
		"such as card number, cvv code, valid thru, etc.\n" +
		"'Card' command helps to work with stored cards list."
)

type cardCommand struct {
	*baseCommand
}

func NewCardCommand(stream io.CommandStream, children ...cli.Command) *cardCommand {
	command := &cardCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		cardCommandName,
		cardShortDescription,
		cardFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *cardCommand) invoke(map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command arguments. See '%s help'.\n", c.FullName()))
	return nil
}
