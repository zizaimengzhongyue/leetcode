package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Wrapper struct {
	Depth int
	N     *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*Wrapper{&Wrapper{Depth: 0, N: root}}
	head := 0
	tail := 1
	for head != tail {
		cur := q[head]
		head++
		if head != tail && q[head].Depth == cur.Depth {
			cur.N.Next = q[head].N
		}
		if cur.N.Left != nil {
			q = append(q, &Wrapper{Depth: cur.Depth + 1, N: cur.N.Left})
			tail++
		}
		if cur.N.Right != nil {
			q = append(q, &Wrapper{Depth: cur.Depth + 1, N: cur.N.Right})
			tail++
		}
	}
	return root
}

func connect02(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else if root.Next != nil {
			root.Left.Next = getNext(root.Next)
		}
	}
	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = getNext(root.Next)
		}
	}
	connect02(root.Right)
	connect02(root.Left)
	return root
}

func getNext(root *Node) *Node {
	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}
	if root.Next != nil {
		return getNext(root.Next)
	}
	return nil
}

func output(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Val, root.Next)
	output(root.Left)
	output(root.Right)
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}
	n13 := &Node{Val: 13}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n4.Left = n8
	n5.Left = n9
	n5.Right = n10
	n7.Left = n11
	n7.Right = n12
	n11.Left = n13
	output(connect02(n1))
	output(connect(n1))
}
