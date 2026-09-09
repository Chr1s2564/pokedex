package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		}, {
			input:    "CAPSLOCK",
			expected: []string{"capslock"},
		}, {
			input:    "lowered multiple words",
			expected: []string{"lowered", "multiple", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("unmatching result")
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("unmatching word")
			}
		}
	}
}
