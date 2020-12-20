package main

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}

// @lc code=end
