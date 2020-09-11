package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	indexPre := -1
	indexIn := -1
	for i := 0; i < len(preorder); i++ {
		for j := 0; j < len(inorder); j++ {
			if inorder[j] == preorder[i] {
				indexPre = i
				indexIn = j
				goto outer
			}
		}
	}
outer:
	if indexPre == -1 || indexIn == -1 {
		return nil
	}
	root := &TreeNode{Val: preorder[indexPre], Left: nil, Right: nil}
	if indexPre+1 < len(preorder) {
		root.Left = buildTree(preorder[indexPre+1:], inorder[0:indexIn])
	}
	if indexPre+1 < len(preorder) && indexIn+1 < len(inorder) {
		root.Right = buildTree(preorder[indexPre+1:], inorder[indexIn+1:])
	}
	return root
}

func Print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	Print(root.Left)
	Print(root.Right)
}

func main() {
	nums1 := []int{3, 9, 20, 15, 7}
	nums2 := []int{9, 3, 15, 20, 7}
	root := buildTree(nums1, nums2)
	Print(root)
	nums1 = []int{1, 2}
	nums2 = []int{1, 2}
	root = buildTree(nums1, nums2)
	Print(root)
}
