package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func clone(node *Node, m map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := m[node.Val]; ok {
		return m[node.Val]
	}
	m[node.Val] = &Node{Val: node.Val, Neighbors: []*Node{}}
	for _, v := range node.Neighbors {
		m[node.Val].Neighbors = append(m[node.Val].Neighbors, clone(v, m))
	}
	return m[node.Val]
}

func cloneGraph(node *Node) *Node {
	return clone(node, map[int]*Node{})
}

func p(node *Node, mark map[int]bool) {
	if mark[node.Val] {
		return
	}
	fmt.Println(node.Val)
	mark[node.Val] = true
	for _, v := range node.Neighbors {
		p(v, mark)
	}
}

func output(node *Node) {
	p(node, map[int]bool{})
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n3)
	n3.Neighbors = append(n3.Neighbors, n4)
	n4.Neighbors = append(n4.Neighbors, n1)
	n := cloneGraph(n1)
	output(n)
}
