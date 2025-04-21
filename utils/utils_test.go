package utils

import (
	"testing"
)

func TestNormalizeSuiAddress(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		description string
	}{
		{
			input:       "abc",
			expected:    "0x0000000000000000000000000000000000000000000000000000000000000abc",
			description: "lowercase without 0x",
		},
		{
			input:       "0xabc",
			expected:    "0x0000000000000000000000000000000000000000000000000000000000000abc",
			description: "with 0x, no forceAdd0x",
		},
		{
			input:       "0xABC",
			expected:    "0x0000000000000000000000000000000000000000000000000000000000000abc",
			description: "uppercase input, normalize to lowercase",
		},
		{
			input:       "abc",
			expected:    "0x0000000000000000000000000000000000000000000000000000000000000abc",
			description: "forceAdd0x with plain input",
		},
		{
			input:       "0xabc",
			expected:    "0x0000000000000000000000000000000000000000000000000000000000000abc",
			description: "forceAdd0x true keeps 0xabc and still pads correctly",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := NormalizeSuiAddress(test.input)
			if string(result) != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
