package educative_test

import (
	"slices"
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestProductOfArray(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "test case 1",
			input:  []int{0, -1, 2, -3, 4, -2},
			output: []int{-48, 0, 0, 0, 0, 0},
		},
		{
			name:   "test case 2",
			input:  []int{5, 3, -1, 6, 4},
			output: []int{-72, -120, 360, -60, -90},
		},
		{
			name:   "test case 3",
			input:  []int{-7, 6, 4, 3, 1, 2},
			output: []int{144, -168, -252, -336, -1008, -504},
		},
		{
			name:   "test case 4",
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			name:   "test case 5",
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
		{
			name:   "test case 6",
			input:  []int{1, 0, 3, 5, 2},
			output: []int{0, 30, 0, 0, 0},
		},
		{
			name:   "test case 7",
			input:  []int{12, 10, 7, 2, 1, 15},
			output: []int{2100, 2520, 3600, 12600, 25200, 1680},
		},
		{
			name:   "test case 8",
			input:  []int{0, -1, 2, -3, 4, -2},
			output: []int{-48, 0, 0, 0, 0, 0},
		},
		{
			name:   "test case 9",
			input:  []int{5, 3, -1, 6, 4},
			output: []int{-72, -120, 360, -60, -90},
		},
		{
			name:   "test case 10",
			input:  []int{-7, 6, 4, 3, 1, 2},
			output: []int{144, -168, -252, -336, -1008, -504},
		},
		{
			name:   "test case 11",
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
		{
			name:   "test case 12",
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
	}

	for k, v := range educative.ProductOfArray {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := v(c.input)

				if !slices.Equal(res, c.output) {
					t.Errorf("for input %v of %q, expected %v; got %v", c.input, c.name, c.output, res)
				}
			}
		})
	}

}
