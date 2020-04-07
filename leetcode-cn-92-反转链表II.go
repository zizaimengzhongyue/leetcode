package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	tmpHead := &ListNode{Val: 0, Next: head}
	index := 0
	cur := tmpHead
	tmpStart := cur
	tmpEnd := cur
	var last *ListNode
	last = nil
	for cur != nil {
		tmp := cur.Next
		if index == m-1 {
			tmpStart = cur
			tmpEnd = cur.Next
		}
		if index >= m && index <= n {
			cur.Next = last
		}
		if index == n {
			tmpStart.Next = cur
			if tmpEnd != nil {
				tmpEnd.Next = tmp
			}
		}
		last = cur
		cur = tmp
		index++
	}
	return tmpHead.Next
}

func main() {
	//	val1 := &ListNode{Val: 1, Next: nil}
	//	val2 := &ListNode{Val: 2, Next: nil}
	//	val3 := &ListNode{Val: 3, Next: nil}
	//	val4 := &ListNode{Val: 4, Next: nil}
	//	val5 := &ListNode{Val: 5, Next: nil}
	//	val1.Next = val2
	//	val2.Next = val3
	//	val3.Next = val4
	//	val4.Next = val5
	//	ans := reverseBetween(val1, 2, 4)
	//	for ans != nil {
	//		fmt.Println(ans.Val)
	//		ans = ans.Next
	//	}
	val1 := &ListNode{Val: 1, Next: nil}
	val2 := &ListNode{Val: 2, Next: nil}
	val1.Next = val2
	ans := reverseBetween(val1, 1, 2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}
