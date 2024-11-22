package problems

import "github.com/amir-mln/algorithm/grokking"

// TODO: Fast & Slow Pointers

var LinkedListCycle = map[string]func(list *grokking.LinkedList) bool{
	"fast-slow-pointers": func(list *grokking.LinkedList) bool {
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
