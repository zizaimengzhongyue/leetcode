package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := dfs(root.Left)
	rh := dfs(root.Right)
	if lh == -1 || rh == -1 || lh-rh > 1 || rh-lh > 1 {
		return -1
	} else {
		if lh > rh {
			return lh + 1
		} else {
			return rh + 1
		}
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root) != -1
}

func main() {
	n1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 := &TreeNode{Val: 20, Left: nil, Right: nil}
	n4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	n5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	fmt.Println(isBalanced(n1))

	n1 = &TreeNode{Val: 3, Left: nil, Right: nil}
	n2 = &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 = &TreeNode{Val: 20, Left: nil, Right: nil}
	n4 = &TreeNode{Val: 15, Left: nil, Right: nil}
	n5 = &TreeNode{Val: 7, Left: nil, Right: nil}
	n1.Left = n2
	n2.Left = n3
	n3.Left = n4
	n4.Left = n5
	fmt.Println(isBalanced(n1))

	n1 = &TreeNode{Val: 3, Left: nil, Right: nil}
	n2 = &TreeNode{Val: 9, Left: nil, Right: nil}
	n3 = &TreeNode{Val: 20, Left: nil, Right: nil}
	n1.Right = n2
	n2.Right = n3
	fmt.Println(isBalanced(n1))
}
