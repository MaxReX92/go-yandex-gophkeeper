package commands

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/cli"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/io"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/client/storage"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/identity"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBinaryAddCommand_CommonChecks(t *testing.T) {
	ctx := context.Background()
	childCommandName := "childCommand"
	childCommandDescription := "childDescription"
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

	command := NewBinaryAddCommand(stream, generator, secretStorage, childCommand)
	command.ShowHelp()
	actualError := command.Invoke(ctx, keys)

	assert.Equal(t, "add", command.Name())
	assert.Equal(t, "add", command.FullName())
	assert.Equal(t, "add binary secret", command.ShortDescription())
	assert.NoError(t, actualError)
	childCommand.AssertCalled(t, "SetParent", command.baseCommand)
	childCommand.AssertCalled(t, "Invoke", ctx, mock.MatchedBy(func(k []string) bool {
		assert.Len(t, k, 0)
		return true
	}))
}

func TestNewBinaryAddCommand_Invoke(t *testing.T) {
	ctx := context.Background()
	secretIdentity := "secretIdentity"
	secretName := "secretName"
	secretContent := "secretContent"
	secretComment := "secretComment"

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
			name:          "no_name",
			keys:          []string{"-b", secretContent},
			expectedError: "failed to invoke add command: name is missed: required arg not found",
		}, {
			name:          "no_content",
			keys:          []string{"-n", secretName},
			expectedError: "failed to add binary: either file path or base64 string is expected: invalid arguments",
		}, {
			name: "success_short_key_no_comment",
			keys: []string{"-n", secretName, "-b", secretContent},
		}, {
			name:            "success_short_key_comment",
			keys:            []string{"-n", secretName, "-b", secretContent, "-c", secretComment},
			expectedComment: secretComment,
		}, {
			name: "success_full_key_no_comment",
			keys: []string{"--name", secretName, "--base64", secretContent},
		}, {
			name:            "success_full_key_comment",
			keys:            []string{"--name", secretName, "--base64", secretContent, "--comment", secretComment},
			expectedComment: secretComment,
		}, {
			name:          "storage_error",
			keys:          []string{"-n", secretName, "-b", secretContent},
			storageError:  errors.New("test error message"),
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

			command := NewBinaryAddCommand(stream, generator, secretStorage, childCommand)
			actualError := command.Invoke(ctx, tt.keys)

			if tt.expectedError != "" {
				assert.ErrorContains(t, actualError, tt.expectedError)
			} else {
				secretStorage.AssertCalled(t, "AddSecret", ctx, mock.MatchedBy(func(s model.Secret) bool {
					assert.Equal(t, s.GetIdentity(), secretIdentity)
					assert.Equal(t, s.GetType(), model.Binary)
					assert.Equal(t, s.GetComment(), tt.expectedComment)

					binary, ok := s.(*secret.BinarySecret)
					assert.True(t, ok)
					assert.Equal(t, binary.Name, secretName)
					assert.NotNil(t, binary.Reader)

					return true
				}))
			}

			if tt.expectedOutput != "" {
				stream.AssertCalled(t, "Write", tt.expectedOutput)
			}
		})
	}
}
