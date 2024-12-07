package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
# Sliding Window Maximum

Given an integer list, [nums], find the maximum values in all the contiguous
sub-arrays (windows) of size [w]. If the window size is greater than the array
size, we consider the entire array as a single window.

Constraints:

 1. 1 ≤ len(Problem41.Nums) ≤ 1000
 2. -10000 ≤ Problem41.Nums[i] ≤ 10000
 3. 1 ≤ Problem41.W

Description:
To solve this problem, we need a datastructure such as a queue or a stack.
We first loop through the first [W] numbers and each iteration:
 1. If the stack contains indexes of numbers that are smaller than or equal
    to the current iteration index (Numbers[i]), we pop them.
 2. As we popped the numbers that were smaller the the current number, we can
    safely push the index of current number to that stack, knowing that it
    wont contain any number that is smaller than the current number.

The bottom most element of the stack now contains the index of the largest
number of the fist [W] numbers (the largest number in first window). Then,
as we loop through the remaining numbers (starting from index [W]), we first
pop index of numbers that are smaller than current number. We then check if
the bottom most element of the stack, is an index that doesn't belong to
the current window (if bottom <= i - W). If it is, we need to remove it from
the stack. We then append number that the bottom element of stack points to
as we are certain that the bottom most element is the largest element.
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
