package main

import "testing"

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
			input:    "hello ",
			expected: []string{"hello"},
		},
		{
			input:    "HELLO, world",
			expected: []string{"hello,", "world"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("number of strings do not match: %v, %v", c.expected, actual)
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match")
				t.Fail()
			}
		}
	}
}
