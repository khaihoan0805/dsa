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

func bfs(root *TreeNode, level int, traversed *[]int) []int {
	if root == nil {
		return *traversed
	}
	if level == len(*traversed)+1 {
		*traversed = append(*traversed, root.Val)
	}

	if root.Right != nil {
		bfs(root.Right, level+1, traversed)
	}
	if root.Left != nil {
		bfs(root.Left, level+1, traversed)
	}

	return *traversed
}

func rightSideView(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	result = append(result, root.Val)
	if root.Right != nil {
		bfs(root.Right, 2, &result)
	}

	if root.Left != nil {
		bfs(root.Left, 2, &result)
	}

	return result
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
		newInt(1),
		newInt(2),
		// newInt(3),
		// newInt(4),
		// nil,
		// nil,
		// nil,
		// newInt(5),
	}

	root := ArrayToBinaryTree(arr)
	result := rightSideView(root)

	fmt.Println(result)
}
