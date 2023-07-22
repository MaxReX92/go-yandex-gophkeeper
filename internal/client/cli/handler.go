package cli

import (
	"context"
	"strings"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type handler struct {
	ioStream       io.CommandStream
	initialCommand Command
}

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

			err := h.initialCommand.Invoke(strings.SplitN(message, " ", -1))
			if err != nil {
				return logger.WrapError("invoke command", err)
			}
		case <-ctx.Done():
			logger.Info("stopping...")
			return ctx.Err()
		}
	}
}
