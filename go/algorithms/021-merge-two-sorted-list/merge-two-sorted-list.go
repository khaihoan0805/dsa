package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	pointer := result

	if list1 == nil && list2 != nil {
		return list2
	}

	if list2 == nil && list1 != nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pointer.Next = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			list2 = list2.Next
		}

		pointer = pointer.Next
	}

	if list1 != nil {
		pointer.Next = list1
	} else {
		pointer.Next = list2
	}

	return result.Next
}

func main() {
	fmt.Println(`test`)
	list1 := new(ListNode)
	list1.Val = 1
	list1.Next = new(ListNode)
	list1.Next.Val = 2
	list2 := new(ListNode)
	list2.Val = 3

	fmt.Println(mergeTwoLists(list1, list2))
}
