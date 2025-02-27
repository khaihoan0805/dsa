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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		// do delete operation

		// case 1: this is a leaf node, directly delete
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// case 2: it has only one child, let the one child to replace it
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}

		// case 3: it has both left and right child
		if root.Left != nil && root.Right != nil {
			// Found the smallest node on the right to replace it
			minSubTreeNode := getMin(root.Right)
			leftSubTree := root.Left
			rightSubTree := deleteNode(root.Right, minSubTreeNode.Val)
			minSubTreeNode.Left = leftSubTree
			minSubTreeNode.Right = rightSubTree
			return minSubTreeNode
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	}

	return getMin(root.Left)
}

func main() {
	fmt.Println(`test`)

	// [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
	// arr := []*int{
	// 	newInt(1),
	// 	nil,
	// 	newInt(2),
	// 	newInt(3),
	// 	newInt(4),
	// 	nil,
	// 	nil,
	// 	newInt(5),
	// 	newInt(6),
	// 	nil,
	// 	newInt(7),
	// 	nil,
	// 	nil,
	// 	nil,
	// 	newInt(8),
	// }

	arr := []*int{
		newInt(5),
		newInt(3),
		newInt(6),
		newInt(2),
		newInt(4),
		nil,
		newInt(7),
		// newInt(3),
		// newInt(4),
		// nil,
		// nil,
		// nil,
		// newInt(5),
	}

	root := ArrayToBinaryTree(arr)
	result := deleteNode(root, 3)

	fmt.Println(result)
}
