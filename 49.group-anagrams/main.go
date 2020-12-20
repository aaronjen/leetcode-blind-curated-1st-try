package main

import "sort"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	table := map[string][]string{}

	for i := 0; i < len(strs); i++ {
		s := strs[i]
		ar := []byte(s)

		sort.Slice(ar, func(i, j int) bool {
			return ar[i] < ar[j]
		})

		key := string(ar)
		if _, ok := table[key]; ok {
			table[key] = append(table[key], s)
		} else {
			table[key] = []string{s}
		}
	}

	res := [][]string{}
	for _, stringArr := range table {
		res = append(res, stringArr)
	}

	return res
}

// @lc code=end
