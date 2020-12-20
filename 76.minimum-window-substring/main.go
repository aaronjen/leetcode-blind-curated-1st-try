package main

import "math"

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	need := map[byte]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	count := 0
	for range need {
		count++
	}

	res := ""
	l := math.MaxInt32
	start := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		if _, ok := need[b]; ok {
			need[b]--
			if need[b] == 0 {
				count--
			}
		}

		for count == 0 {
			tmp := s[start : i+1]
			if len(tmp) < l {
				res = tmp
				l = len(res)
			}
			start++
			if _, ok := need[s[start-1]]; ok {
				need[s[start-1]]++
				if need[s[start-1]] > 0 {
					count++
				}
			}
		}
	}

	return res
}

// @lc code=end
