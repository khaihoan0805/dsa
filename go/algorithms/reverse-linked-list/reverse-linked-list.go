package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) *ListNode {
// 	temp := []int{}

// 	current := head

// 	for current != nil {
// 		temp = append(temp, current.Val)
// 		current = current.Next
// 	}

// 	result := &ListNode{}
// 	currentR := result
// 	for index := len(temp) - 1; index >= 0; index-- {
// 		currentR.Next = &ListNode{Val: temp[index]}
// 		currentR = currentR.Next
// 	}

// 	return result.Next
// }

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Previous node starts as nil
	curr := head       // Current node starts at the head

	// Traverse the list
	for curr != nil {
		next := curr.Next // Save the next node

		curr.Next = prev // Reverse the link

		// Move pointers forward
		prev = curr // Move prev to the current node
		curr = next // Move curr to the next node
	}

	// prev is now the new head of the reversed list
	return prev
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

func main() {
	fmt.Println(`test`)
	list := fromArrToLinkedList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("list: ", list)
	isCycle := reverseList(list)

	fmt.Println("isCycle: ", isCycle)
}
