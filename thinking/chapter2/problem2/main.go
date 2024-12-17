package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
# Descendant Distance

In this problem, we are given a family tree and a specified distance `d`. The score
for each node in the tree is the number of descendants it has at distance `d`. Our
task is to output the nodes with high scores; I’ll explain exactly how many nodes
that is in the Output section. Generalizing, we can say that, for any node, the
number of descendants at distance `d` is the number of nodes that are exactly `d`
edges down the tree from that node.

Input:

The first line of input gives the number of test cases that will follow. Each test
case consists of the following lines:
  - A line containing two integers, `n` and `d`, where n tells us how many
    more lines there are for this test case and `d` specifies the descendant
    distance of interest.
  - n lines used to build the tree. Each of these lines consists of the
    name of a node, an integer m, and m node names giving the children
    of this node. Each name is at most 10 characters long. These lines
    can come in any order—there’s no requirement that parent lines come
    before their descendant lines.

There are at most 1,000 nodes in any test case. Here is a possible input, asking for
the nodes with the most descendants at a distance of two:

	1
	7 2
	Lucas 1 Enzo
	Zara 1 Amber
	Sana 2 Gabriel Lucas
	Enzo 2 Min Becky
	Kevin 2 Jad Cassie
	Amber 4 Vlad Sana Ashley Kevin
	Vlad 1 Omar

Output:

The output for each test case has two parts. First, output the following line:

	Tree i:

Where i is 1 for the first test case, 2 for the second test case, and so on. Then,
output information for the nodes with high scores (where the score for a node is
the number of descendants it has at distance `d`), sorted from most to least. Output
the names that are tied for the number of descendants at distance `d` in alphabetical
order. Use the following rules to determine how many names to output:

  - If there are three or fewer names with descendants at distance `d`,
    output them all.
  - If there are more than three names with descendants at distance `d`,
    start by outputting those with the top three scores, starting from
    the highest score. Then, output each other name whose score is
    the same as the third score from the top. For example, if we have
    names with eight, eight, two, two, two, one, and one descendants at
    distance `d`, we would output information for five names: those with
    eight, eight, two, two, and two descendants at distance `d`.

For each name that we’re required to output, we output a line consisting of the name,
followed by a space, followed by its number of descendants at distance `d`. Output
for each test case is separated from the next by a blank line. Here is the output for
the above sample input:

	Tree 1:
	Amber 5
	Zara 4
	Lucas 2

The time limit for solving the test cases is 0.6 seconds.
Link:

https://dmoj.ca/problem/ecna05b
*/
type Node struct {
	Value    string
	Children []*Node
}

func (n *Node) Descendants(d int) int {
	if d == 0 {
		return 1
	}

	if len(n.Children) == 0 {
		return 0
	}

	sum := 0
	for _, child := range n.Children {
		sum += child.Descendants(d - 1)
	}

	return sum
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		nodes := make(map[string]*Node)
		sc.Scan()
		txt := strings.Fields(sc.Text())
		count, _ := strconv.Atoi(txt[0])
		d, _ := strconv.Atoi(txt[1])

		for j := 0; j < count; j++ {
			sc.Scan()
			fields := strings.Fields(sc.Text())
			parent := fields[0]
			children := fields[2:]

			if _, ok := nodes[parent]; !ok {
				nodes[parent] = &Node{Value: parent}
			}
			for _, child := range children {
				if _, ok := nodes[child]; !ok {
					nodes[child] = &Node{Value: child}
				}
				nodes[parent].Children = append(nodes[parent].Children, nodes[child])
			}
		}

		out.WriteString(fmt.Sprintf("Tree %d:\n", i+1))

		type result struct {
			n string
			c int
		}
		results := make([]result, 0)
		for _, n := range nodes {
			if count := n.Descendants(d); count > 0 {
				results = append(results, result{n: n.Value, c: count})
			}
		}
		slices.SortFunc(results, func(a, b result) int {
			if a.c > b.c {
				return -1
			} else if a.c < b.c {
				return +1
			} else {
				return strings.Compare(a.n, b.n)
			}
		})
		threshold := results[min(2, len(results)-1)]
		for _, res := range results {
			if res.c >= threshold.c {
				out.WriteString(fmt.Sprintln(res.n, res.c))
			}
		}
		out.WriteString("\n")
	}

	out.Flush()
}
