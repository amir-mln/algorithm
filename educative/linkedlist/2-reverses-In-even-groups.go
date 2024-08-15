package linkedlist

import (
	"github.com/amir-mln/algorithm/educative/common"
)

var ReverseInEvenGroups = map[string]func(list *common.LinkedList){
	"linear-solution": func(list *common.LinkedList) {
		var prevTail, current *common.LinkedListNode = nil, list.Head

		for i := 1; current != nil; i++ {
			var prev, tail *common.LinkedListNode = nil, current

			for j := 1; j <= i && current != nil; j++ {
				n := current.Next
				current.Next = prev
				prev = current
				current = n
			}

			if prevTail != nil {
				prevTail.Next = prev
			}

			prevTail = tail
		}
	},
}
