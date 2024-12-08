package problems

/*
# Longest Repeating Character Replacement

Given a string and an integer, find the length of the longest substring
in the string, where all characters are identical, after replacing, at
most, as many characters as the input integer with any other lowercase
english character.

Constraints:

 1. 1 ≤ len(string) ≤ 1000
 2. 0 ≤ integer ≤ len(string)
 3. string consists only of lower case english letters.

Description:

The function uses a sliding window approach with two pointers, `start`
and `end`, to maintain a window of characters in the given string.
It also uses an array, `chars`, to keep track of the frequency of each
character within the current window. The `majority“ variable tracks
the count of the most frequently occurring character in the current window.

The core of the solution is a loop that iterates through the string,
expanding the window by moving the end pointer. For each new character
added to the window, its frequency is updated in `chars`, and the
`majority` variable is updated to reflect the highest frequency of any
character in the window. The number of characters in the current window
that we'd need to replace to make all characters the same is calculated
by the `end-start+1-majority` expression. If this number exceeds the
amount that is given by the problem (`C`), the function shrinks the
window by moving the `start` pointer to the right and decrementing the
frequency of the character at the `start` position.

At each step, the function calculates the maximum length of a valid
substring (`maxSub“) by comparing the current window size with the
previously recorded maximum. The final result is returned as `maxSub`.
*/
type Problem44 struct {
	S string
	C int
}

func (p Problem44) Solve() int {
	chars := [26]int{}
	maxSub := 0
	start := 0
	majority := 0

	for end, r := range p.S {
		chars[r-'a']++
		currentCount := chars[r-'a']
		majority = max(majority, currentCount)

		if end-start+1-majority > p.C {
			chars[p.S[start]-'a']--
			start++
		}

		maxSub = max(end-start+1, maxSub)
	}

	return maxSub
}
