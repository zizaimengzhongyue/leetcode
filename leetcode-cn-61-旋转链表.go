package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	ans := ListNode{}

	ln := 0
	tmp := head
	for tmp != nil {
		ln++
		tmp = tmp.Next
	}
	if ln == 0 {
		return head
	}
	newK := ln - k%ln

	tmp = head
	index := 0
	for tmp != nil {
		if index == ln-1 {
			tmp.Next = head
		}
		if index == newK {
			ans = *tmp
		}
		ttmp := tmp.Next
		if index == newK-1 {
			tmp.Next = nil
		}
		tmp = ttmp
		index++
	}
	return &ans
}

func main() {
	val5 := &ListNode{Val: 5, Next: nil}
	val4 := &ListNode{Val: 4, Next: val5}
	val3 := &ListNode{Val: 3, Next: val4}
	val2 := &ListNode{Val: 2, Next: val3}
	val1 := &ListNode{Val: 1, Next: val2}
	ans := rotateRight(val1, 2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
	fmt.Println()
	val3 = &ListNode{Val: 3, Next: nil}
	val2 = &ListNode{Val: 2, Next: val3}
	val1 = &ListNode{Val: 1, Next: val2}
	ans = rotateRight(val1, 4)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
	fmt.Println()
	val3 = &ListNode{Val: 3, Next: nil}
	val2 = &ListNode{Val: 2, Next: val3}
	val1 = &ListNode{Val: 1, Next: val2}
	ans = rotateRight(val1, 0)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
	fmt.Println()
	val1 = &ListNode{}
	ans = rotateRight(val1, 4)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
	fmt.Println()
}
