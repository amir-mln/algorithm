package problems_test

import (
	"testing"

	"github.com/amir-mln/algorithm/grokking"
	"github.com/amir-mln/algorithm/grokking/problems"
)

func TestListCycle(t *testing.T) {
	cases := []struct {
		name        string
		list        []int
		loopedIndex int
		output      bool
	}{
		{
			name:        "test case 1",
			list:        []int{2, 4, 6, 8, 10},
			loopedIndex: 2,
			output:      true,
		},
		{
			name:        "test case 2",
			list:        []int{1, 3, 5, 7, 9},
			loopedIndex: -1,
			output:      false,
		},
		{
			name:        "test case 3",
			list:        []int{1, 2, 3, 4, 5},
			loopedIndex: 3,
			output:      true,
		},
		{
			name:        "test case 4",
			list:        []int{0, 2, 3, 5, 6},
			loopedIndex: -1,
			output:      false,
		},
		{
			name:        "test case 5",
			list:        []int{3, 6, 9, 10, 11},
			loopedIndex: 0,
			output:      true,
		},
		{
			name:        "test case 6",
			list:        []int{2, 2, 4, 5, 6, 3, 2, 4},
			loopedIndex: 1,
			output:      true,
		},
		{
			name:        "test case 7",
			list:        []int{2, 3, 5, 6, 7, 2, 4, 2},
			loopedIndex: 2,
			output:      true,
		},
		{
			name:        "test case 8",
			list:        []int{4, 4, 4, 4, 4, 4, 4},
			loopedIndex: -1,
			output:      false,
		},
		{
			name:        "test case 9",
			list:        []int{0, 1, 2, 3, 4},
			loopedIndex: 0,
			output:      true,
		},
		{
			name:        "test case 10",
			list:        []int{11, 31, 35, 6, 19},
			loopedIndex: 1,
			output:      true,
		},
	}

	for k, f := range problems.LinkedListCycle {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &grokking.LinkedList{}
				l.InsertFromList(c.list)

				n := l.Head
				var loopedNode *grokking.LinkedListNode
				for n.Next != nil {
					if c.loopedIndex == 0 {
						loopedNode = n
					}
					c.loopedIndex--
					n = n.Next
				}
				n.Next = loopedNode

				o := f(l)

				if o != c.output {
					t.Errorf(
						"in %q with list %v, expected %v; got %v ",
						c.name,
						c.list,
						c.output,
						o,
					)
				}
			}
		})

	}
}
