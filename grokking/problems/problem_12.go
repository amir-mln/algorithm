package problems

import "github.com/amir-mln/algorithm/grokking"

// TODO: Fast & Slow Pointers

var MiddleOfLinkedList = map[string]func(head *grokking.LinkedListNode) *grokking.LinkedListNode{
	"fast-slow-pointer": func(head *grokking.LinkedListNode) *grokking.LinkedListNode {
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
