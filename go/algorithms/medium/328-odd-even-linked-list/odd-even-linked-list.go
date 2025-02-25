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

// func oddEvenList(head *ListNode) *ListNode {
// 	odd := &ListNode{}
// 	even := &ListNode{}
// 	current := head

// 	lastNodeOfodd := odd
// 	lastNodeOfEven := even
// 	for i := 1; current != nil; i++ {
// 		newNode := &ListNode{Val: current.Val}
// 		if i%2 == 1 {
// 			lastNodeOfodd.Next = newNode
// 			lastNodeOfodd = lastNodeOfodd.Next
// 		} else {
// 			lastNodeOfEven.Next = newNode
// 			lastNodeOfEven = lastNodeOfEven.Next
// 		}
// 		current = current.Next
// 	}

// 	lastNodeOfodd.Next = even.Next

// 	return odd.Next
// }

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	} else if head.Next == nil {
		return head
	}

	evenHead := head
	oddHead := head.Next

	evenTail := evenHead
	oddTail := oddHead

	isEven := true
	curr := oddHead.Next
	for curr != nil {
		if isEven {
			evenTail.Next = curr
			evenTail = curr
		} else {
			oddTail.Next = curr
			oddTail = curr
		}
		curr = curr.Next
		isEven = !isEven
	}

	// if it is currently even tail's (or odd) turn, but the list ran out,
	// then it didn't have a chance to set its next elmt, we need to set it to nil
	// (note that we are modifying the pointers in the original list,
	//  instead of creating a new list. the original `next` elmt came from original list)
	if isEven {
		evenTail.Next = nil
	} else {
		oddTail.Next = nil
	}

	evenTail.Next = oddHead

	return head
}

func main() {
	fmt.Println(`test`)

	arr := []int{1, 2, 3, 4, 5}

	// Create the linked list
	head := createLinkedList(arr)

	var result *ListNode = oddEvenList(head)
	fmt.Println(result)

}
