package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
)

type helpCommand struct {
	name   string
	parent cli.Command
}

func NewHelpCommand() *helpCommand {
	return &helpCommand{
		name: "help",
	}
}

func (c *helpCommand) Name() string {
	return c.name
}

func (c *helpCommand) FullName() string {
	return fmt.Sprintf("%s %s", c.parent.FullName(), c.name)
}

func (c *helpCommand) ShortDescription() string {
	return "View help information about a command"
}

func (c *helpCommand) SetParent(command cli.Command) {
	c.parent = command
}

func (c *helpCommand) Invoke(keys []string) error {
	c.parent.ShowHelp()
	return nil
}

func (c *helpCommand) ShowHelp() {
	c.parent.ShowHelp()
}
