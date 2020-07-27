package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func getList(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    nums := []int{}
    queue := []*TreeNode{}
    head := 0
    end := 0
    queue = append(queue, root)
    end = 1
    for head != end {
        cur := queue[head]
        head++
        nums = append(nums, cur.Val)
        if cur.Left != nil {
            queue = append(queue, cur.Left)
            end++
        }
        if cur.Right != nil {
            queue = append(queue, cur.Right)
            end++
        }
    }
    return nums
}

func f(nums []int) bool {
    ln := len(nums)
    for i:=1;i<=(ln-1-i);i++ {
        if nums[i] != nums[ln-1-i] {
            return false
        }
    }
    return true
}

func check(nums []int) bool {
    ln := len(nums)
    begin := 0
    step := 1
    for i := 1;begin + step<ln;i++ {
        if !f(nums[begin:begin + step]) {
            return false
        }
        begin = begin + step
        step = step * 2
    }
    return true
}

func isSymmetric(root *TreeNode) bool {
    nums := getList(root)
    return check(nums)
}

func main() {
    val1 := &TreeNode{Val: 1, Left: nil, Right: nil}
    val2 := &TreeNode{Val: 2, Left: nil, Right: nil}
    val3 := &TreeNode{Val: 2, Left: nil, Right: nil}
    val4 := &TreeNode{Val: 3, Left: nil, Right: nil}
    val5 := &TreeNode{Val: 4, Left: nil, Right: nil}
    val6 := &TreeNode{Val: 4, Left: nil, Right: nil}
    val7 := &TreeNode{Val: 3, Left: nil, Right: nil}
    val1.Left = val2
    val1.Right = val3
    val2.Left  = val4
    val2.Right = val5
    val3.Left = val6
    val3.Right = val7
    fmt.Println(isSymmetric(val1))
}
