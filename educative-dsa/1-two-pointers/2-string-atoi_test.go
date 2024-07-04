package twopointers_test

import (
	twopointers "educative-dsa/1-two-pointers"
	"testing"
)

func TestStringAtoi(t *testing.T) {
	cases := []struct {
		title  string
		input  string
		output int32
	}{
		{
			title:  "test case 1",
			input:  "1",
			output: 1,
		},
		{
			title:  "test case 1",
			input:  "79",
			output: 79,
		},
		{
			title:  "test case 2",
			input:  "king 25",
			output: 0,
		},
		{
			title:  "test case 3",
			input:  "  -33",
			output: -33,
		},
		{
			title:  "test case 4",
			input:  "-91283472332",
			output: -2147483648,
		},
		{
			title:  "test case 5",
			input:  "91283472332",
			output: 2147483647,
		},
	}

	for k, v := range twopointers.StringAtoi {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				result := v(c.input)

				if c.output != result {
					t.Errorf(
						"Failed test for case \"%s\" with input %s.\n expected %d got %d",
						c.title,
						c.input,
						c.output,
						result,
					)

				}
			}
		})
	}

}
