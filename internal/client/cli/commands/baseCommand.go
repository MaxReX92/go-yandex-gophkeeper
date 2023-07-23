package commands

import (
	"fmt"
	"sort"
	"strings"

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
	invokeMethod     func(map[string]string) error
}

func newBaseCommand(
	stream io.CommandStream,
	name string,
	shortDescription string,
	fullDescription string,
	children []cli.Command,
	args []cli.Argument,
	invokeMethod func(map[string]string) error,
) *baseCommand {
	base := &baseCommand{
		stream:           stream,
		name:             name,
		shortDescription: shortDescription,
		fullDescription:  fullDescription,
		children:         make(map[string]cli.Command, len(children)),
		args:             make(map[string]cli.Argument, len(args)),
		invokeMethod:     invokeMethod,
	}

	if children != nil {
		for _, child := range children {
			child.SetParent(base)
			base.children[child.Name()] = child
		}
	}

	if args != nil {
		for _, arg := range args {
			for _, key := range arg.Keys() {
				base.args[key] = arg
			}
		}
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

func (c *baseCommand) SetParent(command cli.Command) {
	c.parent = command
}

func (c *baseCommand) Invoke(keys []string) error {
	args := make(map[string]string)
	keysLenhth := len(keys)
	if keysLenhth == 0 {
		return c.invokeMethod(args)
	}

	if childCommand, ok := c.children[keys[0]]; ok {
		return childCommand.Invoke(keys[1:])
	}

	for i := 0; i < keysLenhth; i++ {
		key := keys[i]
		arg, ok := c.args[key]
		if !ok {
			c.stream.Write(fmt.Sprintf("Unexpected argument: %s. See '%s help'.\n", key, c.FullName()))
			return nil
		}

		var value string
		if arg.NextArgIsValue() {
			i++
			if i > keysLenhth {
				c.stream.Write(fmt.Sprintf("Argument value missed: %s. See '%s help'.\n", key, c.FullName()))
				return nil
			}

			value = keys[i]
		}

		args[key] = value
	}

	return c.invokeMethod(args)
}

func (c *baseCommand) ShowHelp() {
	c.stream.Write(c.fullDescription)
	c.stream.Write("\n\n")
	c.stream.Write("Usage:\n")
	c.stream.Write(fmt.Sprintf("\t%s", c.FullName()))

	// TODO: ignore help command
	childrenLen := len(c.children)
	if len(c.children) > 0 {
		c.stream.Write(fmt.Sprintf(" <command>"))
	}

	argsLen := len(c.args)
	if len(c.args) > 0 {
		c.stream.Write(fmt.Sprintf(" [arguments]"))
	}
	c.stream.Write("\n\n")

	if childrenLen > 0 {
		c.stream.Write("The commands are:\n")
		i := 0
		commandNames := make([]string, childrenLen)
		for childName := range c.children {
			commandNames[i] = childName
			i++
		}
		sort.Strings(commandNames)

		for _, commandName := range commandNames {
			command := c.children[commandName]
			c.stream.Write(fmt.Sprintf("\t%s\t\t%s\n", command.Name(), command.ShortDescription()))
		}

		c.stream.Write("\n")
	}

	if argsLen > 0 {
		c.stream.Write("The arguments are:\n")

		i := 0
		argNames := make([]string, argsLen)
		for argName := range c.args {
			argNames[i] = argName
			i++
		}

		argsMap := make(map[cli.Argument]interface{})

		for _, arg := range c.args {
			keys := strings.Join(arg.Keys(), ",")
			_, ok := argsMap[arg]
			if !ok {
				c.stream.Write(fmt.Sprintf("\t%s\t%s\n", keys, arg.Description()))
				argsMap[arg] = nil
			}
		}
	}
}
