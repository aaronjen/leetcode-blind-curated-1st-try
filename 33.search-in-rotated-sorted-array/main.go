package main

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func bs(nums []int, target int, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		n := nums[mid]
		if n == target {
			return mid
		} else if n > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	pivot := 1
	for ; pivot < len(nums); pivot++ {
		if nums[pivot] < nums[pivot-1] {
			break
		}
	}
	pivot--

	res := bs(nums, target, 0, pivot)
	if res != -1 {
		return res
	}

	return bs(nums, target, pivot+1, len(nums)-1)
}

// @lc code=end
