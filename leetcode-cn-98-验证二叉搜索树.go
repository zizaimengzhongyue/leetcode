package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, cur int) (bool, int) {
	if root == nil {
		return true, cur
	}
	ans := true
	if root.Left != nil {
		ans, cur = dfs(root.Left, cur)
		if !ans {
			return ans, cur
		}
	}
	if root.Val <= cur {
		return false, root.Val
	}
	cur = root.Val
	if root.Right != nil {
		ans, cur = dfs(root.Right, cur)
		if !ans {
			return ans, cur
		}
	}
	return ans, cur
}

func isValidBST(root *TreeNode) bool {
	ans, _ := dfs(root, ^int(^uint(0)>>1))
	return ans
}

func main() {
	val3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	val2 := &TreeNode{Val: 1, Left: nil, Right: nil}
	val1 := &TreeNode{Val: 2, Left: val2, Right: val3}
	fmt.Println(isValidBST(val1))
	val5 := &TreeNode{Val: 3, Left: nil, Right: nil}
	val4 := &TreeNode{Val: 6, Left: nil, Right: nil}
	val3 = &TreeNode{Val: 4, Left: val5, Right: val4}
	val2 = &TreeNode{Val: 1, Left: nil, Right: nil}
	val1 = &TreeNode{Val: 5, Left: val2, Right: val3}
	fmt.Println(isValidBST(val1))
	val5 = &TreeNode{Val: 20, Left: nil, Right: nil}
	val4 = &TreeNode{Val: 6, Left: nil, Right: nil}
	val3 = &TreeNode{Val: 15, Left: val4, Right: val5}
	val2 = &TreeNode{Val: 5, Left: nil, Right: nil}
	val1 = &TreeNode{Val: 10, Left: val2, Right: val3}
	fmt.Println(isValidBST(val1))
}
