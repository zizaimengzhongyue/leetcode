package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	list := []*TreeNode{root}
	ans := [][]int{}
	head := 0
	tail := 1
	for head < tail {
		tmp := []int{}
		ln := tail - head
		for i := 0; i < ln; i++ {
			cur := list[head]
			head++
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				list = append(list, cur.Left)
				tail++
			}
			if cur.Right != nil {
				list = append(list, cur.Right)
				tail++
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	node1 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 20, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5
	fmt.Println(levelOrder(node1))
}
