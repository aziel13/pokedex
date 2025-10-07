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
			input:    "  One tWo  tHree gO ",
			expected: []string{"one", "two", "three", "go"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) == len(c.expected) {
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				// Check each word in the slice
				// if they don't match, use t.Errorf to print an error message
				// and fail the test
				if word != expectedWord {
					t.Errorf("actual does not match expected\nactual: %v\nexpected: %v", word, expectedWord)
				}
			}
		} else {
			t.Errorf("actual length does not match expected  length\nactual: %v\nexpected: %v", len(actual), len(c.expected))
		}
	}
}
