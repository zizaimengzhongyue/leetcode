package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	ln := len(arr)
	mid := ln / 2
	left := mid - 1
	right := mid + 1
	root := &TreeNode{Val: arr[mid]}
	if left >= 0 {
		root.Left = buildTree(arr[0:mid])
	}
	if right < ln {
		root.Right = buildTree(arr[right:])
	}
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return buildTree(arr)
}

func outputList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func outputTreeNode(head *TreeNode) {
	if head == nil {
		return
	}
	if head.Left != nil {
		outputTreeNode(head.Left)
	}
	fmt.Printf("%d ", head.Val)
	if head.Right != nil {
		outputTreeNode(head.Right)
	}
}

func main() {
	a := []*ListNode{}
	for i := 0; i < 5; i++ {
		tmp := &ListNode{Val: i}
		if i-1 >= 0 {
			a[i-1].Next = tmp
		}
		a = append(a, tmp)
	}
	outputTreeNode(sortedListToBST(a[0]))
	fmt.Println()
}
