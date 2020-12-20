package main

/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func findmax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}

	res := 0

	table := make([]int, 26)
	max := 0
	start := 0
	for i := 0; i < len(s); i++ {
		table[s[i]-'A']++
		max = findmax(table[s[i]-'A'], max)
		l := i - start + 1
		if l-max <= k {
			if l > res {
				res = l
			}
		} else {
			table[s[start]-'A']--
			start++
		}
	}

	return res
}

// @lc code=end
