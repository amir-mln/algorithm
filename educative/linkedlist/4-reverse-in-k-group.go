package linkedlist

import (
	"github.com/amir-mln/algorithm/datastructure"
)

var ReverseInKGroup = map[string]func(*datastructure.LinkedList, int){
	"linear-solution": func(ll *datastructure.LinkedList, k int) {
		var tracker *datastructure.LinkedListNode = ll.Head
		var headIsChanged bool

		moveP2 := func() bool {
			for i := 1; i < k && tracker != nil; i++ {
				tracker = tracker.Next
			}

			if tracker == nil {
				return false
			}

			tracker = tracker.Next
			return true
		}

		var prevTail, current *datastructure.LinkedListNode = nil, ll.Head
		for tracker != nil {
			if !moveP2() {
				if prevTail != nil {
					prevTail.Next = current
				}
				return
			}

			var prev, tail *datastructure.LinkedListNode = nil, current

			for i := 1; i <= k; i++ {
				n := current.Next
				current.Next = prev
				prev = current
				current = n
			}
			if !headIsChanged {
				headIsChanged = true
				ll.Head = prev
			}

			current = tracker

			if prevTail != nil {
				prevTail.Next = prev
			}

			prevTail = tail

		}
	},
}
