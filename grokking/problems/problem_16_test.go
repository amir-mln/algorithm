package problems_test

import (
	"testing"

	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestMinInRotatedArray(t *testing.T) {
	cases := []struct {
		name    string
		numbers []int
		output  int
	}{
		{
			name:    "test case 1",
			numbers: []int{5, 7, 17, 21, 22, 2, 4},
			output:  2,
		},
		{
			name:    "test case 2",
			numbers: []int{0, 12, 15, 100, 121, 289, -17, -2, -1},
			output:  -17,
		},
		{
			name:    "test case 3",
			numbers: []int{34, 45, 72, 85, -98, -74, -2, -1},
			output:  -98,
		},
		{
			name:    "test case 4",
			numbers: []int{55, 67, 88, 4, 8, 12, 15, 19, 28},
			output:  4,
		},
		{
			name:    "test case 5",
			numbers: []int{-33, -23, -11, -7, -100, -70, -50, -45},
			output:  -100,
		},
		{
			name:    "test case 6",
			numbers: []int{-1, 0, -4, -2},
			output:  -4,
		},
		{
			name:    "test case 7",
			numbers: []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4},
			output:  1,
		},
		{
			name:    "test case 8",
			numbers: []int{34},
			output:  34,
		},
		{
			name:    "test case 9",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, -19, -18},
			output:  -19,
		},
		{
			name:    "test case 10",
			numbers: []int{1, 2, 3, -3, -2, -1},
			output:  -3,
		},
	}

	for k, f := range problems.MinInRotatedArray {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				o := f(c.numbers)

				if o != c.output {
					t.Errorf(
						"in %q with numbers %v, expected %d; got %d ",
						c.name,
						c.numbers,
						c.output,
						o,
					)
				}
			}
		})

	}
}
