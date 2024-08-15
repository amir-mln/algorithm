package educative

import (
	"fmt"
	"strconv"

	"github.com/amir-mln/algorithm/datastructure"
)

var p21Operators = map[string]func(int, int) int{
	"+": func(i1, i2 int) int { return i1 + i2 },
	"-": func(i1, i2 int) int { return i1 - i2 },
	"*": func(i1, i2 int) int { return i1 * i2 },
	"/": func(i1, i2 int) int { return i1 / i2 },
}

func p21RecursiveSolution(tokens []string, index *int) int {
	if *index < 0 {
		panic("index can not be negative")
	}

	var num1, num2 int
	var operator func(int, int) int
	token := tokens[*index]

	if op, ok := p21Operators[token]; ok {
		*index -= 1
		operator = op
	} else {
		panic(fmt.Sprintf("could not find operator %q", token))
	}

	token = tokens[*index]
	if _, ok := p21Operators[token]; !ok {
		num2, _ = strconv.Atoi(token)
		*index -= 1
	} else {
		num2 = p21RecursiveSolution(tokens, index)
	}

	token = tokens[*index]
	if _, ok := p21Operators[token]; !ok {
		num1, _ = strconv.Atoi(token)
		*index -= 1
	} else {
		num1 = p21RecursiveSolution(tokens, index)
	}

	return operator(num1, num2)
}

func p21StackSolution(tokens []string, _ *int) int {
	s := datastructure.Stack[int]{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if op, ok := p21Operators[token]; ok {
			n2 := s.Pop()
			n1 := s.Pop()
			s.Push(op(n1, n2))
			continue
		}

		num, _ := strconv.Atoi(token)
		s.Push(num)
	}

	return s.Top()
}

/*
Reverse Polish Notation #Stacks #Recursive

Given an arithmetic expression in a Reverse Polish Notation (RPN) in form of a slice of tokens,
your task is to evaluate and return the value of the given expression. Valid operators are +, -, *,
and /. Each operand can be an integer or another expression. The division between two integers should
truncate toward zero. The given Reverse Polish Notation expression is guaranteed to be valid. This ensures
that the expression always evaluates to a result and that there are no division-by-zero operations.

Constraints:

	a) 1 ≤ len(tokens) ≤ 10**3
	b) 'tokens'[i] is either an operator (+, -, *, or /) or an integer
		in the range [−200,200]
*/
var P21Solutions = map[string]func(tokens []string, index *int) int{
	"recursive-solution": p21RecursiveSolution,
	"stack-solution":     p21StackSolution,
}
