package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(root *TreeNode, values *[]int) []int {
	if root == nil {
		return *values
	}

	walk(root.Left, values)
	*values = append(*values, root.Val)
	fmt.Println(values)
	walk(root.Right, values)

	return *values
}

func inorderTraversal(root *TreeNode) []int {
	result := new([]int)

	return walk(root, result)
}

func main() {
	fmt.Println(`test`)

	var test = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	var result []int = inorderTraversal(test)
	fmt.Println(result)
}
