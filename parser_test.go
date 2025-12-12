package unitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedValue float64
		expectedUnit  string
		expectError   bool
	}{
		{
			name:          "whitespace-none",
			input:         "7k€",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "whitespace-center",
			input:         "7 k€",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "whitespace-beginning",
			input:         "  7k€",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "whitespace-end",
			input:         "7k€  ",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "whitespace-all",
			input:         "  7 k€  ",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "whitespace-beginning",
			input:         " 7k€",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "format-beginning",
			input:         "k€7",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "format-beginning-space",
			input:         "k€ 7",
			expectedValue: 7.0,
			expectedUnit:  "k€",
		},
		{
			name:          "decimal",
			input:         "5.02€",
			expectedValue: 5.02,
			expectedUnit:  "€",
		},
		{
			name:          "thousands-separator",
			input:         "5,000.02 €",
			expectedValue: 5_000.02,
			expectedUnit:  "€",
		},
		{
			name:          "thousands-thousands-separator",
			input:         "5,000,000.02 €",
			expectedValue: 5_000_000.02,
			expectedUnit:  "€",
		},
		{
			name:        "decimal-twice",
			input:       "8.000.32 €",
			expectError: true,
		},
		{
			name:          "unit-twice",
			input:         "€ 8.32 €",
			expectedValue: 8.32,
			expectedUnit:  "€",
		},
		{
			name:          "unit-none",
			input:         "5",
			expectedValue: 5.0,
			expectedUnit:  "",
		},
		{
			name:          "text-multiple",
			input:         "Value: 5€",
			expectedValue: 5.0,
			expectedUnit:  "€",
		},
	}

	p := &BasicNumberParser{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			returnedValue, returnedUnit, returnedErr := p.Parse(tc.input, "€")
			// Check error
			assert.Equal(t, tc.expectError, returnedErr != nil)
			// Check output
			if returnedErr == nil {
				assert.Equal(t, tc.expectedValue, returnedValue)
				assert.Equal(t, tc.expectedUnit, returnedUnit)
			}
		})
	}
}
