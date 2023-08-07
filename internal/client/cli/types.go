package cli

import "context"

type Command interface {
	Name() string
	FullName() string
	ShortDescription() string
	SetParent(command Command)
	Invoke(ctx context.Context, keys []string) error
	ShowHelp()
}

type Argument interface {
	Keys() []string
	Description() string
	NextArgIsValue() bool
}
