package educative_test

import (
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestTrappingRainWater(t *testing.T) {
	cases := []struct {
		name    string
		heights []int
		output  int
	}{
		{
			name:    "test case 1",
			heights: []int{1, 0, 3, 0, 1, 2},
			output:  4,
		},
		{
			name:    "test case 2",
			heights: []int{4, 2, 0, 3, 1, 5},
			output:  10,
		},
		{
			name:    "test case 3",
			heights: []int{0, 3, 0, 2, 1, 0, 1, 4, 2, 1, 2, 0},
			output:  12,
		},
		{
			name:    "test case 4",
			heights: []int{1, 2, 2, 2, 2, 2},
			output:  0,
		},
		{
			name:    "test case 5",
			heights: []int{4},
			output:  0,
		},
		{
			name:    "test case 6",
			heights: []int{1, 1, 3, 1, 1, 2},
			output:  2,
		},
		{
			name:    "test case 7",
			heights: []int{4, 2, 0, 3, 1, 5},
			output:  10,
		},
		{
			name:    "test case 8",
			heights: []int{7, 2, 3, 4, 0, 7, 1, 8, 9, 4, 9, 1, 10, 4, 1, 5, 5, 0, 1, 3, 10, 6, 8, 2, 4, 7, 3, 3, 2, 5, 5, 1, 9, 5, 2, 1, 7, 10, 2, 1, 6, 1, 0, 10, 10, 9, 6, 6, 2, 6, 7, 0, 1, 1, 8, 9, 3, 7, 10, 10, 6, 7, 7, 0, 8, 10, 5, 4, 6, 0, 4, 7, 6, 10, 3, 9, 6, 2, 1, 0, 6, 10, 4, 6, 9, 3, 7, 0, 3, 1, 8, 7, 1, 8, 3, 7, 5, 5, 0, 6},
			output:  441,
		},
		{
			name:    "test case 9",
			heights: []int{0, 0, 3, 2, 2, 2, 0, 0, 1},
			output:  2,
		},
		{
			name:    "test case 10",
			heights: []int{4, 1, 0, 0, 0, 1, 0, 0, 0, 4},
			output:  30,
		},
	}

	for k, s := range educative.TrappingRainWater {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := s(c.heights)

				if res != c.output {
					t.Errorf(
						"in %q with heights: %v, expected %d; got %d ",
						c.name,
						c.heights,
						c.output,
						res,
					)
				}
			}
		})
	}
}
