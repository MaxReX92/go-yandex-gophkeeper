package cli

import (
	"context"
	"errors"
	"strings"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/auth"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type handler struct {
	ioStream       io.CommandStream
	initialCommand Command
}

// NewHandler creates a new instance of cli command handler.
func NewHandler(ioStream io.CommandStream, initialCommand Command) *handler {
	return &handler{
		ioStream:       ioStream,
		initialCommand: initialCommand,
	}
}

func (h *handler) Start(ctx context.Context) error {
	readStream := h.ioStream.Read()
	for {
		select {
		case message := <-readStream:
			if message == "" {
				continue
			}

			err := h.initialCommand.Invoke(ctx, strings.Split(message, " "))
			if err != nil {
				if errors.Is(err, auth.ErrUnauthorized) {
					h.ioStream.Write("You`ll need to be authorized first\n")
				} else {
					h.ioStream.Write(logger.WrapError("invoke command", err).Error())
				}
			}
		case <-ctx.Done():
			logger.Info("handler stopping...")
			return ctx.Err()
		}
	}
}
