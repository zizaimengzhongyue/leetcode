package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	if root.Left != nil {
		ans = dfs(root.Left, ans)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = dfs(root.Right, ans)
	}
	return ans
}

func inorderTraversal(root *TreeNode) []int {
	return dfs(root, []int{})
}

func main() {
	val1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	val2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	val3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	val1.Right = val2
	val2.Left = val3
	fmt.Println(inorderTraversal(val1))
}
