package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/amir-mln/algorithm/datastructure"
	"github.com/amir-mln/algorithm/educative/linkedlist"
)

func TestCircularArrayLoop(t *testing.T) {
	cases := []struct {
		name   string
		list   []int
		output []int
	}{
		{
			name:   "test case 1",
			list:   []int{1, -2, 3, 4, -5, 4, 3, -2, 1},
			output: []int{1, -2, 3, 4, -5, 4, 3, -2, 1},
		},
		{
			name:   "test case 2",
			list:   []int{-1, -5, -3, -7, -8, -6, -2},
			output: []int{-2, -6, -8, -7, -3, -5, -1},
		},
		{
			name:   "test case 3",
			list:   []int{-1, 2, -3, 4},
			output: []int{4, -3, 2, -1},
		},
		{
			name:   "test case 4",
			list:   []int{1, -1, -2, 3, -4, 5},
			output: []int{5, -4, 3, -2, -1, 1},
		},
		{
			name:   "test case 5",
			list:   []int{28, 21, 14, 7},
			output: []int{7, 14, 21, 28},
		},
		{
			name:   "test case 6",
			list:   []int{11, -12, 13, -14},
			output: []int{-14, 13, -12, 11},
		},
		{
			name:   "test case 7",
			list:   []int{-10},
			output: []int{-10},
		},
		{
			name:   "test case 8",
			list:   []int{-55, -677, -896, -434, -56, -34, -55, -677, -23, -455, -78, -9, -544, -32, -588, -553, -667},
			output: []int{-667, -553, -588, -32, -544, -9, -78, -455, -23, -677, -55, -34, -56, -434, -896, -677, -55},
		},
	}

	for k, f := range linkedlist.ReverseLinkedList {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &datastructure.LinkedList{}
				l.InsertFromList(c.list)
				f(&l.Head)

				if fmt.Sprint(l) != fmt.Sprint(c.output) {
					t.Errorf(
						"in %q with list %v, expected %v; got %v ",
						c.name,
						c.list,
						c.output,
						l,
					)
				}
			}
		})

	}
}
