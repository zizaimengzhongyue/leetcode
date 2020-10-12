package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, tmp []int, sum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	sum = sum - root.Val
	nTmp := append([]int{}, tmp...)
	nTmp = append(nTmp, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		return append(ans, nTmp)
	}
	if root.Left != nil {
		ans = append(ans, dfs(root.Left, nTmp, sum)...)
	}
	if root.Right != nil {
		ans = append(ans, dfs(root.Right, nTmp, sum)...)
	}
	return ans
}

func pathSum(root *TreeNode, sum int) [][]int {
	return dfs(root, []int{}, sum)
}

func main() {
	val1 := &TreeNode{Val: 5, Left: nil, Right: nil}
	val2 := &TreeNode{Val: 4, Left: nil, Right: nil}
	val3 := &TreeNode{Val: 8, Left: nil, Right: nil}
	val4 := &TreeNode{Val: 11, Left: nil, Right: nil}
	val5 := &TreeNode{Val: 13, Left: nil, Right: nil}
	val6 := &TreeNode{Val: 4, Left: nil, Right: nil}
	val7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	val8 := &TreeNode{Val: 2, Left: nil, Right: nil}
	val9 := &TreeNode{Val: 5, Left: nil, Right: nil}
	val10 := &TreeNode{Val: 1, Left: nil, Right: nil}
	val1.Left = val2
	val1.Right = val3
	val2.Left = val4
	val3.Left = val5
	val3.Right = val6
	val4.Left = val7
	val4.Right = val8
	val6.Left = val9
	val6.Right = val10
	fmt.Println(pathSum(val1, 22))
}
