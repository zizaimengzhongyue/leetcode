package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Wrapper struct {
	T     *TreeNode
	Index int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*Wrapper{}
	head := 0
	queue = append(queue, &Wrapper{T: root, Index: 0})
	tail := 1
	index := 0
	for head < tail {
		cur := queue[head]
		head++
		if cur.Index == 0 {
			ans = append(ans, cur.T.Val)
			index = 0
		}
		if cur.T.Right != nil {
			queue = append(queue, &Wrapper{T: cur.T.Right, Index: index})
			index++
			tail++
		}
		if cur.T.Left != nil {
			queue = append(queue, &Wrapper{T: cur.T.Left, Index: index})
			index++
			tail++
		}
	}
	return ans
}

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	root.Left = node2
	root.Right = node4
	node2.Right = node3
	node4.Left = node5
	fmt.Println(rightSideView(root))
}
