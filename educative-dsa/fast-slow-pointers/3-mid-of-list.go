package fastslowpointers

import "educative-dsa/common"

var MiddleOfLinkedList = map[string]func(head *common.LinkedListNode) *common.LinkedListNode{
	"fast-slow-pointer": func(head *common.LinkedListNode) *common.LinkedListNode {
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
