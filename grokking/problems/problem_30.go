package problems

import "github.com/amir-mln/algorithm/grokking"

/*
# Next Greater Element

Given the two distinct integer arrays, nums1 and nums2, where
nums1 is a subset of nums2, find all the next greater elements
for nums1 values in the corresponding places of nums2.

In general, the next greater element of an element, x, in an array
is the first greater element present on the right side of x in the
same array. However, in the context of this problem, for each element
x in nums1, find the next greater element present on the right side
of x in nums2 and store it in the ans array. If there is no such element,
store −1 for this number. The ans array should be of the same length
as nums1, and the order of the elements in the ans array should correspond
to the order of the elements in nums1.

Constraints:

 1. 1 ≤ len(nums1) ≤ len(num2) ≤ 1000
 2. 1 ≤ nums1[i], nums2[i] ≤ 10_000
 3. Each nums1 and nums2 have distinct integers.
 4. All integers in nums1 also appear in nums2.

Notes:
 1. The way we use stack in this problem is similar to how we use it in
    problem 25
*/
type Problem30 struct {
	Nums1 []int
	Nums2 []int
}

func (p Problem30) Solve() []int {
	stck := grokking.Stack[int]{}
	dict := make(map[int]int)

	for _, v := range p.Nums2 {
		for t, ok := stck.Top(); ok && v > t; _, ok = stck.Pop() {
			dict[t] = v
		}

		stck.Push(v)
	}

	res := make([]int, len(p.Nums1))
	for i, v := range p.Nums1 {
		if gte, ok := dict[v]; ok {
			res[i] = gte
		} else {
			res[i] = -1
		}
	}

	return res
}
