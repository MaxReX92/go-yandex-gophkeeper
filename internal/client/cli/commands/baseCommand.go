package commands

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
)

type baseCommand struct {
	stream           io.CommandStream
	name             string
	shortDescription string
	fullDescription  string
	parent           cli.Command
	children         map[string]cli.Command
	args             map[string]cli.Argument
}

func newBaseCommand(
	stream io.CommandStream,
	name string,
	shortDescription string,
	fullDescription string,
	parent cli.Command,
	args []cli.Argument,
	invokeMethod func() error,
) *baseCommand {
	base := &baseCommand{
		stream:           stream,
		name:             name,
		shortDescription: shortDescription,
		fullDescription:  fullDescription,
		children:         make(map[string]cli.Command),
		args:             make(map[string]cli.Argument, len(args)),
	}
	for _, arg := range args {
		base.args[arg.ShortName()] = arg
		base.args[arg.FullName()] = arg
	}

	if parent != nil {
		parent.AddChild(base)
	}

	return base
}

func (c *baseCommand) Name() string {
	return c.name
}

func (c *baseCommand) FullName() string {
	if c.parent == nil {
		return c.name
	}

	return fmt.Sprintf("%s %s", c.parent.FullName(), c.name)
}

func (c *baseCommand) ShortDescription() string {
	return c.shortDescription
}

func (c *baseCommand) AddChild(command cli.Command) {
	c.children[command.Name()] = command
}

func (c *baseCommand) Invoke(keys []string) error {
	c.showHelp()
	return nil
}

func (c *baseCommand) showHelp() {
	c.stream.Write(c.fullDescription)
	c.stream.Write("\n\n")
	c.stream.Write("Usage:\n")
	c.stream.Write(fmt.Sprintf("\t%s", c.FullName()))
	if len(c.children) > 0 {
		c.stream.Write(fmt.Sprintf(" <command>"))
	}
	if len(c.args) > 0 {
		c.stream.Write(fmt.Sprintf(" [arguments]"))
	}
	c.stream.Write("\n")

	if len(c.children) > 0 {
		c.stream.Write("The commands are:\n")
		for _, command := range c.children {
			c.stream.Write(fmt.Sprintf("\t%s\t\t%s\n", command.Name(), command.ShortDescription()))
		}
	}

	if len(c.args) > 0 {
		c.stream.Write("The arguments are:\n")
		for _, arg := range c.args {
			c.stream.Write(fmt.Sprintf("\t[%s | %s]\t\t%s\n", arg.FullName(), arg.ShortName(), arg.Description()))
		}
	}
}
