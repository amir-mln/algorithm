package twopointers_test

import (
	twopointers "educative-dsa/two-pointers"
	"testing"
)

func TestAddTwoBinary(t *testing.T) {
	cases := []struct {
		input1 string
		input2 string
		result string
	}{
		{
			input1: "1111",
			input2: "1011",
			result: "11010",
		},
		{
			input1: "101010",
			input2: "110110",
			result: "1100000",
		},
		{
			input1: "1111",
			input2: "111",
			result: "10110",
		},
		{
			input1: "101",
			input2: "1101",
			result: "10010",
		},
		{
			input1: "1",
			input2: "1",
			result: "10",
		},
		{
			input1: "0",
			input2: "0",
			result: "0",
		},
		{
			input1: "110010",
			input2: "11110",
			result: "1010000",
		},
		{
			input1: "10000",
			input2: "1000",
			result: "11000",
		},
		{
			input1: "101010101011011111011001101011",
			input2: "1",
			result: "101010101011011111011001101100",
		},
	}

	for k, v := range twopointers.AddBinary {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := v(c.input1, c.input2)
				if res != c.result {
					t.Errorf("Failed test for inputs %s and %s; expected %s got %s", c.input1, c.input2, c.result, res)
				}
			}
		})
	}
}
