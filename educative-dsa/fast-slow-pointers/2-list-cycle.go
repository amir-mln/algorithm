package fastslowpointers

import "educative-dsa/common"

var LinkedListCycle = map[string]func(list *common.LinkedList) bool{
	"fast-slow-pointers": func(list *common.LinkedList) bool {
		fp, sp := list.Head.Next, list.Head

		for fp != nil && fp != sp {
			if fp.Next != nil {
				fp = fp.Next
			}

			if fp != nil {
				fp = fp.Next
			}

			sp = sp.Next
		}

		return fp == sp
	},
}
