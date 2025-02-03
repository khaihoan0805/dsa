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
func insertNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value}
	}
	if value < root.Val {
		root.Left = insertNode(root.Left, value)
	} else {
		root.Right = insertNode(root.Right, value)
	}
	return root
}

// Function to create a tree from an array
func createTreeFromArray(arr []int) *TreeNode {
	var root *TreeNode
	for _, value := range arr {
		root = insertNode(root, value)
	}
	return root
}

// Function to print the tree (in-order traversal)
// func printTree(root *TreeNode) {
// 	if root != nil {
// 		printTree(root.Left)
// 		fmt.Println(root.Val)
// 		printTree(root.Right)
// 	}
// }

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	root.Right = invertTree(root.Right)
	root.Left = invertTree(root.Left)

	return root
}

func main() {
	fmt.Println(`test`)
	root := createTreeFromArray([]int{4, 2, 7, 1, 3, 6, 9})
	var result *TreeNode = invertTree(root)
	fmt.Println(result)
}
