package educative

import "github.com/amir-mln/algorithm/datastructure"

// TODO: Fast & Slow Pointers

var MiddleOfLinkedList = map[string]func(head *datastructure.LinkedListNode) *datastructure.LinkedListNode{
	"fast-slow-pointer": func(head *datastructure.LinkedListNode) *datastructure.LinkedListNode {
		if head.Next == nil {
			return head
		}

		sp, fp := head, head

		for {
			if fp == nil || fp.Next == nil {
				return sp
			}

			sp = sp.Next
			fp = fp.Next
			fp = fp.Next
		}
	},
}
