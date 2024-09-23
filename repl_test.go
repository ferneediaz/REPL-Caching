package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  trim   spaces   ",
			expected: []string{"trim", "spaces"},
		},
		{
			input:    "UPPERCASE lowercase",
			expected: []string{"uppercase", "lowercase"},
		},
		{
			input:    "multiple   spaces   between   words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "special!@#characters",
			expected: []string{"specialcharacters"},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("For input '%s': expected length %d, got %d", cs.input, len(cs.expected), len(actual))
			continue
		}

		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("For input '%s': at index %d, expected '%s', got '%s'", cs.input, i, cs.expected[i], actual[i])
			}
		}
	}
}
