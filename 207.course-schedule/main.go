package main

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	degrees := map[int]int{}
	nextList := map[int][]int{}
	for i := 0; i < len(prerequisites); i++ {
		pre := prerequisites[i]
		course, req := pre[0], pre[1]
		degrees[course]++
		if _, ok := nextList[req]; ok {
			nextList[req] = append(nextList[req], course)
		} else {
			nextList[req] = []int{course}
		}
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if degrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		tmp := []int{}
		for i := 0; i < len(queue); i++ {
			numCourses--
			for _, nextCourse := range nextList[queue[i]] {
				degrees[nextCourse]--
				if degrees[nextCourse] == 0 {
					tmp = append(tmp, nextCourse)
				}
			}
		}
		queue = tmp
	}

	return numCourses == 0
}

// @lc code=end
