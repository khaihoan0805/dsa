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
func InOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		InOrderTraversal(root.Right)
	}
}

func newInt(val int) *int {
	return &val
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
	}

	return false
}

func isSymmetric(root *TreeNode) bool {
	var left, right *TreeNode = root.Left, root.Right

	if left == nil && right == nil {
		return true
	}

	if left != nil && right == nil {
		return false
	}

	if left == nil && right != nil {
		return false
	}

	if left.Val == right.Val {
		return isSameTree(left.Left, right.Right) && isSameTree(left.Right, right.Left)
	}

	return false
}

func main() {
	fmt.Println(`test`)

	arr := []*int{
		newInt(2),
		newInt(3),
		newInt(3),
		newInt(4),
		newInt(5),
		newInt(5),
		newInt(4),
		nil,
		nil,
		newInt(8),
		newInt(9),
		newInt(9),
		newInt(8),
	}
	root := ArrayToBinaryTree(arr)

	var result bool = isSymmetric(root)
	fmt.Println(result)
}
