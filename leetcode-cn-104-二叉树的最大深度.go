package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Tree *TreeNode
	Deep int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deep := 0
	cur := &Node{Tree: root, Deep: 1}
	queue := []*Node{cur}
	for len(queue) != 0 {
		cur = queue[0]
		queue = queue[1:]
		if cur.Deep > deep {
			deep = cur.Deep
		}
		if cur.Tree.Left != nil {
			queue = append(queue, &Node{Tree: cur.Tree.Left, Deep: cur.Deep + 1})
		}
		if cur.Tree.Right != nil {
			queue = append(queue, &Node{Tree: cur.Tree.Right, Deep: cur.Deep + 1})
		}
	}
	return deep
}

func main() {
	root := &TreeNode{Val: 0, Left: nil, Right: nil}
	n1 := &TreeNode{Val: 1, Left: nil, Right: nil}
	n2 := &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = n1
	n1.Right = n2
	fmt.Println(maxDepth(root))
}
