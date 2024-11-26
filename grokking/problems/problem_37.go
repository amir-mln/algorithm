package problems

/*
# Anagrams in a String

Given two strings, [a] and [b], return an array of all the start indexes of anagrams
of [b] in [a]. We may return the answer in any order.

An anagram is a word or phrase created by rearranging the letters of another word or
phrase while utilizing each of the original letters exactly once. For example, “act”
is the anagram of “cat”, and “flow” is the anagram of “wolf”.

Constraints:

 1. 1 ≤ len(a), len(b) ≤ 10 ** 3
 2. The strings contain lowercase English letters.

Description:
Since the input strings contain lowercase english letters only, we can use fixed size
arrays instead of hash maps. We first iterate through the first [len(b)] characters of
the input strings and increase their count in each corresponding array. After this
iteration, if the two array are equal, we append 0 to the result slice. Then, we iterate
through the first string [a] from [len(b)]index to the end. In each iteration, we
need to count the characters in a [len(b)] size window. Since we already have measured
the count of the first [len(b)] characters of the first string in our initial iteration,
we already have processed the first window. This is why we start iterating on the first
sting from the [len(b)] character in the second iteration loop. In each iteration of
second loop, we slide the window forward. To do this, we need to first remove the first
element of the window (decrease its count in the array) and then add the new element (
increase its count int the array). The first element that we want to remove is located
at [i - len(b)] and the new element is located at [i].
*/
type Problem37 []string

func (p Problem37) Solve() []int {
	res := make([]int, 0)
	a, b := []byte(p[0]), []byte(p[1])
	arrA, arrB := [26]byte{}, [26]byte{}
	lenA, lenB := len(a), len(b)
	const byteLowA byte = 'a'

	for i := 0; i < lenB; i++ {
		arrA[a[i]-byteLowA]++
		arrB[b[i]-byteLowA]++
	}

	if arrA == arrB {
		res = append(res, 0)
	}

	for i := lenB; i < lenA; i++ {
		arrA[a[i-lenB]-byteLowA]--
		arrA[a[i]-byteLowA]++

		if arrA == arrB {
			res = append(res, i-lenB+1)
		}
	}

	return res
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
