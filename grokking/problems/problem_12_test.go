package problems_test

import (
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestMiddleOfLinkedList(t *testing.T) {
	cases := []struct {
		name   string
		list   []int
		output int
	}{
		{
			name:   "test case 1",
			list:   []int{1, 2, 3, 4, 5},
			output: 3,
		},
		{
			name:   "test case 2",
			list:   []int{1, 2, 3, 4, 5, 6},
			output: 4,
		},
		{
			name:   "test case 3",
			list:   []int{3, 2, 1},
			output: 2,
		},
		{
			name:   "test case 4",
			list:   []int{10},
			output: 10,
		},
		{
			name:   "test case 5",
			list:   []int{1, 2},
			output: 2,
		},
		{
			name:   "test case 6",
			list:   []int{1, 2, 8, 3, 4},
			output: 8,
		},
	}

	for k, f := range problems.MiddleOfLinkedList {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &grokking.LinkedList{}
				l.InsertFromList(c.list)

				o := f(l.Head)

				if o.Value != c.output {
					t.Errorf(
						"in %q with list %v, expected %v; got %v ",
						c.name,
						c.list,
						c.output,
						o.Value,
					)
				}
			}
		})

	}
}
