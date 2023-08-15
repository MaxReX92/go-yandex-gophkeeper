package cli

import "context"

// Command represent client cli command.
type Command interface {
	// Name returns command name.
	Name() string
	// FullName return command full name in command tree hierarchy.
	FullName() string
	// ShortDescription return command brief description.
	ShortDescription() string
	// SetParent set command parent in command tree hierarchy.
	SetParent(command Command)
	// Invoke run command.
	Invoke(ctx context.Context, keys []string) error
	// ShowHelp present command summary help information.
	ShowHelp()
}

// Argument represent client command argument.
type Argument interface {
	// Keys returns list of key string that points to this argument.
	Keys() []string
	// Description is an argument full text specification.
	Description() string
	// NextArgIsValue returns true, if next arg in arguments list is a value of this argument.
	NextArgIsValue() bool
}
