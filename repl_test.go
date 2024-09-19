package main

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
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
			expected: []string{"special", "characters"},
		},
		}
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input){
			if len(actual) != len(cs.expected){
				t.Errorf("the length are not equal: %v vs %v"),
					len(actual),
					len(cs.expected).

			}
			continue
		}
		for i := range actual{
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v")
			}
		}
	} 
