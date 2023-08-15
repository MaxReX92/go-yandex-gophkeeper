package commands

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCardGetCommand_CommonChecks(t *testing.T) {
	ctx := context.Background()
	childCommandName := "childCommand"
	childCommandDescription := "childDescription"
	keys := []string{childCommandName}

	stream := new(io.CommandStreamMock)
	secretStorage := new(storage.ClientSecretsStorageMock)
	childCommand := new(cli.CommandMock)

	childCommand.On("Name").Return(childCommandName)
	childCommand.On("ShortDescription").Return(childCommandDescription)
	childCommand.On("SetParent", mock.Anything)
	childCommand.On("Invoke", ctx, mock.Anything).Return(nil)
	stream.On("Write", mock.Anything)

	command := NewCardGetCommand(stream, secretStorage, childCommand)
	command.ShowHelp()
	actualError := command.Invoke(ctx, keys)

	assert.Equal(t, "get", command.Name())
	assert.Equal(t, "get", command.FullName())
	assert.Equal(t, "get card secret", command.ShortDescription())
	assert.NoError(t, actualError)
	childCommand.AssertCalled(t, "SetParent", command.baseCommand)
	childCommand.AssertCalled(t, "Invoke", ctx, mock.MatchedBy(func(k []string) bool {
		assert.Len(t, k, 0)
		return true
	}))
}

func TestNewCardGetCommand_Invoke(t *testing.T) {
	ctx := context.Background()
	secretIdentity := "secretIdentity"
	tests := []struct {
		name            string
		keys            []string
		storageGetError error
		expectedError   string
		expectedOutput  string
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
			keys:          []string{"-r"},
			expectedError: "failed to invoke get command: secret identity is missed: required arg not found",
		}, {
			name:           "success_short_key_no_reveal",
			keys:           []string{"-id", secretIdentity},
			expectedOutput: "\tsecretIdentity\t\tsecretNumber\t\t2025-05-01 00:00:00 +0000 UTC\t\t***\t\tsecretComment",
		}, {
			name:           "success_short_key_reveal",
			keys:           []string{"-id", secretIdentity, "-r"},
			expectedOutput: "\tsecretIdentity\t\tsecretNumber\t\t2025-05-01 00:00:00 +0000 UTC\t\t123\t\tsecretComment",
		}, {
			name:           "success_full_key_no_reveal",
			keys:           []string{"--identity", secretIdentity},
			expectedOutput: "\tsecretIdentity\t\tsecretNumber\t\t2025-05-01 00:00:00 +0000 UTC\t\t***\t\tsecretComment",
		}, {
			name:           "success_full_key_reveal",
			keys:           []string{"--identity", secretIdentity, "--reveal"},
			expectedOutput: "\tsecretIdentity\t\tsecretNumber\t\t2025-05-01 00:00:00 +0000 UTC\t\t123\t\tsecretComment",
		}, {
			name:            "storage_get_error",
			keys:            []string{"-id", secretIdentity},
			storageGetError: errors.New("test error message"),
			expectedError:   "failed to get card secret: test error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, _ := parser.ToTime("05/25")
			currentSecret := secret.NewCardSecret("secretNumber", 123, valid, secretIdentity, "secretComment")

			stream := new(io.CommandStreamMock)
			secretStorage := new(storage.ClientSecretsStorageMock)
			childCommand := new(cli.CommandMock)

			childCommand.On("Name").Return("childCommand")
			childCommand.On("SetParent", mock.Anything)
			stream.On("Write", mock.Anything)
			secretStorage.On("GetSecretByID", ctx, model.Card, secretIdentity).Return(currentSecret, tt.storageGetError)

			command := NewCardGetCommand(stream, secretStorage, childCommand)
			actualError := command.Invoke(ctx, tt.keys)

			if tt.expectedError != "" {
				assert.ErrorContains(t, actualError, tt.expectedError)
			}

			if tt.expectedOutput != "" {
				stream.AssertCalled(t, "Write", tt.expectedOutput)
			}
		})
	}
}
