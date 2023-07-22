package commands

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
)

type cardsCommand struct {
}

func NewCardsCommand(childCommands ...cli.Command) *cardsCommand {
	return &cardsCommand{}
}
