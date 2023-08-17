package secret

import (
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

// CardSecret represent a card secret.
type CardSecret struct {
	*BaseSecret

	Number string
	CVV    int32
	Valid  time.Time
}

// NewCardSecret creates a new instance of card secret.
func NewCardSecret(number string, cvv int32, valid time.Time, identity string, comment string) *CardSecret {
	return &CardSecret{
		BaseSecret: newBaseSecret(identity, model.Card, comment),
		Number:     number,
		CVV:        cvv,
		Valid:      valid,
	}
}
