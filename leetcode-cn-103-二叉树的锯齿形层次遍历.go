package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	list := []*TreeNode{root}
	head := 0
	tail := 1
	ans := [][]int{}
	target := 1
	for head < tail {
		size := tail - head
		tmp := make([]int, size, size)
		if target%2 == 0 {
			target++
			for i := size - 1; i >= 0; i-- {
				cur := list[head]
				head++
				tmp[i] = cur.Val
				if cur.Left != nil {
					list = append(list, cur.Left)
					tail++
				}
				if cur.Right != nil {
					list = append(list, cur.Right)
					tail++
				}
			}
		} else {
			target++
			for i := 0; i < size; i++ {
				cur := list[head]
				head++
				tmp[i] = cur.Val
				if cur.Left != nil {
					list = append(list, cur.Left)
					tail++
				}
				if cur.Right != nil {
					list = append(list, cur.Right)
					tail++
				}
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

func main() {
	root := &TreeNode{Val: 3, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 20, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Left = node1
	root.Right = node2
	node2.Left = node3
	node2.Right = node4
	fmt.Println(zigzagLevelOrder(root))

}
