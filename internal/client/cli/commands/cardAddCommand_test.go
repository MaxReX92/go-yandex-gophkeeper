package commands

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/identity"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/test"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
)

func TestNewCardAddCommand_CommonChecks(t *testing.T) {
	ctx := context.Background()
	const childCommandName = "childCommand"
	const childCommandDescription = "childDescription"
	keys := []string{childCommandName}

	stream := new(io.CommandStreamMock)
	generator := new(identity.GeneratorMock)
	secretStorage := new(storage.ClientSecretsStorageMock)
	childCommand := new(cli.CommandMock)

	childCommand.On("Name").Return(childCommandName)
	childCommand.On("ShortDescription").Return(childCommandDescription)
	childCommand.On("SetParent", mock.Anything)
	childCommand.On("Invoke", ctx, mock.Anything).Return(nil)
	stream.On("Write", mock.Anything)

	command := NewCardAddCommand(stream, generator, secretStorage, childCommand)
	command.ShowHelp()
	actualError := command.Invoke(ctx, keys)

	assert.Equal(t, "add", command.Name())
	assert.Equal(t, "add", command.FullName())
	assert.Equal(t, "add card secret", command.ShortDescription())
	assert.NoError(t, actualError)
	childCommand.AssertCalled(t, "SetParent", command.baseCommand)
	childCommand.AssertCalled(t, "Invoke", ctx, mock.MatchedBy(func(k []string) bool {
		assert.Len(t, k, 0)
		return true
	}))
}

func TestNewCardAddCommand_Invoke(t *testing.T) {
	ctx := context.Background()
	const secretIdentity = "secretIdentity"
	secretNumber := "secretNumber"
	secretCVV := "123"
	secretValid := "04/25"
	const secretComment = "secretComment"

	tests := []struct {
		name            string
		keys            []string
		storageError    error
		expectedError   string
		expectedOutput  string
		expectedComment string
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
			name:          "no_number",
			keys:          []string{"--cvv", secretCVV, "-v", secretValid},
			expectedError: cli.ErrRequiredArgNotFound.Error(),
		}, {
			name:          "no_cvv",
			keys:          []string{"-n", secretNumber, "-v", secretValid},
			expectedError: cli.ErrRequiredArgNotFound.Error(),
		}, {
			name:          "no_valid_thru",
			keys:          []string{"-n", secretNumber, "--cvv", secretCVV},
			expectedError: cli.ErrRequiredArgNotFound.Error(),
		}, {
			name: "success_short_key_no_comment",
			keys: []string{"-n", secretNumber, "--cvv", secretCVV, "-v", secretValid},
		}, {
			name:            "success_short_key_comment",
			keys:            []string{"-n", secretNumber, "--cvv", secretCVV, "-v", secretValid, "-c", secretComment},
			expectedComment: secretComment,
		}, {
			name: "success_full_key_no_comment",
			keys: []string{"--number", secretNumber, "--cvv", secretCVV, "--validThru", secretValid},
		}, {
			name:            "success_full_key_comment",
			keys:            []string{"--number", secretNumber, "--cvv", secretCVV, "--validThru", secretValid, "--comment", secretComment},
			expectedComment: secretComment,
		}, {
			name:          "storage_error",
			keys:          []string{"-n", secretNumber, "--cvv", secretCVV, "-v", secretValid},
			storageError:  test.ErrTest,
			expectedError: "failed to add secret: test error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := new(io.CommandStreamMock)
			generator := new(identity.GeneratorMock)
			secretStorage := new(storage.ClientSecretsStorageMock)
			childCommand := new(cli.CommandMock)

			childCommand.On("Name").Return("childCommand")
			childCommand.On("SetParent", mock.Anything)
			stream.On("Write", mock.Anything)
			generator.On("GenerateNewIdentity").Return(secretIdentity)
			secretStorage.On("AddSecret", ctx, mock.Anything).Return(tt.storageError)

			command := NewCardAddCommand(stream, generator, secretStorage, childCommand)
			actualError := command.Invoke(ctx, tt.keys)

			if tt.expectedError != "" {
				assert.ErrorContains(t, actualError, tt.expectedError)
			} else {
				secretStorage.AssertCalled(t, "AddSecret", ctx, mock.MatchedBy(func(s model.Secret) bool {
					assert.Equal(t, s.GetIdentity(), secretIdentity)
					assert.Equal(t, s.GetType(), model.Card)
					assert.Equal(t, s.GetComment(), tt.expectedComment)

					card, ok := s.(*secret.CardSecret)
					assert.True(t, ok)
					assert.Equal(t, card.Number, secretNumber)
					cvv, _ := parser.ToInt32(secretCVV)
					valid, _ := parser.ToTime(secretValid)
					assert.Equal(t, card.CVV, cvv)
					assert.Equal(t, card.Valid, valid)

					return true
				}))
			}

			if tt.expectedOutput != "" {
				stream.AssertCalled(t, "Write", tt.expectedOutput)
			}
		})
	}
}
