package main

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	table := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		table[nums[i]] = true
	}

	res := 1
	for key, value := range table {
		if !value {
			continue
		}
		tmp := 1
		right := key + 1
		for table[right] {
			table[right] = false
			tmp++
			right++

		}
		left := key - 1
		for table[left] {
			table[left] = false
			tmp++
			left--
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}

// @lc code=end
