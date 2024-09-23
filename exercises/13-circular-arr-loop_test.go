package exercises_test

import (
	"testing"

	"github.com/amir-mln/algorithm/exercises"
)

func TestCircularArrayLoop(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output bool
	}{
		{
			name:   "test case 1",
			input:  []int{1, 3, -2, -4, 1},
			output: true,
		},
		{
			name:   "test case 2",
			input:  []int{2, 1, -1, -2},
			output: false,
		},
		{
			name:   "test case 3",
			input:  []int{5, 4, -2, -1, 3},
			output: false,
		},
		{
			name:   "test case 4",
			input:  []int{1, 2, -3, 3, 4, 7, 1},
			output: true,
		},
		{
			name:   "test case 5",
			input:  []int{3, 3, 1, -1, 2},
			output: true,
		},
		{
			name:   "test case 6",
			input:  []int{2, 2, 2, 7, 2, -1, 2, -1, -1},
			output: true,
		},
		{
			name:   "test case 7",
			input:  []int{1, -1, 2, -2},
			output: false,
		},
		{
			name:   "test case 8",
			input:  []int{3, 2, 1, 1, -4, 1},
			output: false,
		},
		{
			name:   "test case 9",
			input:  []int{-1, 2, -2, 5, 3, -6, -1, 1, 3, 7, -4, 1, 4, 6, 8},
			output: true,
		},
		{
			name:   "test case 10",
			input:  []int{2, -1, 1, 4, 3, -2, -4, -1, -5, -3, 7, 6, 2, -8, 1, 4, 9, -7, -6, -2},
			output: false,
		},
	}

	for k, f := range exercises.CircularArrayLoop {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				o := f(c.input)

				if o != c.output {
					t.Errorf(
						"in %q with nums %v, expected %v; got %v ",
						c.name,
						c.input,
						c.output,
						o,
					)
				}
			}
		})

	}
}
