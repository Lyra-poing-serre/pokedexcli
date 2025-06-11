package tests

import (
	"testing"

	"github.com/Lyra-poing-serre/pokedexcli/internal/repl"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "cashy  OS FTW",
			expected: []string{"cashy", "os", "ftw"},
		},
	}
	for _, c := range cases {
		actual := repl.CleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Actual len != than Expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word != than Expected word")
			}
		}
	}
}
