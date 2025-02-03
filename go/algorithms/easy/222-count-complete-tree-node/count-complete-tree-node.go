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

// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return 1 + countNodes(root.Left) + countNodes(root.Right)
// }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	fmt.Println(`test`)
	root := createTreeFromArray([]int{1, 2, 3, 4, 5, 6})
	var result int = countNodes(root)
	fmt.Println(result)
}
