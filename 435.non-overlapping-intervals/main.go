package main

import "sort"

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	prev := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if start >= prev {
			count++
			prev = end
		}
	}
	return len(intervals) - count
}

// @lc code=end
