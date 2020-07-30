package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func cmp(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil || left != nil && right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	if left.Left != nil && right.Right == nil || left.Left == nil && right.Right != nil {
		return false
	}
	if left.Right != nil && right.Left == nil || left.Right == nil && right.Left != nil {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !cmp(root.Left, root.Right) {
		return false
	}
	if root.Left == nil {
		return true
	}
	leftQ := []*TreeNode{root.Left}
	rightQ := []*TreeNode{root.Right}
	head := 0
	end := 1
	for head < end {
		left := leftQ[head]
		right := rightQ[head]
		head++
		if !cmp(left.Left, right.Right) {
			return false
		}
		if left.Left != nil {
			leftQ = append(leftQ, left.Left)
			rightQ = append(rightQ, right.Right)
			end++
		}
		if !cmp(left.Right, right.Left) {
			return false
		}
		if left.Right != nil {
			leftQ = append(leftQ, left.Right)
			rightQ = append(rightQ, right.Left)
			end++
		}
	}
	return true
}

func main() {
	val1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	val2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	val3 := &TreeNode{Val: 2, Left: nil, Right: nil}
	val4 := &TreeNode{Val: 3, Left: nil, Right: nil}
	val5 := &TreeNode{Val: 4, Left: nil, Right: nil}
	val6 := &TreeNode{Val: 4, Left: nil, Right: nil}
	val7 := &TreeNode{Val: 3, Left: nil, Right: nil}
	val1.Left = val2
	val1.Right = val3
	val2.Left = val4
	val2.Right = val5
	val3.Left = val6
	val3.Right = val7
	fmt.Println(isSymmetric(val1))
}
