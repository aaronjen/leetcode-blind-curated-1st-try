package main

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	table := make([]int, 26)
	for i := 0; i < len(s); i++ {
		table[int(s[i]-'a')]++
	}

	for i := 0; i < len(t); i++ {
		table[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if table[i] != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
