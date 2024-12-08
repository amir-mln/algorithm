package problems

/*
# Minimum Window Substring

Given two strings, `str1` and `str2`, find the shortest substring in `str1` such that
`str2` is a subsequence of that substring. For `str2` to be the subsequence of this
shortest substring of `str1`, the substring must contain all the characters of `str2`
or more. For example, if `str1` is `ABBECBA` and `str2` is `AC`, the shortest substring
of which `str2` is a subsequence, would be `cba`.

If there is no substring in `str1` that covers all characters in `str2`, return an
empty string. If there are multiple minimum-length substrings that meet the subsequence
requirement, return the one with the left-most starting index (the first one).

Constraints:

 1. 1 ≤ len(str1) ≤ 2000
 2. 1 ≤ len(str2) ≤ 100
 3. `str1` and `str2` consist only of uppercase english letters.

Description:

The goal is to find the smallest substring in `str1` that contains all the
characters of `str2`. To achieve this, two fixed-size arrays, `freq1` and
`freq2`, are used to track the character frequencies of `str1` and `str2`,
respectively. The `freq2` array is initialized to store the counts of
characters in str2.

Using a two-pointer technique, the code iterates through `str1` with a `start`
and `end` pointer. The `end` pointer expands the window by including characters,
while the `start` pointer contracts it when the window already satisfies the
condition of containing all characters in `str2`. The helper function `isSubstring`
verifies whether the current window meets this condition by comparing the character
frequencies in `freq1` against `freq2`. As soon as the window satisfies the condition,
the substring is checked, and the `start` pointer is incremented to potentially
find a smaller valid window. The process continues until the end of the string
is reached, and the smallest valid substring is returned.
*/
type Problem45 [2]string

func (p Problem45) Solve() string {
	str1, str2 := p[0], p[1]
	freq1, freq2 := [26]int{}, [26]int{}
	start, end := 0, 0
	sub := ""
	isSubstring := func() bool {
		for i, c := range freq2 {
			if freq1[i] < c {
				return false
			}
		}

		return true
	}

	for _, r := range str2 {
		freq2[r-'A']++
	}

	for end < len(str1) {
		ch := str1[end]
		freq1[ch-'A']++

		for isSubstring() && start <= end {
			sub = str1[start : end+1]
			if len(sub) == len(str2) {
				return sub
			}

			freq1[str1[start]-'A']--
			start++

		}

		end++
	}

	return sub
}
