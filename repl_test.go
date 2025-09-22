package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  jackie chan    ",
			expected: []string{"jackie", "chan"},
		},
		{
			input:    "some words",
			expected: []string{"some", "words"},
		},
		{
			input:    "khasab oman",
			expected: []string{"khasab", "oman"},
		},
		{
			input:    "   dubai    abudhabi   ",
			expected: []string{"dubai", "abudhabi"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			fmt.Println("actual length: ", len(actual))
			fmt.Println("c.expected length: ", len(c.expected))
			t.Errorf("Error: slice lengths are not same")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Error: words do not match")
			}
		}
	}
}
