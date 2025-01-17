package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func fromArrToLinkedList(arr []int, pos int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head
	var posNode *ListNode = nil

	for index, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next

		if index+1 == pos {
			posNode = current
		}
	}

	current.Next = posNode

	return head
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA := headA

	for currentA != nil {
		currentB := headB
		for currentB != nil {
			if currentA == currentB {
				return currentA
			}

			currentB = currentB.Next
		}

		currentA = currentA.Next
	}

	return nil
}

func main() {
	fmt.Println(`test`)
	intersect := fromArrToLinkedList([]int{3, 2, 0, -4}, 1)
	list1 := &ListNode{Val: 3, Next: intersect}
	list2 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersect}}
	fmt.Println(getIntersectionNode(list1, list2))
}
