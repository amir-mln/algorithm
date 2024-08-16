package linkedlist_test

import (
	"fmt"
	"testing"

	"github.com/amir-mln/algorithm/educative/common"
	"github.com/amir-mln/algorithm/educative/linkedlist"
)

func TestReverseInKGroup(t *testing.T) {
	cases := []struct {
		name   string
		list   []int
		k      int
		output []int
	}{
		{
			name:   "test case 1",
			list:   []int{55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667},
			k:      7,
			output: []int{55, 34, 56, 434, 896, 677, 55, 32, 544, 9, 78, 455, 23, 677, 588, 553, 667},
		},
		{
			name:   "test case 2",
			list:   []int{78, 78, 9, 544, 32, 588, 553, 667, 2, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 455, 30, 789, 32, 68, 45, 78, 9, 544, 32, 588, 553, 667, 2, 4, 553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 3, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 666, 12, 3, 44, 56, 355, 567, 1, 2, 3, 4, 5, 356, 900, 78, 43, 23, 56, 34, 55, 544, 999, 345, 612},
			k:      100,
			output: []int{681, 93, 68, 32, 789, 32, 45, 78, 32, 85, 56, 24, 66, 445, 3, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 2, 667, 553, 6, 5, 4, 445, 3, 4, 2, 667, 553, 4, 2, 667, 553, 588, 32, 544, 9, 78, 45, 68, 32, 789, 30, 455, 35, 56, 78, 56, 444, 787, 56, 43, 90, 67, 35, 2, 1, 0, 434, 896, 2, 667, 553, 588, 32, 544, 9, 78, 78, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 666, 12, 3, 44, 56, 355, 567, 1, 2, 3, 4, 5, 356, 900, 78, 43, 23, 56, 34, 55, 544, 999, 345, 612},
		},
		{
			name:   "test case 3",
			list:   []int{1},
			k:      1,
			output: []int{1},
		},
		{
			name:   "test case 4",
			list:   []int{2, 3, 4, 5, 6, 2, 3, 8, 90, 4, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 4, 3, 445, 66, 78, 43, 23, 56, 34, 55, 677, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 45, 93, 68, 45, 89, 444, 788, 534, 789, 46, 6, 666, 355, 567, 356, 900, 544, 999, 345, 612},
			k:      50,
			output: []int{85, 56, 24, 35, 56, 78, 56, 444, 787, 56, 43, 90, 67, 35, 2, 1, 0, 434, 896, 677, 55, 34, 56, 23, 43, 78, 66, 445, 3, 4, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 4, 90, 8, 3, 2, 6, 5, 4, 3, 2, 32, 78, 45, 32, 789, 32, 68, 45, 93, 68, 45, 89, 444, 788, 534, 789, 46, 6, 666, 355, 567, 356, 900, 544, 999, 345, 612},
		},
		{
			name:   "test case 5",
			list:   []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			k:      4,
			output: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			name:   "test case 6",
			list:   []int{2, 4, 6, 7, 8},
			k:      5,
			output: []int{8, 7, 6, 4, 2},
		},
		{
			name:   "test case 7",
			list:   []int{6, 7, 7, 8, 0, 1, 2, 5, 0, 1, 6, 8, 3, 9, 1, 0, 4, 2, 7, 9, 3, 6},
			k:      10,
			output: []int{1, 0, 5, 2, 1, 0, 8, 7, 7, 6, 9, 7, 2, 4, 0, 1, 9, 3, 8, 6, 3, 6},
		},
		{
			name:   "test case 8",
			list:   []int{553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 455, 30, 789, 32, 68, 45, 78, 9, 544, 32, 588, 553, 667, 2, 4, 553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 3, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 45, 32, 444, 788, 534, 16, 34, 68, 32, 13, 789, 46, 6, 60, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 455, 30, 789, 32, 68, 45, 78, 9, 544, 32, 588, 553, 667, 2, 4, 553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 3, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 78, 9, 544, 32, 588, 553, 667, 2, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 455, 30, 789, 32, 68, 45, 78, 9, 544, 32, 588, 553, 667, 2, 4, 553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 3, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 666, 12, 3, 44, 56, 355, 567, 1, 2, 3, 4, 5, 356, 900, 78, 43, 23, 56, 34, 55, 544, 999, 345, 612},
			k:      300,
			output: []int{3, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 2, 667, 553, 6, 5, 4, 445, 3, 4, 2, 667, 553, 4, 2, 667, 553, 588, 32, 544, 9, 78, 45, 68, 32, 789, 30, 455, 35, 56, 78, 56, 444, 787, 56, 43, 90, 67, 35, 2, 1, 0, 434, 896, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 2, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 60, 6, 46, 789, 13, 32, 68, 34, 16, 534, 788, 444, 32, 45, 78, 32, 85, 56, 24, 35, 89, 45, 1, 78, 5, 11, 68, 23, 681, 93, 68, 32, 789, 32, 45, 78, 32, 85, 56, 24, 66, 445, 3, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 2, 667, 553, 6, 5, 4, 445, 3, 4, 2, 667, 553, 4, 2, 667, 553, 588, 32, 544, 9, 78, 45, 68, 32, 789, 30, 455, 35, 56, 78, 56, 444, 787, 56, 43, 90, 67, 35, 2, 1, 0, 434, 896, 2, 667, 553, 588, 32, 544, 9, 78, 455, 23, 677, 55, 34, 56, 434, 896, 677, 55, 23, 43, 78, 66, 9, 78, 455, 23, 4, 90, 8, 3, 2, 3, 4, 2, 667, 553, 6, 5, 4, 445, 3, 4, 2, 667, 553, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 78, 9, 544, 32, 588, 553, 667, 2, 896, 434, 0, 1, 2, 35, 67, 90, 43, 56, 787, 444, 56, 78, 56, 35, 455, 30, 789, 32, 68, 45, 78, 9, 544, 32, 588, 553, 667, 2, 4, 553, 667, 2, 4, 3, 445, 4, 5, 6, 553, 667, 2, 4, 3, 2, 3, 8, 90, 4, 23, 455, 78, 9, 66, 78, 43, 23, 55, 677, 896, 434, 56, 34, 55, 677, 23, 455, 78, 9, 544, 32, 588, 553, 667, 2, 3, 445, 66, 24, 56, 85, 32, 78, 45, 32, 789, 32, 68, 93, 681, 23, 68, 11, 5, 78, 1, 45, 89, 35, 24, 56, 85, 32, 78, 666, 12, 3, 44, 56, 355, 567, 1, 2, 3, 4, 5, 356, 900, 78, 43, 23, 56, 34, 55, 544, 999, 345, 612},
		},
		{
			name:   "test case 9",
			list:   []int{55, 677, 55, 677, 3, 9, 1, 0, 4, 2, 7, 9, 896, 434, 56, 34, 3, 6, 23, 455, 5, 2, 1, 0, 8, 78, 9, 544, 32, 588, 553, 667},
			k:      2,
			output: []int{677, 55, 677, 55, 9, 3, 0, 1, 2, 4, 9, 7, 434, 896, 34, 56, 6, 3, 455, 23, 2, 5, 0, 1, 78, 8, 544, 9, 588, 32, 667, 553},
		},
		{
			name:   "test case 10",
			list:   []int{55, 677, 2, 4, 3, 2, 3, 8, 55, 677, 3, 9, 1, 0, 4, 2, 7, 9, 896, 434, 56, 34, 3, 6, 23, 455, 5, 2, 1, 0, 8, 78, 9, 544, 32, 588, 55, 23, 43, 78, 66, 9, 78, 553, 667},
			k:      45,
			output: []int{667, 553, 78, 9, 66, 78, 43, 23, 55, 588, 32, 544, 9, 78, 8, 0, 1, 2, 5, 455, 23, 6, 3, 34, 56, 434, 896, 9, 7, 2, 4, 0, 1, 9, 3, 677, 55, 8, 3, 2, 3, 4, 2, 677, 55},
		},
		{
			name:   "test case 11",
			list:   []int{1, 2, 3, 4, 5},
			k:      2,
			output: []int{2, 1, 4, 3, 5},
		},
		{
			name:   "test case 12",
			list:   []int{3, 4, 5, 6, 2, 8, 7, 7},
			k:      3,
			output: []int{5, 4, 3, 8, 2, 6, 7, 7},
		},
		{
			name:   "test case 13",
			list:   []int{1, 2, 3, 4, 5},
			k:      1,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "test case 14",
			list:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			output: []int{3, 2, 1, 6, 5, 4, 7},
		},
		{
			name:   "test case 15",
			list:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      7,
			output: []int{7, 6, 5, 4, 3, 2, 1},
		},
	}

	for k, f := range linkedlist.ReverseInKGroup {
		t.Run(k, func(t *testing.T) {
			for _, c := range cases {
				l := &common.LinkedList{}
				l.InsertFromList(c.list)
				f(l, c.k)

				if fmt.Sprint(l) != fmt.Sprint(c.output) {
					t.Errorf(
						"in %q with list %v and k %d, expected %v; got %v ",
						c.name,
						c.list,
						c.k,
						c.output,
						l,
					)
				}
			}
		})

	}
}