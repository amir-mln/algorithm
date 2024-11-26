package problems

/*
# Two Sum
For the given array of integers and a target, you have to identify the two
indices that add up to generate the target. Moreover, you can’t use the same
index twice, and there will be only one solution. The order of output indices
don't matter.

Constraints:

 1. 1 ≤ len(array) ≤ 10 ** 3
 2. -10 ** 9 ≤ array[i] ≤ 10 ** 9
 3. -10 ** 9 ≤ target ≤ 10 ** 9
 2. Only one valid answer exists.

Description:

We simply need a hash map to track the numbers that exist in the array. We
then loop through the array to see if the map has any key that equals to
[target - array[i]].
*/
type Problem38 struct {
	Numbers []int
	Target  int
}

func (p Problem38) Solve() []int {
	mp := make(map[int]int)
	for i, num := range p.Numbers {
		mp[num] = i
	}

	for i, num := range p.Numbers {
		if j, ok := mp[p.Target-num]; ok {
			return []int{i, j}
		}
	}

	return []int{}
}

// func (p Problem37) Solve() []int {
// 	a, b := []byte(p[0]), []byte(p[1])
// 	mp := make(map[byte]int)
// 	res := make([]int, 0)
// 	for _, r := range b {
// 		mp[r]++
// 	}
// 	slices.Sort(b)
// 	for i := 0; i <= len(a)-len(b); i++ {
// 		char := a[i]
// 		if _, ok := mp[char]; ok {
// 			sub := bytes.Clone(a[i : i+len(b)])
// 			slices.Sort(sub)
// 			if slices.Equal(sub, b) {
// 				res = append(res, i)
// 			}
// 		}
// 	}
// 	return res
// }
