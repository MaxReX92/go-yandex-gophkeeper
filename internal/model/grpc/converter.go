package grpc

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/serialization"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type Converter struct {
	serializer serialization.Serializer
}

func NewConverter(serializer serialization.Serializer) *Converter {
	return &Converter{
		serializer: serializer,
	}
}

func (c *Converter) ToModelSecret(generatedSecret *generated.Secret) (model.Secret, error) {
	secretType := generatedSecret.Type
	switch secretType {
	case generated.SecretType_BINARY:
		return c.toModelSecret(generatedSecret, &secret.BinarySecret{})
	case generated.SecretType_CARD:
		return c.toModelSecret(generatedSecret, &secret.CardSecret{})
	case generated.SecretType_CREDENTIAL:
		return c.toModelSecret(generatedSecret, &secret.CredentialSecret{})
	case generated.SecretType_NOTE:
		return c.toModelSecret(generatedSecret, &secret.NoteSecret{})
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert secret with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) FromModelSecret(modelSecret model.Secret) (*generated.Secret, error) {
	secretType := modelSecret.GetType()
	switch secretType {
	case model.Binary:
		return c.fromModelSecret(modelSecret, generated.SecretType_BINARY)
	case model.Card:
		return c.fromModelSecret(modelSecret, generated.SecretType_CARD)
	case model.Credential:
		return c.fromModelSecret(modelSecret, generated.SecretType_CREDENTIAL)
	case model.Note:
		return c.fromModelSecret(modelSecret, generated.SecretType_NOTE)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert secret with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) ToModelEvent(generatedEvent *generated.SecretEvent) (*model.SecretEvent, error) {
	secretType := generatedEvent.Type
	switch secretType {
	case generated.EventType_INITIAL:
		return c.toModelEvent(generatedEvent, model.Initial)
	case generated.EventType_ADD:
		return c.toModelEvent(generatedEvent, model.Add)
	case generated.EventType_EDIT:
		return c.toModelEvent(generatedEvent, model.Edit)
	case generated.EventType_REMOVE:
		return c.toModelEvent(generatedEvent, model.Remove)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert event with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) FromModelEvent(modelEvent *model.SecretEvent) (*generated.SecretEvent, error) {
	secretType := modelEvent.Type
	switch secretType {
	case model.Initial:
		return c.fromModelEvent(modelEvent, generated.EventType_INITIAL)
	case model.Add:
		return c.fromModelEvent(modelEvent, generated.EventType_ADD)
	case model.Edit:
		return c.fromModelEvent(modelEvent, generated.EventType_EDIT)
	case model.Remove:
		return c.fromModelEvent(modelEvent, generated.EventType_REMOVE)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert event with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) toModelSecret(generatedSecret *generated.Secret, modelSecret model.Secret) (model.Secret, error) {
	err := c.serializer.Deserialize(generatedSecret.Content, modelSecret)
	if err != nil {
		return nil, logger.WrapError("deserialize secret", err)
	}

	return modelSecret, nil
}

func (c *Converter) fromModelSecret(secret model.Secret, secretType generated.SecretType) (*generated.Secret, error) {
	bytes, err := c.serializer.Serialize(secret)
	if err != nil {
		return nil, logger.WrapError("serialize secret", err)
	}

	return &generated.Secret{
		Identity: secret.GetIdentity(),
		Type:     secretType,
		Content:  bytes,
	}, nil
}

func (c *Converter) toModelEvent(
	generatedEvent *generated.SecretEvent,
	eventType model.EventType,
) (*model.SecretEvent, error) {
	modelSecret, err := c.ToModelSecret(generatedEvent.Secret)
	if err != nil {
		return nil, logger.WrapError("convert event secret", err)
	}

	return &model.SecretEvent{
		Type:   eventType,
		Secret: modelSecret,
	}, nil
}

func (c *Converter) fromModelEvent(
	modelEvent *model.SecretEvent,
	eventType generated.EventType,
) (*generated.SecretEvent, error) {
	generatedSecret, err := c.FromModelSecret(modelEvent.Secret)
	if err != nil {
		return nil, logger.WrapError("convert event secret", err)
	}

	return &generated.SecretEvent{
		Type:   eventType,
		Secret: generatedSecret,
	}, nil
}
