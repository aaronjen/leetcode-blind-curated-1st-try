package main

import "sort"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	curStart, curEnd := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		itvStart, itvEnd := intervals[i][0], intervals[i][1]
		if curEnd >= itvStart {
			if curEnd < itvEnd {
				curEnd = itvEnd
			}
		} else {
			res = append(res, []int{curStart, curEnd})
			curStart = itvStart
			curEnd = itvEnd
		}
	}
	res = append(res, []int{curStart, curEnd})

	return res
}

// @lc code=end
