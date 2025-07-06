package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},{
			input: "   something   is com ing  today  ",
			expected: []string{"something", "is", "com", "ing", "today"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch for input '%s': expected %d words, got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Mismatch at index %d for input '%s': expected '%s', got '%s'", i, c.input, expectedWord, word)
			}
		}
	}
}
