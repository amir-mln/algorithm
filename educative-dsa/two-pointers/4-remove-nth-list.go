package twopointers

import common "educative-dsa/common"

var RemoveNthNode = map[string]func(l *common.LinkedList, n int){
	"two-pointer-solution": func(l *common.LinkedList, n int) {
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
	"iterative-solution": func(l *common.LinkedList, n int) {
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
