package datastructure

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
