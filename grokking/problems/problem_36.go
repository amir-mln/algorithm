package problems

/*
# Valid Anagram

Given two strings, str1 and str2, check whether str2 is an anagram of str1.
An anagram is a word or phrase created by rearranging the letters of another
word or phrase while utilizing each of the original letters exactly once.

Constraints:

 1. 1 ≤ len(str1), len(str2) ≤ 10 ** 3
 2. The strings contain lowercase English letters.

Description:

My solution involves two maps, and four linear iterations. In the first two
iterations we count the number of characters in each string. In the last two
iterations we check each map for missing or unequal characters.

The problem can also be solved using one hashmap. In this approach, we iterate
through the first string characters and increase the corresponding key-value
pair in the map. As we iterate through the second string, we decrease the
corresponding key-value pair, if it exists in the map. If it doesn't, we return
false. Finally, in one last iteration, we iterate through the map. If we found
any key that doesn't have a 0 value, we return false.
*/
type Problem36 []string

func (p Problem36) Solve() bool {
	str1, str2 := p[0], p[1]
	str1Mp, str2Mp := make(map[rune]int), make(map[rune]int)

	for _, r := range str1 {
		str1Mp[r]++
	}

	for _, r := range str2 {
		str2Mp[r]++
	}

	for k, count1 := range str1Mp {
		if count2, ok := str2Mp[k]; !ok || count1 != count2 {
			return false
		}
	}

	for k, count2 := range str2Mp {
		if count1, ok := str1Mp[k]; !ok || count1 != count2 {
			return false
		}
	}

	return true
}
