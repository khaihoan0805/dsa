package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a new value into the binary tree
func (t *TreeNode) Insert(val int) {
	if val < t.Val {
		if t.Left == nil {
			t.Left = &TreeNode{Val: val}
		} else {
			t.Left.Insert(val)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Val: val}
		} else {
			t.Right.Insert(val)
		}
	}
}

// ArrayToBinaryTree converts an array of integers to a binary tree
func ArrayToBinaryTree(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		current := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: *arr[i]}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

// func preorderTraversal(head *TreeNode) []int {
// 	var result []int = make([]int, 0)

// 	if head != nil {
// 		result = append([]int{head.Val}, preorderTraversal(head.Left)...)
// 		result = append(result, preorderTraversal(head.Right)...)
// 	}

// 	return result
// }

func postorderTraversal(root *TreeNode) []int {
	var result []int = make([]int, 0)

	if root != nil {
		result = append(result, postorderTraversal(root.Left)...)
		result = append(result, postorderTraversal(root.Right)...)

		result = append(result, root.Val)
	}

	return result
}

func newInt(val int) *int {
	return &val
}
func main() {
	fmt.Println(`test`)

	// arr := []*int{
	// 	newInt(5),
	// 	newInt(4),
	// 	newInt(8),
	// 	newInt(11),
	// 	nil,
	// 	newInt(13),
	// 	newInt(4),
	// 	newInt(7),
	// 	newInt(2),
	// 	nil,
	// 	nil,
	// 	nil,
	// 	newInt(1),
	// }

	arr := []*int{
		newInt(-2),
		nil,
		newInt(-3),
	}

	root := ArrayToBinaryTree(arr)
	var result []int = postorderTraversal(root)
	fmt.Println(result)

	// var result bool = hasPathSum(root, -5)
	// fmt.Println(result)
}
