package cli

type Command interface {
	Name() string
	FullName() string
	ShortDescription() string
	AddChild(command Command)
	Invoke(keys []string) error
}

type Argument interface {
	FullName() string
	ShortName() string
	Description() string
	NextArgIsValue() bool
}
