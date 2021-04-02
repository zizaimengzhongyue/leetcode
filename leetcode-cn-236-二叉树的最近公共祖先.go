package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	var l *TreeNode
	var r *TreeNode
	if root.Left != nil {
		l = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Right != nil {
		r = lowestCommonAncestor(root.Right, p, q)
	}
	if l != nil && r == nil {
		return l
	}
	if r != nil && l == nil {
		return r
	}
	if l != nil && r != nil {
		return root
	}
	return nil
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node5.Left = node6
	fmt.Println(lowestCommonAncestor(root, node4, node6))
}
