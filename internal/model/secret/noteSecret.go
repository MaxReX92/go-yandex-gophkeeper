package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

type NoteSecret struct {
	*BaseSecret

	Text string
}

func NewNoteSecret(text string, identity string, comment string) *NoteSecret {
	return &NoteSecret{
		BaseSecret: newBaseSecret(identity, model.Note, comment),
		Text:       text,
	}
}
