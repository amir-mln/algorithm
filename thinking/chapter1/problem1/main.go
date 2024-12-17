package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
# Unique Snowflakes

We’re given a collection of snowflakes, and we have to determine whether
any of the snowflakes in the collection are identical. A snowflake is
represented by six integers, where each integer gives the length of one
of the snowflake’s arms. For example, this is a snowflake:

	3, 9, 15, 2, 1, 10

Snowflakes can also have repeated integers, such as

	8, 4, 8, 9, 2, 8

What does it mean for two snowflakes to be identical? Let’s work up to
that definition through a few examples. First, we’ll look at these two
snowflakes:

	1, 2, 3, 4, 5, 6	4, 5, 6, 1, 2, 3

These are also identical. We can see this by starting at the 1 in the second
snowflake and moving right. We see the integers 1, 2, and 3 and then, wrap-
ping around to the left, we see 4, 5, and 6. These two pieces together give us
the first snowflake.

Let’s try a different kind of example:

	1, 2, 3, 4, 5, 6	3, 2, 1, 6, 5, 4

If we begin at the 1 in the second snowflake and move left instead of right,
then we do get exactly 1, 2, 3, 4, 5, 6! Moving left from the 1 gives us
1, 2, 3, and wrapping around to the right, we can proceed leftward to collect
4, 5, 6.

Putting it all together, we can conclude that two snowflakes are identical
if they are the same, if we can make them the same by moving rightward
through one of the snowflakes (moving clockwise), or if we can make them
the same by moving leftward through one of the snowflakes (moving
counter clockwise).

Input:

The first line of input is an integer n, the number of snowflakes that we’ll be
processing. The value n will be between 1 and 100,000. Each of the follow-
ing n lines represents one snowflake: each line has six integers, where each
integer is at least 0 and at most 10,000,000.

Output:

Our output will be a single line of text:

  - If there are no identical snowflakes, output exactly No two snowflakes
    are alike.
  - If there are at least two identical snowflakes, output exactly Twin
    snowflakes found.

The time limit for solving the test cases is one second.

Link:

https://dmoj.ca/problem/cco07p2
*/
type Problem struct {
	Result     bool
	Snowflakes [][6]int
}

func (p *Problem) Solve() {
	const size = 6
	mp := make(map[int][][size]int)
	sum := func(s [size]int) int {
		sum := 0
		for _, num := range s {
			sum += num
		}
		return sum
	}
	areIdentical := func(s1, s2 [size]int) bool {
		for i := 0; i < size; i++ {
			identical := true
			// moving right to see if they are identical
			for j := 0; j < 6; j++ {
				idx := (i + j) % 6

				if s1[idx] != s2[j] {
					identical = false
					break
				}
			}

			if identical {
				return true
			}

			identical = true
			// moving left to see if they are identical
			for j := 0; j < 6; j++ {
				idx := i - j
				if idx < 0 {
					idx = idx + size
				}

				if s1[j] != s2[idx] {
					identical = false
					break
				}
			}

			if identical {
				return true
			}
		}

		return false
	}

	for _, snowflake := range p.Snowflakes {
		s := sum(snowflake)
		// we're putting the snowflakes that have the same sum amount into
		// the same group. this might put snowflakes that are not equal but
		// have the same sum amount into the same group, but we're ok with it.
		mp[s] = append(mp[s], snowflake)
	}

	for _, v := range mp {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				if areIdentical(v[i], v[j]) {
					p.Result = true
					return
				}
			}
		}
	}
}

func main() {
	problem := Problem{}
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		txt := strings.Fields(sc.Text())
		nums := [6]int{}
		for i, s := range txt {
			nums[i], _ = strconv.Atoi(s)
		}
		problem.Snowflakes = append(problem.Snowflakes, nums)
	}

	problem.Solve()
	if problem.Result {
		fmt.Println("Twin snowflakes found.")
	} else {
		fmt.Println("No two snowflakes are alike.")
	}
}
