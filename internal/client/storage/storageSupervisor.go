package storage

import (
	"context"
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/service"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type storageSupervisor struct {
	service       service.SecretService
	memoryStorage ClientSecretsStorage
}

func NewStorageSupervisor(service service.SecretService, memoryStorage ClientSecretsStorage) *storageSupervisor {
	return &storageSupervisor{
		service:       service,
		memoryStorage: memoryStorage,
	}
}

func (s *storageSupervisor) Start(ctx context.Context) error {
	eventStream := s.service.SecretEvents(ctx)

	for {
		select {
		case <-ctx.Done():
			logger.Info("supervisor stopping...")
			return ctx.Err()
		case event := <-eventStream:
			err := s.handleEvent(ctx, event)
			if err != nil {
				logger.ErrorFormat("failed to handler event: %v", err)
			}
		}
	}
}

func (s *storageSupervisor) handleEvent(ctx context.Context, event *model.SecretEvent) error {
	logger.InfoFormat("event received: %v", event)

	eventType := event.Type
	switch eventType {
	case model.Initial, model.Add:
		return s.memoryStorage.AddSecret(ctx, event.Secret)
	case model.Edit:
		return s.memoryStorage.ChangeSecret(ctx, event.Secret)
	case model.Remove:
		return s.memoryStorage.RemoveSecret(ctx, event.Secret)
	default:
		return logger.WrapError(fmt.Sprintf("handle event with type %v", eventType), model.ErrUnknownType)
	}
}
