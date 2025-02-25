package main

import (
	"fmt"
	"math"
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

// func pairSum(head *ListNode) int {
// 	max := 0

// 	var arr []int

// 	current := head
// 	for current != nil {
// 		arr = append(arr, current.Val)
// 		current = current.Next
// 	}

// 	for i := 0; i < len(arr)/2; i++ {
// 		sum := arr[i] + arr[len(arr)-1-i]
// 		if sum > max {
// 			max = sum
// 		}
// 	}

// 	return max
// }

func pairSum(head *ListNode) int {
	ans := 0
	var newList *ListNode
	current := head
	currHalf := head

	// Traverse the first half and reverse it
	for currHalf != nil && currHalf.Next != nil {
		currHalf = currHalf.Next.Next
		temp := current.Next
		current.Next = newList
		newList = current
		current = temp
	}

	// Compare the first half and reversed second half
	for current != nil {
		ans = int(math.Max(float64(ans), float64(current.Val+newList.Val)))
		current = current.Next
		newList = newList.Next
	}

	return ans
}
func main() {
	fmt.Println(`test`)

	arr := []int{5, 4, 2, 1}

	// Create the linked list
	head := createLinkedList(arr)

	var result int = pairSum(head)
	fmt.Println(result)

}
