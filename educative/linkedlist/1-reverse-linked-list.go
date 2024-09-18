package linkedlist

import "github.com/amir-mln/algorithm/datastructure"

var ReverseLinkedList = map[string]func(head **datastructure.LinkedListNode){
	"linear-solution": func(head **datastructure.LinkedListNode) {
		var prev, current *datastructure.LinkedListNode = nil, *head

		for current != nil {
			n := current.Next
			current.Next = prev
			prev = current
			current = n
		}

		*head = prev
	},
}
