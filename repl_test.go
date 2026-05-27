package main

import (
	"testing"
)

func TestXxx(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "How are you",
			expected: []string{"how", "are", "you"},
		},
	}

	for _, test := range cases {
		result := clearInput(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("Length mismatch: expected %d, got %d\n", len(test.expected), len(result))
		}
		for i, _ := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Expected %v\n Actual %v\n ", test.expected, result)
			}
		}
	}
}
