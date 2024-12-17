package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
# Login Mayhem

Imagine that you are wanting to join a (hopefully theoretical) social network
website that has a major security concern: other passwords besides yours can
be used to get into your account! Specifically, if someone tries a password
that has your password as a substring, then they’re in. If your password were
"dish", for example, then passwords like "brandish" and "radishes" would work
to get into your account because the string "dish" is in them. You don’t know
what password to choose for your account so at various points you will ask:
“If I chose this password, how many current users’ passwords would get in to
my account?”

We need to support two types of operations:
  - Add Sign up a new user with the given password.
  - Query Given a proposed password p, return the number of current
    users’ passwords that could be used to get into an account whose pass-
    word is p.

Input:

The input consists of the following lines:

  - A line containing q, the number of operations to be performed. q is
    between 1 and 100,000.
  - q lines, each giving one add or query operation to be performed.

Here are the operations that can be performed in those q lines:

  - An add operation is specified as the number 1, a space, and then the
    new user’s password. It indicates that a new user has joined with the
    provided password. This operation doesn’t result in any output.

  - A query operation is specified as the number 2, a space, and then a
    proposed password p. It indicates that we should output the num-
    ber of current users’ passwords that could be used to get into an
    account whose password is p.

All passwords provided in these operations are between 1 and 10 lower-
case characters.

Output:

Output the result of each query operation, one per line.
The time limit for solving the test case is three seconds.

Link:

https://dmoj.ca/problem/coci17c1p3hard
*/
type Problem struct {
	passwords map[string]int
}

func (p *Problem) Add(word string) {
	if p.passwords == nil {
		p.passwords = make(map[string]int)
	}

	set := make(map[string]struct{})
	for i := 1; i <= len(word); i++ {
		for j := 0; j <= len(word)-i; j++ {
			set[word[j:j+i]] = struct{}{}
		}
	}

	for k := range set {
		p.passwords[k]++
	}
}

func (p *Problem) Query(word string) int {
	if p.passwords == nil {
		p.passwords = make(map[string]int)
	}

	return p.passwords[word]
}

func main() {
	problem := Problem{}
	out := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		txt := strings.Fields(sc.Text())

		if txt[0] == "1" {
			problem.Add(txt[1])
		} else {
			out.WriteString(fmt.Sprintln(problem.Query(txt[1])))
		}
	}

	out.Flush()
}
