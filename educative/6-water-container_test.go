package educative_test

import (
	"testing"

	"github.com/amir-mln/algorithm/educative"
)

func TestContainerWithMostWater(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "test case 1",
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			name:   "test case 2",
			input:  []int{1, 1},
			output: 1,
		},
		{
			name:   "test case 3",
			input:  []int{2, 8, 6, 3, 5, 4, 7},
			output: 35,
		},
		{
			name:   "test case 4",
			input:  []int{1, 5},
			output: 1,
		},
		{
			name:   "test case 5",
			input:  []int{7, 7, 2},
			output: 7,
		},
		{
			name:   "test case 6",
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			name:   "test case 7",
			input:  []int{1, 1},
			output: 1,
		},
		{
			name:   "test case 8",
			input:  []int{2, 8, 6, 3, 5, 4, 7},
			output: 35,
		},
		{
			name:   "test case 9",
			input:  []int{1, 5},
			output: 1,
		},
		{
			name:   "test case 10",
			input:  []int{7, 7, 2},
			output: 7,
		},
		{
			name:   "test case 11",
			input:  []int{1, 7, 2},
			output: 2,
		},
		{
			name:   "test case 12",
			input:  []int{1, 7, 3, 2},
			output: 4,
		},
	}

	for k, s := range educative.ContainerWithMostWater {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := s(c.input)

				if res != c.output {
					t.Errorf("for input slice %v, expected %d; got %d ", c.input, c.output, res)
				}
			}
		})
	}
}
