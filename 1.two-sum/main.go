package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	table := map[int]int{}

	for i, v := range nums {
		if prev, ok := table[v]; ok {
			return []int{prev, i}
		}

		table[target-v] = i
	}
	return []int{}
}

// @lc code=end
