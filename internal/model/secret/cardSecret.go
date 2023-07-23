package secret

import (
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
)

type cardSecret struct {
	*baseSecret

	number string
	cvv    int32
	valid  time.Time
}

func newCardSecret(number string, cvv int32, valid time.Time, identity string, comment string) *cardSecret {
	return &cardSecret{
		baseSecret: newBaseSecret(identity, model.Card, comment),
		number:     number,
		cvv:        cvv,
		valid:      valid,
	}
}
