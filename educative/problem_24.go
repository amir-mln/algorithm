package educative

import (
	"github.com/amir-mln/algorithm/datastructure"
)

/*
Valid Parentheses

Given a string, s, that may have matched and unmatched parentheses, remove the minimum
number of parentheses so that the resulting string contains valid parenthesis. For example,
the string “a(b)” contains valid parenthesis while the string “a(b” doesn’t.

Constraints:

	a) 1 ≤ len(input) ≤ 10**3
	b) input consists of '(', ')', and lowercase english letters only.
*/
var P24Solutions = map[string]func(string) string{
	"stacks-solution": func(s string) string {
		res := ""
		st := datastructure.Stack[rune]{}
		const open, close = '(', ')'

		for _, ch := range s {
			if ch == close {
				if st.Top() != open {
					continue
				}

				st.Pop()
			}

			if ch == open {
				st.Push(ch)
			}

			res += string(ch)
		}

		// after the first iteration, the 'result' string may still contain invalid open parenthesis
		temp := ""
		st = datastructure.Stack[rune]{}
		for i := len(res) - 1; i >= 0; i-- {
			ch := rune(res[i])

			if ch == close {
				st.Push(ch)
			}

			if ch == open {
				if st.Top() != close {
					continue
				}

				st.Pop()
			}

			temp = string(ch) + temp
		}

		res = temp

		return res
	},
}
