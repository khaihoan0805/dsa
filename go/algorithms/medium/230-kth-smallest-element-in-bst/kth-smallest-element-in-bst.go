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

func newInt(val int) *int {
	return &val
}

func kthSmallest(root *TreeNode, k int) int {
	result := -1
	bfs(root, &k, &result)
	return result
}

func bfs(root *TreeNode, counter, result *int) {
	if root == nil {
		return
	}

	bfs(root.Left, counter, result)

	*counter--
	if *counter == 0 {
		*result = root.Val
		return
	}
	bfs(root.Right, counter, result)
}

func main() {
	fmt.Println(`test`)

	arr := []*int{
		// newInt(5),
		// newInt(4),
		// newInt(8),
		// newInt(11),
		// nil,
		// newInt(13),
		// newInt(4),
		// newInt(7),
		// newInt(2),
		// nil,
		// nil,
		// nil,
		newInt(5),
		newInt(3),
		newInt(6),
		newInt(2),
		newInt(4),
		nil,
		nil,
		newInt(1),
	}

	root := ArrayToBinaryTree(arr)
	var result int = kthSmallest(root, 6)
	fmt.Println(result)
}
