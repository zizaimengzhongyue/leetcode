package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	sum = sum - root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func main() {
	node1 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 8, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 11, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 13, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node8 := &TreeNode{Val: 2, Left: nil, Right: nil}
	node9 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node3.Left = node5
	node3.Right = node6
	node4.Left = node7
	node4.Right = node8
	node5.Right = node9
	fmt.Println(hasPathSum(node1, 22))
}
