package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander BuLbAsAur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "mewtwo",
			expected: []string{"mewtwo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect slice length. Expected: %v, Actual: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]

			if word != expected_word {
				t.Errorf("Slices do not match. Expected: %v, Actual: %v", c.expected, actual)
			}
		}
	}
}
