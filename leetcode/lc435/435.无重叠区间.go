/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 *
 * https://leetcode.cn/problems/non-overlapping-intervals/description/
 *
 * algorithms
 * Medium (51.10%)
 * Likes:    811
 * Dislikes: 0
 * Total Accepted:    178.9K
 * Total Submissions: 350.1K
 * Testcase Example:  '[[1,2],[2,3],[3,4],[1,3]]'
 *
 * 给定一个区间的集合 intervals ，其中 intervals[i] = [starti, endi] 。返回
 * 需要移除区间的最小数量，使剩余区间互不重叠 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: intervals = [[1,2],[2,3],[3,4],[1,3]]
 * 输出: 1
 * 解释: 移除 [1,3] 后，剩下的区间没有重叠。
 *
 *
 * 示例 2:
 *
 *
 * 输入: intervals = [ [1,2], [1,2], [1,2] ]
 * 输出: 2
 * 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
 *
 *
 * 示例 3:
 *
 *
 * 输入: intervals = [ [1,2], [2,3] ]
 * 输出: 0
 * 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= intervals.length <= 10^5
 * intervals[i].length == 2
 * -5 * 10^4 <= starti < endi <= 5 * 10^4
 *
 *
 */
package lc435

import "sort"

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	return len(intervals) - intervalSchedule(intervals)
}

// 区间调度算法，计算intvs中最多有几个互不相交的区间
func intervalSchedule(intvs [][]int) int {
	n := len(intvs)
	if n == 0 {
		return 0
	}
	// 按 end 升序排序
	sort.Slice(intvs, func(i, j int) bool {
		return intvs[i][1] < intvs[j][1]
	})
	// 每次选择end最早的区间x，并将与x相交的区间移除掉
	count, xEnd := 1, intvs[0][1]
	for _, interval := range intvs {
		start := interval[0]
		if start >= xEnd {
			// 找到下一个选择的区间
			count++
			xEnd = interval[1]
		}
	}
	return count
}

// @lc code=end
