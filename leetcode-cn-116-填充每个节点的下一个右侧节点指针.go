package main

import "fmt"

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

func output(root *Node) {
	if root == nil {
		return
	}
	tmp := root
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
	output(root.Left)
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	output(connect(n1))
}
