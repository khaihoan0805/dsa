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

func subPathSum(root *TreeNode, prevs []int, targetSum int) int {
	if root == nil {
		return 0
	}
	count := 0

	path := append(prevs, root.Val)

	sum := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetSum {
			count++
		}
	}

	return count + subPathSum(root.Left, path, targetSum) + subPathSum(root.Right, path, targetSum)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Val == targetSum {
		count++
	}

	return count + subPathSum(root.Left, []int{root.Val}, targetSum) + subPathSum(root.Right, []int{root.Val}, targetSum)
}

func main() {
	fmt.Println(`test`)

	// 10,5,-3,3,2,null,11,3,-2,null,1
	arr := []*int{
		newInt(10),
		newInt(5),
		newInt(-3),
		newInt(3),
		newInt(2),
		nil,
		newInt(11),
		newInt(3),
		newInt(-2),
		nil,
		newInt(1),
	}

	root := ArrayToBinaryTree(arr)
	result := pathSum(root, 8)

	fmt.Println(result)
}
