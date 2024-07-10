package linkedlistmanipulation

import (
	"educative-dsa/common"
)

var ReorderList = map[string]func(list *common.LinkedList){
	"linear-solution": func(list *common.LinkedList) {
		p1, p2 := list.Head, list.Head

		for p1.Next != nil {
			p1 = p1.Next

			if p1.Next != nil {
				p1 = p1.Next
				p2 = p2.Next
			}
		}

		ReverseLinkedList["linear-solution"](&p2.Next)

		p1 = list.Head
		for p1 != p2 {
			p3, p4 := p1.Next, p2.Next
			p1.Next = p4
			p2.Next = p4.Next
			p4.Next = p3
			p1 = p3
		}
	},
}

// 0 11 -22 -12 21
// 0 -22 -12 21
// 0 11 -22 -12 21
