package problems

/*
# Contains Duplicate

For a given array of integers, [nums], return true if it contains duplicates.
Otherwise, return false.

Constraints:

 1. 1 ≤ len(nums) ≤ 10 ** 3
 2. −10**9 ≤ nums[i] ≤ 10**9
*/
type Problem34 struct {
	Nums []int
}

func (p Problem34) Solve() bool {
	mp := make(map[int]struct{})
	for _, num := range p.Nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = struct{}{}
	}

	return false
}
