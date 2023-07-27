package parser_test

import (
	"testing"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestToFloat64(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expectError bool
		expected    float64
	}{
		{
			name:        "zero_success",
			value:       "0",
			expectError: false,
			expected:    0,
		},
		{
			name:        "positive_success",
			value:       "100",
			expectError: false,
			expected:    100,
		},
		{
			name:        "positive_float_success",
			value:       "100.001",
			expectError: false,
			expected:    100.001,
		},
		{
			name:        "negative_success",
			value:       "-100",
			expectError: false,
			expected:    -100,
		},
		{
			name:        "negative_float_success",
			value:       "-100.001",
			expectError: false,
			expected:    -100.001,
		},
		{
			name:        "emmpty_fail",
			value:       "",
			expectError: true,
		},
		{
			name:        "str_fail",
			value:       "str",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parser.ToFloat64(tt.value)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expectError bool
		expected    int64
	}{
		{
			name:        "zero_success",
			value:       "0",
			expectError: false,
			expected:    0,
		},
		{
			name:        "positive_success",
			value:       "100",
			expectError: false,
			expected:    100,
		},
		{
			name:        "positive_float_fail",
			value:       "100.001",
			expectError: true,
		},
		{
			name:        "negative_success",
			value:       "-100",
			expectError: false,
			expected:    -100,
		},
		{
			name:        "negative_float_fail",
			value:       "-100.001",
			expectError: true,
		},
		{
			name:        "emmpty_fail",
			value:       "",
			expectError: true,
		},
		{
			name:        "str_fail",
			value:       "str",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := parser.ToInt64(tt.value)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		value    float64
	}{
		{
			name:     "zero",
			value:    0,
			expected: "0",
		},
		{
			name:     "positive",
			value:    100,
			expected: "100",
		},
		{
			name:     "positive_float",
			value:    100.001,
			expected: "100.001",
		},
		{
			name:     "negative",
			value:    -100,
			expected: "-100",
		},
		{
			name:     "negative_float",
			value:    -100.001,
			expected: "-100.001",
		},
		{
			name:     "positive_double",
			value:    100.5555555555555555555555,
			expected: "100.55555555555556",
		},
		{
			name:     "negative_double",
			value:    -100.5555555555555555555555,
			expected: "-100.55555555555556",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, parser.FloatToString(tt.value))
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		value    int64
	}{
		{
			name:     "zero",
			value:    0,
			expected: "0",
		},
		{
			name:     "positive",
			value:    100,
			expected: "100",
		},
		{
			name:     "negative",
			value:    -100,
			expected: "-100",
		},
		{
			name:     "positive_long",
			value:    1000000000000000,
			expected: "1000000000000000",
		},
		{
			name:     "negative_long",
			value:    -1000000000000000,
			expected: "-1000000000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, parser.Int64ToString(tt.value))
		})
	}
}

func TestToTime(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		expected      time.Time
		expectedError error
	}{
		{
			name:          "notADate",
			value:         "test",
			expectedError: parser.ErrInvalidFormat,
		},
		{
			name:          "nowTime",
			value:         time.Now().String(),
			expectedError: parser.ErrInvalidFormat,
		},
		{
			name:          "fullYear",
			value:         "03/2026",
			expectedError: parser.ErrInvalidFormat,
		},
		{
			name:     "shortYear",
			value:    "03/26",
			expected: time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, actualError := parser.ToTime(tt.value)
			if tt.expectedError == nil {
				assert.Equal(t, tt.expected, actual)
			}

			assert.ErrorIs(t, actualError, tt.expectedError)
		})
	}
}
