package grpc

import (
	"fmt"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/crypto"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/serialization"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type Converter struct {
	serializer serialization.Serializer
	encryptor  crypto.Encryptor
	decryptor  crypto.Decryptor
}

func NewConverter(serializer serialization.Serializer, encryptor crypto.Encryptor, decryptor crypto.Decryptor) *Converter {
	return &Converter{
		serializer: serializer,
		encryptor:  encryptor,
		decryptor:  decryptor,
	}
}

func (c *Converter) ToModelSecret(generatedSecret *generated.Secret, key string) (model.Secret, error) {
	secretType := generatedSecret.Type
	switch secretType {
	case generated.SecretType_BINARY:
		return c.toModelSecret(generatedSecret, key, &secret.BinarySecret{})
	case generated.SecretType_CARD:
		return c.toModelSecret(generatedSecret, key, &secret.CardSecret{})
	case generated.SecretType_CREDENTIAL:
		return c.toModelSecret(generatedSecret, key, &secret.CredentialSecret{})
	case generated.SecretType_NOTE:
		return c.toModelSecret(generatedSecret, key, &secret.NoteSecret{})
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert secret with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) FromModelSecret(modelSecret model.Secret, key string) (*generated.Secret, error) {
	secretType := modelSecret.GetType()
	switch secretType {
	case model.Binary:
		return c.fromModelSecret(modelSecret, key, generated.SecretType_BINARY)
	case model.Card:
		return c.fromModelSecret(modelSecret, key, generated.SecretType_CARD)
	case model.Credential:
		return c.fromModelSecret(modelSecret, key, generated.SecretType_CREDENTIAL)
	case model.Note:
		return c.fromModelSecret(modelSecret, key, generated.SecretType_NOTE)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert secret with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) ToModelEvent(generatedEvent *generated.SecretEvent, key string) (*model.SecretEvent, error) {
	secretType := generatedEvent.Type
	switch secretType {
	case generated.EventType_INITIAL:
		return c.toModelEvent(generatedEvent, key, model.Initial)
	case generated.EventType_ADD:
		return c.toModelEvent(generatedEvent, key, model.Add)
	case generated.EventType_EDIT:
		return c.toModelEvent(generatedEvent, key, model.Edit)
	case generated.EventType_REMOVE:
		return c.toModelEvent(generatedEvent, key, model.Remove)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert event with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) FromModelEvent(modelEvent *model.SecretEvent, key string) (*generated.SecretEvent, error) {
	secretType := modelEvent.Type
	switch secretType {
	case model.Initial:
		return c.fromModelEvent(modelEvent, key, generated.EventType_INITIAL)
	case model.Add:
		return c.fromModelEvent(modelEvent, key, generated.EventType_ADD)
	case model.Edit:
		return c.fromModelEvent(modelEvent, key, generated.EventType_EDIT)
	case model.Remove:
		return c.fromModelEvent(modelEvent, key, generated.EventType_REMOVE)
	default:
		return nil, logger.WrapError(fmt.Sprintf("convert event with type %v", secretType), model.ErrUnknownType)
	}
}

func (c *Converter) toModelSecret(generatedSecret *generated.Secret, key string, modelSecret model.Secret) (model.Secret, error) {
	decrypted, err := c.decryptor.Decrypt(generatedSecret.Content, []byte(key))
	if err != nil {
		return nil, logger.WrapError("decrypt content", err)
	}

	err = c.serializer.Deserialize(decrypted, modelSecret)
	// err := c.serializer.Deserialize(generatedSecret.Content, modelSecret)
	if err != nil {
		return nil, logger.WrapError("deserialize secret", err)
	}

	return modelSecret, nil
}

func (c *Converter) fromModelSecret(secret model.Secret, key string, secretType generated.SecretType) (*generated.Secret, error) {
	bytes, err := c.serializer.Serialize(secret)
	if err != nil {
		return nil, logger.WrapError("serialize secret", err)
	}

	encrypted, err := c.encryptor.Encrypt(bytes, []byte(key))
	if err != nil {
		return nil, logger.WrapError("encrypt content", err)
	}

	return &generated.Secret{
		Identity: secret.GetIdentity(),
		Type:     secretType,
		Content:  encrypted,
		// Content: bytes,
	}, nil
}

func (c *Converter) toModelEvent(
	generatedEvent *generated.SecretEvent,
	key string,
	eventType model.EventType,
) (*model.SecretEvent, error) {
	modelSecret, err := c.ToModelSecret(generatedEvent.Secret, key)
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
	key string,
	eventType generated.EventType,
) (*generated.SecretEvent, error) {
	generatedSecret, err := c.FromModelSecret(modelEvent.Secret, key)
	if err != nil {
		return nil, logger.WrapError("convert event secret", err)
	}

	return &generated.SecretEvent{
		Type:   eventType,
		Secret: generatedSecret,
	}, nil
}
