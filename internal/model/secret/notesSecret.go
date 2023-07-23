package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type notesSecret struct {
	*baseSecret

	text string
}

func newNotesSecret(text string, identity string, comment string) *notesSecret {
	return &notesSecret{
		baseSecret: newBaseSecret(identity, model.Credentials, comment),
		text:       text,
	}
}
