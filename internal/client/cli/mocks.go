package cli

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

func (c *CommandMock) Name() string {
	args := c.Called()
	return args.String(0)
}

func (c *CommandMock) FullName() string {
	args := c.Called()
	return args.String(0)
}

func (c *CommandMock) ShortDescription() string {
	args := c.Called()
	return args.String(0)
}

func (c *CommandMock) SetParent(command Command) {
	c.Called(command)
}

func (c *CommandMock) Invoke(ctx context.Context, keys []string) error {
	args := c.Called(ctx, keys)
	return args.Error(0)
}

func (c *CommandMock) ShowHelp() {
	c.Called()
}
