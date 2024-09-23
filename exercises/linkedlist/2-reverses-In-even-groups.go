package linkedlist

import (
	"github.com/amir-mln/algorithm/datastructure"
)

var ReverseInEvenGroups = map[string]func(list *datastructure.LinkedList){
	"linear-solution": func(list *datastructure.LinkedList) {
		var prevTail, current *datastructure.LinkedListNode = nil, list.Head

		for i := 1; current != nil; i++ {
			var prev, tail *datastructure.LinkedListNode = nil, current

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
