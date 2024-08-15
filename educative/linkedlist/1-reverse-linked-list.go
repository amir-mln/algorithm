package linkedlist

import "github.com/amir-mln/algorithm/educative/common"

var ReverseLinkedList = map[string]func(head **common.LinkedListNode){
	"linear-solution": func(head **common.LinkedListNode) {
		var prev, current *common.LinkedListNode = nil, *head

		for current != nil {
			n := current.Next
			current.Next = prev
			prev = current
			current = n
		}

		*head = prev
	},
}
