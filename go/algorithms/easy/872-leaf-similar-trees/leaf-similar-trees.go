package main

import (
	"fmt"
	"sort"
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

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
// }

func getAllLeaves(tree *TreeNode) []int {
	var leaves []int

	if tree.Left == nil && tree.Right == nil {
		leaves = append(leaves, tree.Val)
	}

	if tree.Left != nil {
		leaves = append(leaves, getAllLeaves(tree.Left)...)
	}
	if tree.Right != nil {
		leaves = append(leaves, getAllLeaves(tree.Right)...)
	}

	return leaves
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Leaves := getAllLeaves(root1)
	root2Leaves := getAllLeaves(root2)

	if len(root1Leaves) != len(root2Leaves) {
		return false
	}

	sort.Ints(root1Leaves)
	sort.Ints(root2Leaves)

	for i := range root1Leaves {
		if root1Leaves[i] != root2Leaves[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(`test`)

	arr := []*int{
		newInt(1),
		newInt(3),
		newInt(2),
	}

	arr2 := []*int{
		newInt(1),
		newInt(2),
		// newInt(2),
		newInt(3),
		// nil,
		// nil,
		// newInt(3),
		// newInt(4),
		// nil,
		// nil,
		// newInt(4),
	}

	root := ArrayToBinaryTree(arr)
	root2 := ArrayToBinaryTree(arr2)

	result := leafSimilar(root, root2)

	// var result int = maxDepth(root)
	fmt.Println(result)
}
