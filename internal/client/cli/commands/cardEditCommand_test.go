package commands

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/test"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
)

func TestNewCardEditCommand_CommonChecks(t *testing.T) {
	ctx := context.Background()
	const childCommandName = "childCommand"
	const childCommandDescription = "childDescription"
	keys := []string{childCommandName}

	stream := new(io.CommandStreamMock)
	secretStorage := new(storage.ClientSecretsStorageMock)
	childCommand := new(cli.CommandMock)

	childCommand.On("Name").Return(childCommandName)
	childCommand.On("ShortDescription").Return(childCommandDescription)
	childCommand.On("SetParent", mock.Anything)
	childCommand.On("Invoke", ctx, mock.Anything).Return(nil)
	stream.On("Write", mock.Anything)

	command := NewCardEditCommand(stream, secretStorage, childCommand)
	command.ShowHelp()
	actualError := command.Invoke(ctx, keys)

	assert.Equal(t, "edit", command.Name())
	assert.Equal(t, "edit", command.FullName())
	assert.Equal(t, "edit card secret", command.ShortDescription())
	assert.NoError(t, actualError)
	childCommand.AssertCalled(t, "SetParent", command.baseCommand)
	childCommand.AssertCalled(t, "Invoke", ctx, mock.MatchedBy(func(k []string) bool {
		assert.Len(t, k, 0)
		return true
	}))
}

func TestNewCardEditCommand_Invoke(t *testing.T) {
	ctx := context.Background()
	const secretIdentity = "secretIdentity"
	secretNum := "secretNum"
	secretCVV := 123
	secretValid, _ := parser.ToTime("04/25")
	const secretComment = "secretComment"

	tests := []struct {
		name             string
		keys             []string
		storageGetError  error
		storageEditError error
		expectedError    string
		expectedOutput   string
		expectedComment  string
	}{
		{
			name:          "no_keys",
			keys:          []string{},
			expectedError: cli.ErrRequiredArgNotFound.Error(),
		}, {
			name:          "unknown_keys",
			keys:          []string{"--test"},
			expectedError: fmt.Errorf("failed to parse unexpected key: --test: %w", cli.ErrInvalidArguments).Error(),
		}, {
			name:          "no_identity_key",
			keys:          []string{"-n", secretNum, "--cvv", "123", "-v", "04/25"},
			expectedError: "failed to invoke edit command: secret identity is missed: required arg not found",
		}, {
			name:            "success_short_key_no_comment",
			keys:            []string{"-id", secretIdentity, "-n", secretNum, "--cvv", "123", "-v", "04/25"},
			expectedComment: "test",
		}, {
			name:            "success_short_key_comment",
			keys:            []string{"-id", secretIdentity, "-n", secretNum, "--cvv", "123", "-v", "04/25", "-c", secretComment},
			expectedComment: secretComment,
		}, {
			name:            "success_full_key_no_comment",
			keys:            []string{"--identity", secretIdentity, "--number", secretNum, "--cvv", "123", "--validThru", "04/25"},
			expectedComment: "test",
		}, {
			name:            "success_full_key_comment",
			keys:            []string{"--identity", secretIdentity, "--number", secretNum, "--cvv", "123", "--validThru", "04/25", "--comment", secretComment},
			expectedComment: secretComment,
		}, {
			name:            "storage_get_error",
			keys:            []string{"-id", secretIdentity, "-n", secretNum, "--cvv", "123", "-v", "04/25"},
			storageGetError: test.ErrTest,
			expectedError:   "failed to get secret: test error message",
		}, {
			name:             "storage_add_error",
			keys:             []string{"-id", secretIdentity, "-n", secretNum, "--cvv", "123", "-v", "04/25"},
			storageEditError: test.ErrTest,
			expectedError:    "failed to edit secret: test error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentSecret := secret.NewCardSecret("test", 777, time.Now(), secretIdentity, "test")

			stream := new(io.CommandStreamMock)
			secretStorage := new(storage.ClientSecretsStorageMock)
			childCommand := new(cli.CommandMock)

			childCommand.On("Name").Return("childCommand")
			childCommand.On("SetParent", mock.Anything)
			stream.On("Write", mock.Anything)
			secretStorage.On("GetSecretByID", ctx, model.Card, secretIdentity).Return(currentSecret, tt.storageGetError)
			secretStorage.On("ChangeSecret", ctx, mock.Anything).Return(tt.storageEditError)

			command := NewCardEditCommand(stream, secretStorage, childCommand)
			actualError := command.Invoke(ctx, tt.keys)

			if tt.expectedError != "" {
				assert.ErrorContains(t, actualError, tt.expectedError)
			} else {
				secretStorage.AssertCalled(t, "ChangeSecret", ctx, mock.MatchedBy(func(s model.Secret) bool {
					assert.Equal(t, secretIdentity, s.GetIdentity())
					assert.Equal(t, model.Card, s.GetType())
					assert.Equal(t, tt.expectedComment, s.GetComment())

					card, ok := s.(*secret.CardSecret)
					assert.True(t, ok)
					assert.Equal(t, secretNum, card.Number)
					assert.Equal(t, int32(secretCVV), card.CVV)
					assert.Equal(t, secretValid, card.Valid)

					return true
				}))
			}

			if tt.expectedOutput != "" {
				stream.AssertCalled(t, "Write", tt.expectedOutput)
			}
		})
	}
}
