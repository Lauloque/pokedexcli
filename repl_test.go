/* SPDX-License-Identifier: GPL-3.0-or-later */
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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  owo uwu  ",
			expected: []string{"owo", "uwu"},
		},
		{
			input:    "  owouwu  ",
			expected: []string{"owouwu"},
		},
		{
			input:    "one    two   three four ",
			expected: []string{"one", "two", "three", "four"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("incorrect length, expected '%d', got '%d'", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("incorrect word: expected '%s', got '%s'", expectedWord, word)
			}
		}
	}
}
