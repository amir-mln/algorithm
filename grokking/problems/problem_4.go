package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

/*
# Remove Nth Node from End of List

Given a singly linked list, remove the Nth node from the end
of the list and return its head.

Constraints:

 1. The number of nodes in the list is k.
 2. 1 ≤ k ≤ 10**3
 3. -10**3 ≤ Node.Value ≤ 10**3
 4. 1 ≤ n ≤ k
*/
var Problem4 = map[string]func(l *grokking.LinkedList, n int){
	"two-pointer": func(l *grokking.LinkedList, n int) {
		p1, p2 := l.Head, l.Head

		for i := 0; i < n; i++ { // p2 != nil ???
			p2 = p2.Next
		}

		if p2 == nil {
			temp := l.Head
			l.Head = l.Head.Next
			temp.Next = nil
		} else {
			for {
				if p2.Next == nil {
					temp := p1.Next
					p1.Next = p1.Next.Next
					temp.Next = nil
					break
				}
				p1 = p1.Next
				p2 = p2.Next

			}
		}
	},
	"iterative": func(l *grokking.LinkedList, n int) {
		k := 0
		c := l.Head

		for c != nil {
			k++
			c = c.Next
		}

		if k == n {
			temp := l.Head
			l.Head = l.Head.Next
			temp.Next = nil
		} else {

			k -= n
			c = l.Head
			for {
				if k == 1 {
					temp := c.Next
					c.Next = c.Next.Next
					temp.Next = nil
					break
				}

				c = c.Next
				k--
			}

		}
	},
}
