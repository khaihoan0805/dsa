package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func fromArrToLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	var posNode *ListNode = nil

	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	current.Next = posNode

	return head
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

func main() {
	fmt.Println(`test`)
	list := fromArrToLinkedList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println("list: ", list)
	isCycle := removeElements(list, 3)

	fmt.Println("isCycle: ", isCycle)
}
