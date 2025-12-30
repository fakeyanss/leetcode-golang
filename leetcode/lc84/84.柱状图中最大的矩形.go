/*
 * @lc app=leetcode.cn id=84 lang=golang
 * @lcpr version=30305
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (48.57%)
 * Likes:    3038
 * Dislikes: 0
 * Total Accepted:    613.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,1,5,6,2,3]\n[2,4]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 * 输入：heights = [2,1,5,6,2,3]
 * 输出：10
 * 解释：最大的矩形为图中红色区域，面积为 10
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入： heights = [2,4]
 * 输出： 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= heights.length <=10^5
 * 0 <= heights[i] <= 10^4
 *
 *
 */
package lc84

// @lc code=start
func largestRectangleArea(heights []int) int {
	n := len(heights)
	// left[i]存储左侧第一个比left[i]小的元素idx
	left, right := make([]int, n), make([]int, n)
	var stk []int

	for i := 0; i < n; i++ {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		left[i] = -1 // 左侧没有更小元素，默认值
		if len(stk) > 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i] {
			stk = stk[:len(stk)-1]
		}
		right[i] = n // 右侧没有更小元素，默认值
		if len(stk) > 0 {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	var res int
	for i := range n {
		res = max(res, heights[i]*(right[i]-left[i]-1))
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,5,6,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n
// @lcpr case=end

*/
