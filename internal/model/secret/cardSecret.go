package secret

import (
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type CardSecret struct {
	*BaseSecret

	Number string
	CVV    int32
	Valid  time.Time
}

func NewCardSecret(number string, cvv int32, valid time.Time, identity string, comment string) *CardSecret {
	return &CardSecret{
		BaseSecret: newBaseSecret(identity, model.Card, comment),
		Number:     number,
		CVV:        cvv,
		Valid:      valid,
	}
}
