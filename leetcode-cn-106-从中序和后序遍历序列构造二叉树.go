package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	indexI := -1
	indexJ := -1
	for j := len(postorder) - 1; j >= 0; j-- {
		for i := 0; i < len(inorder); i++ {
			if postorder[j] == inorder[i] {
				indexI = i
				indexJ = j
				goto outer
			}
		}
	}
	return nil
outer:
	root := &TreeNode{Val: postorder[indexJ], Left: nil, Right: nil}
	root.Left = buildTree(inorder[0:indexI], postorder[0:indexJ])
	if indexI+1 < len(inorder) {
		root.Right = buildTree(inorder[indexI+1:], postorder[0:indexJ])
	}
	return root
}

func output(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	if root.Left != nil {
		output(root.Left)
	}
	if root.Right != nil {
		output(root.Right)
	}
}

func main() {
	nums1 := []int{9, 3, 15, 20, 7}
	nums2 := []int{9, 15, 7, 20, 3}
	output(buildTree(nums1, nums2))
}
