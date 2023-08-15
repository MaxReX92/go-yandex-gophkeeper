package io

import "github.com/stretchr/testify/mock"

type CommandStreamMock struct {
	mock.Mock
}

func (c *CommandStreamMock) Read() <-chan string {
	args := c.Called()
	return args.Get(0).(<-chan string)
}

func (c *CommandStreamMock) Write(s string) {
	c.Called(s)
}
