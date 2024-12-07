package problems

/*
# Longest Substring without Repeating Characters

Given a string, return the length of the longest substring without
repeating characters.

Constraints:

 1. 1 ≤ len(Problem42) ≤ 1000
 2. Problem42 consists of English letters, digits, and spaces.

Description:

The solution uses a sliding window approach to find the length of the
longest substring without repeating characters. It maintains a `start`
pointer to track the beginning of the current window and a map,
`indexes`, to store the last seen index of each character. As the loop
iterates through the string, each character is checked in the map. If
a character is found in the map and its last-seen index is within the
current window `idx >= start`, it indicates a repetition. In this case,
the start pointer is moved to exclude the repeated character by updating
it to `indexes[character] + 1`. The current character’s index is then
updated in the map.

If the character is in the map but its last occurrence is outside the
current window (its index is < `start`), it is treated as a new unique
character for the current window, and its index is updated without
moving the `start` pointer. After processing each character, the length
of the current substring `i - start + 1` is calculated and compared
with the maximum length found so far. By the end of the iteration, the
function returns the maximum length of any valid substring encountered.
This approach ensures that each character is processed only once, leading
to an O(n) time complexity.
*/
type Problem43 string

func (p Problem43) Solve() int {
	indexes := make(map[byte]int)
	maxSize, start := 0, 0

	for i := 0; i < len(p); i++ {
		if idx, ok := indexes[p[i]]; ok && idx >= start {
			start = idx + 1
		}
		indexes[p[i]] = i
		maxSize = max(maxSize, i-start+1)
	}

	return maxSize
}
