package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(nums [][]int) [][]int {
	ans := [][]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		ans = append(ans, nums[i])
	}
	return ans
}

func merge(left [][]int, right [][]int) [][]int {
	ans := [][]int{}
	left = reverse(left)
	right = reverse(right)
	if len(left) < len(right) {
		for i := 0; i < len(left); i++ {
			ans = append(ans, append(left[i], right[i]...))
		}
		for i := len(left); i < len(right); i++ {
			ans = append(ans, right[i])
		}
	} else {
		for i := 0; i < len(right); i++ {
			ans = append(ans, append(left[i], right[i]...))
		}
		for i := len(right); i < len(left); i++ {
			ans = append(ans, left[i])
		}
	}
	return reverse(ans)
}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	ans = merge(ans, levelOrderBottom(root.Left))
	ans = merge(ans, levelOrderBottom(root.Right))
	return append(ans, []int{root.Val})
}

func main() {
	node := &TreeNode{Val: 3, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 20, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node.Left = node1
	node.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Println(levelOrderBottom(node))
}
