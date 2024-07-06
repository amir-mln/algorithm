package twopointers_test

import (
	twopointers "educative-dsa/two-pointers"
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		output bool
	}{
		{
			name:   "test case 1",
			nums:   []int{1, -1, 0},
			target: -1,
			output: false,
		},
		{
			name:   "test case 2",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 10,
			output: true,
		},
		{
			name:   "test case 3",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 21,
			output: false,
		},
		{
			name:   "test case 4",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			output: true,
		},
		{
			name:   "test case 5",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			output: true,
		},
		{
			name:   "test case 8",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 20,
			output: true,
		},
		{
			name:   "test case 9",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 21,
			output: false,
		},
		{
			name:   "test case 10",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			output: true,
		},
		{
			name:   "test case 11",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			output: true,
		},
		{
			name:   "test case 12",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 7,
			output: false,
		},
		{
			name:   "test case 13",
			nums:   []int{1, -1, 0},
			target: 1,
			output: false,
		},
		{
			name:   "test case 14",
			nums:   []int{1, -1, 0},
			target: 0,
			output: true,
		},
	}

	for k, s := range twopointers.ThreeSum {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				res := s(c.nums, c.target)

				if res != c.output {
					t.Errorf(
						"in %q with nums: %v and target %d, expected %v; got %v ",
						c.name,
						c.nums,
						c.target,
						c.output,
						res,
					)
				}
			}
		})
	}
}
