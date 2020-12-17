package main

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}

	table := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		b := s[i]
		if v, ok := table[b]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0
}

// @lc code=end
