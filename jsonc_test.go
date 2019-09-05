package jsonc

import (
	"testing"
)

type test struct {
	name   string
	input  string
	output string
}

var tests = []test{
	{
		"Strip comments from new lines",
		`
			{
				// Comment
				"foo": "\"foo"
				// Comment
			}
		`,
		`
			{
				
				"foo": "\"foo"
				
			}
		`,
	},
	{
		"Strip comments that occur on same line as a JSON block",
		`
			{
				"foo": "\"foo" // Comment
			}
		`,
		`
			{
				"foo": "\"foo" 
			}
		`,
	},
	{
		"Keep comments that are within a JSON block",
		`
			{
				"foo": "\"foo // comment"
			}
		`,
		`
			{
				"foo": "\"foo // comment"
			}
		`,
	},
}

func TestLexer(t *testing.T) {
	for _, test := range tests {
		output := ToJSON(test.input)
		if test.output != output {
			t.Errorf("%s: %s != %s\n", test.name, test.output, output)
		}
	}
}
