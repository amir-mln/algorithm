package twopointers_test

import (
	common "educative-dsa/common"
	twopointers "educative-dsa/two-pointers"
	"fmt"
	"testing"
)

func TestRemoveNthNode(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		list   []int
		output []int
	}{
		{
			name:   "test case 1",
			n:      2,
			list:   []int{23, 28, 10, 5, 67, 39, 70, 28},
			output: []int{23, 28, 10, 5, 67, 39, 28},
		},
		{
			name:   "test case 2",
			list:   []int{34, 53, 6, 95, 38, 28, 17, 63, 16, 76},
			n:      3,
			output: []int{34, 53, 6, 95, 38, 28, 17, 16, 76},
		},
		{
			name:   "test case 3",
			list:   []int{288, 224, 275, 390, 4, 383, 330, 60, 193},
			n:      4,
			output: []int{288, 224, 275, 390, 4, 330, 60, 193},
		},
		{
			name:   "test case 4",
			list:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			n:      1,
			output: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:   "test case 5",
			list:   []int{69, 8, 49, 106, 116, 112},
			n:      6,
			output: []int{8, 49, 106, 116, 112},
		},
		{
			name:   "test case 6",
			list:   []int{40, 43, 38, 195, 30, 152, 164, 162, 47, 20, 116, 160, 36, 118, 114, 148, 135, 29, 133, 51},
			n:      10,
			output: []int{40, 43, 38, 195, 30, 152, 164, 162, 47, 20, 160, 36, 118, 114, 148, 135, 29, 133, 51},
		},
		{
			name:   "test case 7",
			list:   []int{763, 92, 100, 695, 47, 543, 349, 550, 498, 328},
			n:      4,
			output: []int{763, 92, 100, 695, 47, 543, 550, 498, 328},
		},
		{
			name:   "test case 8",
			list:   []int{27, 7, 26, 22, 44, 2, 5, 9, 36, 4, 12, 8, 38, 48, 11},
			n:      15,
			output: []int{7, 26, 22, 44, 2, 5, 9, 36, 4, 12, 8, 38, 48, 11},
		},
		{
			name:   "test case 9",
			list:   []int{0, 91, 36, 78, 33, 6, 93, 36},
			n:      1,
			output: []int{0, 91, 36, 78, 33, 6, 93},
		},
		{
			name:   "test case 10",
			list:   []int{69, 8, 49, 106, 116, 112},
			n:      5,
			output: []int{69, 49, 106, 116, 112},
		},
	}

	for k, v := range twopointers.RemoveNthNode {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &common.LinkedList{}
				l.InsertFromList(c.list)

				v(l, c.n)

				got := fmt.Sprint(l)
				expected := fmt.Sprint(c.output)

				if expected != got {
					t.Errorf("case %q, expected %v; got %v", c.name, expected, got)
				}
			}
		})
	}
}
