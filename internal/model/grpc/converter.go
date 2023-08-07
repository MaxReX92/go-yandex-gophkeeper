package grpc

import (
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/generated"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) FromModelSecret(secret model.Secret) (*generated.Secret, error) {

}

func (c *Converter) ToModelSecret(secret *generated.Secret) (model.Secret, error) {

}

func (c *Converter) ToModelEvent(secretEvent *generated.SecretEvent) (model.SecretEvent, error) {

}
