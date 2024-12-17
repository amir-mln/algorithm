package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

/*
# Spelling Check

In this problem, we are given two strings where the first string has one more
character than the second. Our task is to determine the number of ways in which
one character can be deleted from the first string to arrive at the second string.
For example, there is one way to get from "favour" to "favor": we can remove the
`u` from the first string.

There are three ways to get from "abcdxxxef" to "abcdxxef": we can remove any of
the `x` characters from the first string.

The context for the problem is a spellchecker. The first string might be "bizzarre"
(a misspelled word) and the second might be "bizarre" (a correct spelling). In
this case, there are two ways to fix the misspelling—by removing either one of
the two `z`s from the first string. The problem is more general, though, having
nothing to do with actual English words or spelling mistakes.

Input:

The input data contains two strings, consisting of lower-case Latin letters. The
length of each string is from 1 to 1e6 symbols inclusive, the first string contains
exactly 1 symbol more than the second one.

Output:

In the first line output the number of positions of the symbols in the first string,
after the deleting of which the first string becomes identical to the second one.
In the second line output space-separated positions of these symbols in increasing
order. The positions are numbered starting from 1. If it is impossible to make the
first string identical to the second string by deleting one symbol, output 0.

Link:

https://dmoj.ca/problem/coci17c1p3hard
*/
type Problem struct {
	Str1   string
	Str2   string
	Result []int
}

func (p *Problem) Solve() {
	i, j := 0, 0
	extraIdx := -1
	for i < len(p.Str1) || j < len(p.Str2) {
		if i == len(p.Str1)-1 && extraIdx == -1 {
			extraIdx = i
			break
		}

		if p.Str1[i] == p.Str2[j] {
			i++
			j++
		} else {
			if extraIdx == -1 {
				extraIdx = i
				i++
			} else {
				p.Result = []int{}
				return
			}
		}
	}

	for i := extraIdx; i >= 0; i-- {
		if p.Str1[i] != p.Str1[extraIdx] {
			break
		}

		p.Result = append(p.Result, i+1)
	}
}

func main() {
	problem := Problem{}
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	problem.Str1 = sc.Text()
	sc.Scan()
	problem.Str2 = sc.Text()

	problem.Solve()
	if len(problem.Result) == 0 {
		fmt.Println(0)
	} else {
		slices.Sort(problem.Result)
		fmt.Println(len(problem.Result))
		for _, n := range problem.Result {
			fmt.Print(n, " ")
		}
	}
}
