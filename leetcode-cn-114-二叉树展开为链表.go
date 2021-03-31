package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLast(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	if root.Left == nil {
		return
	}
	last := getLast(root.Left)
	last.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}

func output(root *TreeNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Right
	}
	fmt.Println()
}

func main() {
	root := &TreeNode{Val: 1}
	a1 := &TreeNode{Val: 2}
	a2 := &TreeNode{Val: 3}
	a3 := &TreeNode{Val: 4}
	a4 := &TreeNode{Val: 5}
	a5 := &TreeNode{Val: 6}
	root.Left = a1
	root.Right = a4
	a1.Left = a2
	a1.Right = a3
	a4.Right = a5
	flatten(root)
	output(root)
}
