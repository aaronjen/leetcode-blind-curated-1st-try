package main

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	backtrack(candidates, target, []int{}, &res, 0)

	return res
}

func backtrack(candidates []int, target int, tmp []int, res *[][]int, ind int) {
	if target < 0 {
		return
	}

	if target == 0 {
		tt := make([]int, len(tmp))
		copy(tt, tmp)
		*res = append(*res, tt)
	} else {
		for i := ind; i < len(candidates); i++ {
			tmp = append(tmp, candidates[i])
			backtrack(candidates, target-candidates[i], tmp, res, i)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

// @lc code=end
