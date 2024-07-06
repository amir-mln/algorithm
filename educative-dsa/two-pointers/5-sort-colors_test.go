package twopointers_test

import (
	twopointers "educative-dsa/two-pointers"
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "test case 1",
			input:  []int{0, 1, 0},
			output: []int{0, 0, 1},
		},
		{
			name:   "test case 2",
			input:  []int{1},
			output: []int{1},
		},
		{
			name:   "test case 3",
			input:  []int{2, 2},
			output: []int{2, 2},
		},
		{
			name:   "test case 4",
			input:  []int{1, 1, 0, 2},
			output: []int{0, 1, 1, 2},
		},
		{
			name:   "test case 5",
			input:  []int{2, 1, 1, 0, 0},
			output: []int{0, 0, 1, 1, 2},
		},
		{
			name:   "test case 6",
			input:  []int{1},
			output: []int{1},
		},
		{
			name:   "test case 7",
			input:  []int{2, 2},
			output: []int{2, 2},
		},
		{
			name:   "test case 8",
			input:  []int{0, 1, 0},
			output: []int{0, 0, 1},
		},
		{
			name:   "test case 9",
			input:  []int{1, 1, 0, 2},
			output: []int{0, 1, 1, 2},
		},
		{
			name:   "test case 10",
			input:  []int{2, 1, 1, 0, 0},
			output: []int{0, 0, 1, 1, 2},
		},
		{
			name:   "test case 11",
			input:  []int{2, 2, 2, 0, 1, 0},
			output: []int{0, 0, 1, 2, 2, 2},
		},
		{
			name:   "test case 12",
			input:  []int{2, 1, 1, 0, 1, 0, 2},
			output: []int{0, 0, 1, 1, 1, 2, 2},
		},
		{
			name:   "test case 13",
			input:  []int{1, 0, 1},
			output: []int{0, 1, 1},
		},
	}

	for k, v := range twopointers.SortColors {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				cp := make([]int, len(c.input))
				copy(cp, c.input)

				v(cp)

				if !slices.Equal(cp, c.output) {
					t.Errorf("for input slice %v, expected %v; got %v", c.input, c.output, cp)
				}
			}
		})
	}
}
