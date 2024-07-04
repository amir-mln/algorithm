package twopointers

import "fmt"

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) String() string {
	s := make([]int, 0)
	c := l.Head

	for c != nil {
		s = append(s, c.Value)
		c = c.Next
	}

	return fmt.Sprint(s)
}

func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
}

func (l *LinkedList) InsertFromList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := &LinkedListNode{
			Value: lst[i],
			Next:  nil,
		}

		l.InsertNodeAtHead(newNode)
	}
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

var RemoveNthNode = map[string]func(l *LinkedList, n int){
	"two-pointer-solution": func(l *LinkedList, n int) {
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
	"iterative-solution": func(l *LinkedList, n int) {
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
