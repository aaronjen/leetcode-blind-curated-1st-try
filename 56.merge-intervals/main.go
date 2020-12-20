package main

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	curStart := intervals[0][0]
	curEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		itv := intervals[i]
		start, end := itv[0], itv[1]
		if start <= curEnd {
			if end > curEnd {
				curEnd = end
			}
		} else {
			res = append(res, []int{curStart, curEnd})
			curStart = start
			curEnd = end
		}
	}
	res = append(res, []int{curStart, curEnd})

	return res
}

// @lc code=end
