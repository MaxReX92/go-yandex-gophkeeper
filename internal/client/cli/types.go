package cli

type Command interface {
	Name() string
	FullName() string
	ShortDescription() string
	SetParent(command Command)
	Invoke(keys []string) error
	ShowHelp()
}

type Argument interface {
	Keys() []string
	Description() string
	NextArgIsValue() bool
}
