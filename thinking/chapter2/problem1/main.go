package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
# Halloween Haul

You just moved into a strange neighborhood. You notice that the roads in your neighborhood
form a binary tree, with the houses forming the leaves of the tree. But this doesn't really
matter to you, because it's all about the candy during Halloween! However having just moved
in, you might end up getting lost. So you create a plan: start at the root of your
neighborhood and walk to every house and ask for candy! Before setting off on your adventure
however, you want to calculate the minimum number of roads you'll need to walk and the total
amount of candy you'll get.

Input:

The input will contain 5 test cases. Each test case is a line containing a single string, less
than 256 characters long, describing the tree your neighborhood forms. A binary tree can be
recursively described as either:

-	A leaf `c`, 1 ≤ `c` ≤ 20, the amount of candy received from the house
-	(`c` `c`) - where `c` represents a tree leaf node

Output:

The output will contain 5 lines of output, each a pair of integers where the first integer is
the minimum number of roads needed to be traversed to get all the candy (starting from the root
of the tree, and not needing to return) and the second integer	represents the total amount of
candy you'll collect.

The time limit for solving the test cases is 2 second.

Link:

https://dmoj.ca/problem/dwite12c1p4
*/
type Problem struct {
	Root *Node
}

func (p *Problem) MinPath() int {
	var sumPath func(node *Node) int
	sumPath = func(node *Node) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}
		sum := 0
		if node.Left != nil {
			sum += 2 + sumPath(node.Left)
		}
		if node.Right != nil {
			sum += 2 + sumPath(node.Right)
		}

		return sum
	}

	var maxDepth func(node *Node) int
	maxDepth = func(node *Node) int {
		if node.Left == nil && node.Right == nil {
			return 0
		}

		l, r := 0, 0
		if node.Left != nil {
			l = maxDepth(node.Left)
		}
		if node.Right != nil {
			r = maxDepth(node.Right)
		}

		return 1 + max(l, r)
	}

	return sumPath(p.Root) - maxDepth(p.Root)
}

func (p *Problem) Candies() int {
	sum := 0
	idx := 0
	nodes := []*Node{p.Root}
	for idx < len(nodes) {
		node := nodes[idx]
		if node.Left == nil && node.Right == nil {
			sum += *node.Value
		} else {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}

			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		idx++
	}

	return sum
}

type Node struct {
	Value *int
	Left  *Node
	Right *Node
}

func ParseTree(str string) *Node {
	digits := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	var helper func(*int) *Node
	helper = func(i *int) *Node {
		switch rune(str[*i]) {
		case '(':
			*i++
			node := helper(i)
			*i++ // skipping the space
			if *i < len(str)-1 {
				temp := node
				node = &Node{}
				node.Left = temp
				node.Right = helper(i)
			}

			return node
		case '0', '1', '2', '3', '4',
			'5', '6', '7', '8', '9':
			n := &Node{}
			left := &Node{Value: new(int)}
			for d, ok := digits[rune(str[*i])]; ok; {
				if *left.Value != 0 {
					*left.Value *= 10
				}

				*left.Value += d
				*i++
				d, ok = digits[rune(str[*i])]
			}

			if rune(str[*i]) == ')' {
				return left
			}

			n.Left = left
			*i++ // skipping the space
			if rune(str[*i]) == '(' {
				*i++
				n.Right = helper(i)
			} else {
				right := &Node{Value: new(int)}
				for d, ok := digits[rune(str[*i])]; ok; {
					if *right.Value != 0 {
						*right.Value *= 10
					}

					*right.Value += d
					*i++
					d, ok = digits[rune(str[*i])]
				}
				n.Right = right
				*i++ // skipping the ')' character
			}

			return n
		default:
			return nil
		}
	}

	i := 0
	return helper(&i)
}

func main() {
	stdIn := bufio.NewScanner(os.Stdin)
	stdOut := bufio.NewWriter(os.Stdout)

	for i := 0; i < 5; i++ {
		stdIn.Scan()
		root := ParseTree(stdIn.Text())
		p := Problem{root}
		stdOut.WriteString(fmt.Sprintln(p.MinPath(), p.Candies()))
	}
	stdOut.Flush()
}
