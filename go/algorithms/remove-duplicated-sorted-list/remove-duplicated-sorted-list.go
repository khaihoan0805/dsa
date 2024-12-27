package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var result *ListNode = head
	var before *ListNode = result
	var after *ListNode = result.Next

	for after != nil {
		if before.Val == after.Val {
			var temp *ListNode = after
			before.Next = temp.Next
			after = temp.Next
		} else {
			before = after
			after = after.Next
		}
	}

	return result
}

func main() {
	fmt.Println(`test`)

	list1 := new(ListNode)
	list1.Val = 1
	list1.Next = new(ListNode)
	list1.Next.Val = 2

	fmt.Println(deleteDuplicates(list1))
}
