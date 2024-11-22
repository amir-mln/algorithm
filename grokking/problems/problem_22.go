package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
Remove Adjacent Duplicates

You are given an input string consisting of lowercase English letters. Repeatedly remove adjacent
duplicate letters, one pair at a time. Both members of a pair of adjacent duplicate letters
need to be removed.

Constraints:

	a) 1 ≤ len(input) ≤ 10**3
	b) input consists of lowercase English alphabets.
*/
var Problem22 = map[string]func(string) string{
	"stacks-solution": func(s string) string {
		st := grokking.Stack[rune]{}

		for _, ch := range s {
			if topCh, _ := st.Top(); topCh == ch {
				st.Pop()
				continue
			}

			st.Push(ch)
		}

		return string(st.Values())
	},
}
