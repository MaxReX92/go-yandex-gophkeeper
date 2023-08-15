package secret

import "github.com/MaxReX92/go-yandex-gophkeeper/internal/model"

// NoteSecret represent a note secret.
type NoteSecret struct {
	*BaseSecret

	Text string
}

// NewNoteSecret creates a new instance of note secret.
func NewNoteSecret(text string, identity string, comment string) *NoteSecret {
	return &NoteSecret{
		BaseSecret: newBaseSecret(identity, model.Note, comment),
		Text:       text,
	}
}
