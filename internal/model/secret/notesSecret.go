package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type NotesSecret struct {
	*BaseSecret

	Text string
}

func NewNotesSecret(text string, identity string, comment string) *NotesSecret {
	return &NotesSecret{
		BaseSecret: newBaseSecret(identity, model.Credentials, comment),
		Text:       text,
	}
}
