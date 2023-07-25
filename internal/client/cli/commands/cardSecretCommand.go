package commands

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
)

type cardCommand struct {
}

func NewCardsCommand(childCommands ...cli.Command) *cardCommand {
	return &cardCommand{}
}
