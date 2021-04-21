package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[0:mid])
	if len(nums) >= mid+1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}

func output(root *TreeNode) {
	if root == nil {
		return
	}
	output(root.Left)
	fmt.Println(root.Val)
	output(root.Right)
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	output(sortedArrayToBST(nums))
}
