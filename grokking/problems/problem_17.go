package problems

import "github.com/amir-mln/algorithm/grokking"

var ReverseLinkedList = map[string]func(head **grokking.LinkedListNode){
	"linear-solution": func(head **grokking.LinkedListNode) {
		var prev, current *grokking.LinkedListNode = nil, *head

		for current != nil {
			n := current.Next
			current.Next = prev
			prev = current
			current = n
		}

		*head = prev
	},
}
