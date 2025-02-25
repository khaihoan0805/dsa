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

// InOrderTraversal prints the tree in order
// func InOrderTraversal(root *TreeNode) {
// 	if root != nil {
// 		InOrderTraversal(root.Left)
// 		fmt.Print(root.Val, " ")
// 		InOrderTraversal(root.Right)
// 	}
// }

func newInt(val int) *int {
	return &val
}

func isGoodNode(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= max {
		max = node.Val
		count++
	}

	return isGoodNode(node.Left, max) + isGoodNode(node.Right, max) + count
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min := root.Val

	return isGoodNode(root.Left, min) + isGoodNode(root.Right, min) + 1
}

func main() {
	fmt.Println(`test`)

	// [3,1,4,3,null,1,5]
	arr := []*int{
		newInt(3),
		newInt(1),
		newInt(4),
		newInt(3),
		nil,
		newInt(1),
		newInt(5),
	}

	root := ArrayToBinaryTree(arr)
	result := goodNodes(root)

	fmt.Println(result)
}
