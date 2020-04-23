package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func dfs(before *ListNode, cur *ListNode, next *ListNode) {
	if next != nil && cur.Val == next.Val {
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}
		cur = next
		before.Next = cur
	}
	if cur == nil || cur.Next == nil {
		return
	}
	if cur.Val == cur.Next.Val {
		dfs(before, cur, cur.Next)
	} else {
		dfs(cur, cur.Next, cur.Next.Next)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	before := &ListNode{Val: -1, Next: head}
	if head.Next != nil {
		dfs(before, head, head.Next)
	}
	return before.Next
}

func main() {
	val1 := &ListNode{Val: 1, Next: nil}
	val2 := &ListNode{Val: 1, Next: nil}
	val3 := &ListNode{Val: 1, Next: nil}
	val4 := &ListNode{Val: 2, Next: nil}
	val5 := &ListNode{Val: 3, Next: nil}
	val1.Next = val2
	val2.Next = val3
	val3.Next = val4
	val4.Next = val5
	head := deleteDuplicates(val1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	val1 = &ListNode{Val: 1, Next: nil}
	val2 = &ListNode{Val: 2, Next: nil}
	val3 = &ListNode{Val: 2, Next: nil}
	val1.Next = val2
	val2.Next = val3
	head = deleteDuplicates(val1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
