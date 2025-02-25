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

func getAncestors(root *TreeNode, p *TreeNode, prev []*TreeNode) ([]*TreeNode, bool) {
	if root == nil {
		return prev, false
	}

	path := append(prev, root)

	if root == p {
		return path, true
	}

	left, found := getAncestors(root.Left, p, path)
	if found {
		return left, found
	} else {
		return getAncestors(root.Right, p, path)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pAns, _ := getAncestors(root, p, []*TreeNode{})
	qAns, _ := getAncestors(root, q, []*TreeNode{})

	set := make(map[*TreeNode]bool)

	for _, val := range pAns {
		set[val] = true
	}

	for index := len(qAns) - 1; index >= 0; index-- {
		if set[qAns[index]] {
			return qAns[index]
		}
	}

	return nil
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
	result := lowestCommonAncestor(root, root, root)

	fmt.Println(result)
}
