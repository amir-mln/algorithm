package problems

import (
	"github.com/amir-mln/algorithm/grokking"
)

var ReverseInEvenGroups = map[string]func(list *grokking.LinkedList){
	"linear-solution": func(list *grokking.LinkedList) {
		var prevTail, current *grokking.LinkedListNode = nil, list.Head

		for i := 1; current != nil; i++ {
			var prev, tail *grokking.LinkedListNode = nil, current

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
