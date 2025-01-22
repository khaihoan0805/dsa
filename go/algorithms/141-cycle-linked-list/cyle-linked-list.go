package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// my own approach

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	current := head
// 	// next := current.Next

// 	looped := []*ListNode{}

// 	for current.Next != nil {
// 		for _, value := range looped {
// 			if current.Next == value {
// 				return true
// 			}
// 		}

// 		looped = append(looped, current)

// 		current = current.Next
// 	}

// 	return false
// }

// 2 pointers (rare and tortoise algorithm)

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}

// 	slow_pointer, fast_pointer := head, head

// 	for fast_pointer.Next != nil && fast_pointer.Next.Next != nil {
// 		slow_pointer = slow_pointer.Next
// 		fast_pointer = fast_pointer.Next.Next
// 		if slow_pointer == fast_pointer {
// 			return true
// 		}
// 	}

// 	return false
// }

// using hash table
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	current := head
	looped := make(map[*ListNode]bool)

	for current != nil {
		if looped[current] {
			return true
		}

		looped[current] = true

		current = current.Next
	}

	return false
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

func main() {
	fmt.Println(`test`)
	list := fromArrToLinkedList([]int{3, 2, 0, -4}, 1)
	fmt.Println("list: ", list)
	isCycle := hasCycle(list)
	fmt.Println("isCycle: ", isCycle)
}
