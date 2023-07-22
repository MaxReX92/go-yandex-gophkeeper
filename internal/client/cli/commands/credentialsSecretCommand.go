package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	credentialsCommandName      = "cred"
	credentialsShortDescription = "work with user credential secrets"
	credentialsFullDescription  = "" +
		"System provide functionality for storing various user credentials data,\n" +
		"such as login, password, token, etc.\n" +
		"'Cred' command helps to work with stored credentials list."
)

type credentialsCommand struct {
	*baseCommand
}

func NewCredentialsCommand(stream io.CommandStream, children ...cli.Command) *credentialsCommand {
	command := &credentialsCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsCommandName,
		credentialsShortDescription,
		credentialsFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *credentialsCommand) invoke(map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command. See '%s help'.\n", c.FullName()))
	return nil
}
