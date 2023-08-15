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
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/model/secret"
	"github.com/MaxReX92/go-yandex-gophkeeper/internal/test"
)

func TestNewCredentialListCommand_CommonChecks(t *testing.T) {
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

	command := NewCredentialListCommand(stream, secretStorage, childCommand)
	command.ShowHelp()
	actualError := command.Invoke(ctx, keys)

	assert.Equal(t, "list", command.Name())
	assert.Equal(t, "list", command.FullName())
	assert.Equal(t, "list of all credential", command.ShortDescription())
	assert.NoError(t, actualError)
	childCommand.AssertCalled(t, "SetParent", command.baseCommand)
	childCommand.AssertCalled(t, "Invoke", ctx, mock.MatchedBy(func(k []string) bool {
		assert.Len(t, k, 0)
		return true
	}))
}

func TestNewCredentialListCommand_Invoke(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name            string
		keys            []string
		storageGetError error
		expectedError   string
		expectedOutput  []string
	}{
		{
			name:          "unknown_keys",
			keys:          []string{"--test"},
			expectedError: fmt.Errorf("failed to parse unexpected key: --test: %w", cli.ErrInvalidArguments).Error(),
		}, {
			name: "no_keys",
			keys: []string{},
			expectedOutput: []string{
				"\tsecretIdentity1\t\tsecretName1\t\t***\t\tsecretComment1\n",
				"\tsecretIdentity2\t\tsecretName2\t\t***\t\tsecretComment2\n",
			},
		}, {
			name: "success_short_reveal",
			keys: []string{"-r"},
			expectedOutput: []string{
				"\tsecretIdentity1\t\tsecretName1\t\tsecretPass1\t\tsecretComment1\n",
				"\tsecretIdentity2\t\tsecretName2\t\tsecretPass2\t\tsecretComment2\n",
			},
		}, {
			name: "success_full_reveal",
			keys: []string{"--reveal"},
			expectedOutput: []string{
				"\tsecretIdentity1\t\tsecretName1\t\tsecretPass1\t\tsecretComment1\n",
				"\tsecretIdentity2\t\tsecretName2\t\tsecretPass2\t\tsecretComment2\n",
			},
		}, {
			name:            "storage_list_error",
			keys:            []string{},
			storageGetError: test.ErrTest,
			expectedError:   "failed to get secrets: test error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentSecret1 := secret.NewCredentialSecret("secretName1", "secretPass1", "secretIdentity1", "secretComment1")
			currentSecret2 := secret.NewCredentialSecret("secretName2", "secretPass2", "secretIdentity2", "secretComment2")

			stream := new(io.CommandStreamMock)
			secretStorage := new(storage.ClientSecretsStorageMock)
			childCommand := new(cli.CommandMock)

			childCommand.On("Name").Return("childCommand")
			childCommand.On("SetParent", mock.Anything)
			stream.On("Write", mock.Anything)
			secretStorage.On("GetAllSecrets", ctx, model.Credential).Return([]model.Secret{currentSecret1, currentSecret2}, tt.storageGetError)

			command := NewCredentialListCommand(stream, secretStorage, childCommand)
			actualError := command.Invoke(ctx, tt.keys)

			if tt.expectedError != "" {
				assert.ErrorContains(t, actualError, tt.expectedError)
			}

			if len(tt.expectedOutput) > 0 {
				for _, expectedOutput := range tt.expectedOutput {
					stream.AssertCalled(t, "Write", expectedOutput)
				}
			}
		})
	}
}
