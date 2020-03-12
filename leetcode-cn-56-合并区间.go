package main

import (
    "fmt"
    "sort"
)

func merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    for i := 0; i < len(intervals); i++ {
        if intervals[i][0] > intervals[i][1] {
            intervals[i][0], intervals[i][1] = intervals[i][1], intervals[i][0]
        }
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
    })
    ans := [][]int{}
    start := intervals[0][0]
    end := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= end {
            if intervals[i][1] > end {
                end = intervals[i][1]
            }
        } else {
            ans = append(ans, []int{start, end})
            start = intervals[i][0]
            end = intervals[i][1]
        }
    }
    ans = append(ans, []int{start, end})
    return ans
}

func main() {
    nums := [][]int{{1,3}, {2,6}, {8,10}, {15,18}}
    nums = [][]int{{1,4}, {4,5}}
    nums = [][]int{}
    fmt.Println(merge(nums))
}

//[[1,3],[2,6],[8,10],[15,18]]
//[[1,6],[8,10],[15,18]]


//[[1,4],[4,5]]
//[[1,5]]
