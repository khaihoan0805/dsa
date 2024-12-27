package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func InOrderTraversal(root *TreeNode) {
// 	if root != nil {
// 		InOrderTraversal(root.Left)
// 		fmt.Print(root.Val, " ")
// 		InOrderTraversal(root.Right)
// 	}
// }

// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) == 0 {
// 		return nil
// 	}

// 	pivotIndex := (len(nums) - 1) / 2

// 	root := &TreeNode{Val: nums[pivotIndex]}

// 	if pivotIndex > 0 {
// 		pivotInsert(nums[0:pivotIndex], root)
// 	}

// 	if pivotIndex < len(nums)-1 {
// 		pivotInsert(nums[pivotIndex+1:], root)
// 	}

// 	return root
// }

// func pivotInsert(nums []int, node *TreeNode) bool {
// 	pivotIndex := (len(nums) - 1) / 2

// 	fmt.Println(nums)
// 	fmt.Println(pivotIndex)

// 	// node = &TreeNode{Val: nums[pivotIndex]}

// 	Insert(nums[pivotIndex], node)
// 	if pivotIndex > 0 {
// 		pivotInsert(nums[0:pivotIndex], node)
// 	}

// 	if pivotIndex < len(nums)-1 {
// 		pivotInsert(nums[pivotIndex+1:], node)
// 	}

// 	return true
// }

// func Insert(val int, root *TreeNode) bool {
// 	var before, after *TreeNode = nil, root

// 	for after != nil {
// 		before = after

// 		if val < before.Val {
// 			after = after.Left
// 		} else if val > before.Val {
// 			after = after.Right
// 		} else {
// 			return false
// 		}
// 	}

// 	newNode := &TreeNode{Val: val}
// 	if before == nil {
// 		root = newNode
// 		return true
// 	}

// 	if newNode.Val < before.Val {
// 		before.Left = newNode
// 	} else {
// 		before.Right = newNode
// 	}

// 	return true
// }

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)

	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	l, r := 0, n-1
	mid := l + (r-l)/2

	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func main() {
	fmt.Println(`test`)

	arr := []int{0, 1}

	root := sortedArrayToBST(arr)

	// var result int = maxDepth(root)
	fmt.Println(root)
}
