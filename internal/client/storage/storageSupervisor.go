package storage

import (
	"context"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type storageSupervisor struct {
}

func (s *storageSupervisor) Start(ctx context.Context) error {
	//readStream := h.ioStream.Read()
	for {
		select {
		case <-ctx.Done():
			logger.Info("supervisor stopping...")
			return ctx.Err()
		}
	}
}
