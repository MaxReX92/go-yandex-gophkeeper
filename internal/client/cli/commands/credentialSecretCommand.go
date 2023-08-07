package commands

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	credentialCommandName      = "cred"
	credentialShortDescription = "work with user credential secrets"
	credentialFullDescription  = "" +
		"System provide functionality for storing various user credential data,\n" +
		"such as login, password, token, etc.\n" +
		"'Cred' command helps to work with stored credential list."
)

type credentialCommand struct {
	*baseCommand
}

func NewCredentialCommand(stream io.CommandStream, children ...cli.Command) *credentialCommand {
	command := &credentialCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		credentialCommandName,
		credentialShortDescription,
		credentialFullDescription,
		children,
		nil,
		command.invoke,
	)
	return command
}

func (c *credentialCommand) invoke(context.Context, map[string]string) error {
	c.stream.Write(fmt.Sprintf("Unexpected command arguments. See '%s help'.\n", c.FullName()))
	return nil
}
