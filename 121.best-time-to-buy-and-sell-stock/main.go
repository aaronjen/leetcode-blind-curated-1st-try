package main

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	cost := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < cost {
			cost = prices[i]
		} else {
			val := prices[i] - cost
			if val > res {
				res = val
			}
		}
	}

	return res
}

// @lc code=end
