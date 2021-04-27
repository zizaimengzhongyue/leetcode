package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l != 0 && r == 0 {
		return l + 1
	}
	if l == 0 && r != 0 {
		return r + 1
	}
	if l < r {
		return l + 1
	}
	return r + 1
}

func main() {
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 9}
	n3 := &TreeNode{Val: 20}
	n4 := &TreeNode{Val: 15}
	n5 := &TreeNode{Val: 7}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	fmt.Println(minDepth(n1))
}
