package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func fromArrayToBinaryTree(values []int) *TreeNode {
	var result *TreeNode = nil

	for index := 0; index < len(values); index++ {
		binarySearchTreeInsert(values[index], result)
	}

	return result
}

func binarySearchTreeInsert(value int, root *TreeNode) *TreeNode {
	// if value == nil {
	// 	return root
	// }

	var before, after *TreeNode = nil, root

	for root != nil {
		before = after

		if value < before.Val {
			after = before.Left
		}
		if value > before.Val {
			after = before.Right
		}

		if value == before.Val {
			return root
		}
	}

	var newNode *TreeNode = &TreeNode{Val: value}

	if root == nil {
		root = newNode
		return root
	} else {
		if value < before.Val {
			before.Left = newNode
		} else {
			before.Right = newNode
		}
	}

	return root
}
