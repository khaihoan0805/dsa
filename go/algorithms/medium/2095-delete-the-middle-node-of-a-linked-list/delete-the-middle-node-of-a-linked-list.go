package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil // Return nil for an empty array
	}

	// Create the head node
	head := &ListNode{Val: arr[0]}
	current := head

	// Iterate through the array and create the linked list
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}

	return head
}

// with this approach, we can delete with any division by changing the division to a paramater of the function
// func deleteMiddle(head *ListNode) *ListNode {
// 	if head.Next == nil {
// 		return nil
// 	}
// 	var current *ListNode = head

// 	currentCount := 1

// 	tail := current
// 	maxCount := 1

// 	for tail.Next != nil {
// 		tail = tail.Next
// 		maxCount++

// 		middle := (maxCount) / 2

// 		for currentCount < middle {
// 			current = current.Next
// 			currentCount++
// 		}
// 	}

// 	current.Next = current.Next.Next

// 	return head
// }

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {
	fmt.Println(`test`)

	arr := []int{1, 2, 3, 4}

	// Create the linked list
	head := createLinkedList(arr)

	var result *ListNode = deleteMiddle(head)
	fmt.Println(result)

}
