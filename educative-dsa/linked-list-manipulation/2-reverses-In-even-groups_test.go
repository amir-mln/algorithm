package linkedlistmanipulation_test

import (
	"educative-dsa/common"
	linkedlistmanipulation "educative-dsa/linked-list-manipulation"
	"fmt"
	"testing"
)

func TestReverseInEvenGroups(t *testing.T) {
	cases := []struct {
		name   string
		list   []int
		output []int
	}{
		{
			name:   "test case 1",
			list:   []int{1, -2, 3, 4, -5, 6},
			output: []int{1, 3, -2, 6, -5, 4},
		},
		{
			name:   "test case 2",
			list:   []int{9},
			output: []int{9},
		},
		{
			name:   "test case 3",
			list:   []int{1, 2, 3, 4},
			output: []int{1, 3, 2, 4},
		},
		{
			name:   "test case 4",
			list:   []int{10, 1, 2, 3, 4, 5},
			output: []int{10, 2, 1, 5, 4, 3},
		},
		{
			name:   "test case 5",
			list:   []int{28, 21, 14, 7},
			output: []int{28, 14, 21, 7},
		},
		{
			name:   "test case 6",
			list:   []int{11, 12, 13, 14, 15},
			output: []int{11, 13, 12, 15, 14},
		},
		{
			name:   "test case 7",
			list:   []int{1, 2},
			output: []int{1, 2},
		},
	}

	for k, f := range linkedlistmanipulation.ReverseInEvenGroups {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &common.LinkedList{}
				l.InsertFromList(c.list)
				f(l)

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
