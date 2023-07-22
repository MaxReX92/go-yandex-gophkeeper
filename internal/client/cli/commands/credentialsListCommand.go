package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

const (
	credentialsListCommandName      = "list"
	credentialsListShortDescription = "list of all credentials"
	credentialsListFullDescription  = "Command list all stored credentials,"
)

type credentialsListCommand struct {
	*baseCommand
}

func NewCredentialsListCommand(stream io.CommandStream, children ...cli.Command) *credentialsListCommand {
	command := &credentialsListCommand{}
	command.baseCommand = newBaseCommand(
		stream,
		credentialsListCommandName,
		credentialsListShortDescription,
		credentialsListFullDescription,
		children,
		[]cli.Argument{
			newArgument("Reveal secret values", false, revealFullArgName, revealShortArgName),
			newArgument("Show full secret info", false, verboseFullArgName, verboseShortArgName),
		},
		command.invoke,
	)
	return command
}

func (c *credentialsListCommand) invoke(args map[string]string) error {
	_, reveal := argValue(args, revealFullArgName, revealShortArgName)

	// stub
	for i := 0; i < 10; i++ {
		value := "***"
		if reveal {
			value = fmt.Sprintf("secret value %v", i)
		}

		c.stream.Write(fmt.Sprintf("\tsecret name %v\t\t%s\n", i, value))
	}

	// storage.GetSecretsByType()
	// c.stream.Write(...)

	return nil
}
