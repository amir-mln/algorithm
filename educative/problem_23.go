package educative

import (
	"github.com/amir-mln/algorithm/datastructure"
)

/*
Valid Parentheses

Given a string that may consist of opening and closing parentheses, your task is to
check whether or not the string contains valid parenthesis. Every opening parenthesis
should be closed by the same kind of parenthesis. Therefore, '{)' and '[(])' strings
are invalid. Every opening parenthesis must be closed in the correct order. Therefore,
')(' and '()(()' are invalid.

Constraints:

	a) 1 ≤ len(input) ≤ 10**3
	b) input consists of following characters only: '(', ')', '[', ']', '{' and '}'.
*/
var P23Solutions = map[string]func(string) bool{
	"stacks-solution": func(s string) bool {
		st := datastructure.Stack[rune]{}
		closings := map[rune]rune{
			')': '(',
			'}': '{',
			']': '[',
		}

		for _, ch := range s {
			if open, ok := closings[ch]; ok {
				if st.Top() != open {
					return false
				}

				st.Pop()
				continue
			}

			st.Push(ch)
		}

		return st.Empty()
	},
}
