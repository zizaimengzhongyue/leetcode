package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		root.Left.Val += root.Val * 10
		dfs(root.Left, sum)
	}
	if root.Right != nil {
		root.Right.Val += root.Val * 10
		dfs(root.Right, sum)
	}
	if root.Left == nil && root.Right == nil {
		(*sum) += root.Val
	}
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	dfs(root, &sum)
	return sum
}

func main() {
	root := &TreeNode{Val: 4}
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 0}
	lleft := &TreeNode{Val: 5}
	lright := &TreeNode{Val: 1}
	root.Left = left
	root.Right = right
	root.Left.Left = lleft
	root.Left.Right = lright
	fmt.Println(sumNumbers(root))
}
