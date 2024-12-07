package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
# Sliding Window Maximum

Given an integer list, [nums], find the maximum values in all the
contiguous sub-arrays (windows) of size [w]. If the window size is
greater than the array size, we consider the entire array as a
single window.

Constraints:

 1. 1 ≤ len(nums) ≤ 1000
 2. -10000 ≤ nums[i] ≤ 10000
 2. DNA[i] is either A, C, G, and T.
 3. 1 ≤ w
*/

type Problem41 struct {
	Nums []int
	W    int
}

func (p Problem41) Solve() []int {
	window := &grokking.Stack[int]{}
	popSmaller := func(i int) {
		for t, ok := window.Top(); ok && p.Nums[i] >= p.Nums[t]; t, ok = window.Top() {
			window.Pop()
		}
	}

	for i := 0; i < p.W; i++ {
		popSmaller(i)
		window.Push(i)
	}
	b, _ := window.Bottom()
	result := make([]int, 0)
	result = append(result, p.Nums[b])

	for i := p.W; i < len(p.Nums); i++ {
		popSmaller(i)
		if b, ok := window.Bottom(); ok && b <= i-p.W {
			window.PopBottom()
		}
		window.Push(i)
		b, _ := window.Bottom()
		result = append(result, p.Nums[b])
	}

	return result
}
