package problems

/*
# Palindrome Permutation

For a given string,find whether or not a permutation of this
string is a palindrome. You should return true if such a permutation
is possible and false if it isn’t possible.

Constraints:

 1. 1 ≤ len(string) ≤ 10 ** 3
 2. The string will contain lowercase English letters.
*/
type Problem35 string

func (p Problem35) Solve() bool {
	chars := make(map[rune]int)

	for _, r := range p {
		chars[r]++
	}

	hasOdd := false
	for _, v := range chars {
		if v%2 != 0 {
			if hasOdd {
				return false
			}
			hasOdd = true
		}
	}

	return true
}
