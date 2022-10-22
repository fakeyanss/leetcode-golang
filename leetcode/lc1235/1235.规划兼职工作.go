/*
 * @lc app=leetcode.cn id=1235 lang=golang
 *
 * [1235] 规划兼职工作
 *
 * https://leetcode.cn/problems/maximum-profit-in-job-scheduling/description/
 *
 * algorithms
 * Hard (48.72%)
 * Likes:    281
 * Dislikes: 0
 * Total Accepted:    22.4K
 * Total Submissions: 40.4K
 * Testcase Example:  '[1,2,3,3]\n[3,4,5,6]\n[50,10,40,70]'
 *
 * 你打算利用空闲时间来做兼职工作赚些零花钱。
 *
 * 这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
 *
 * 给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
 *
 * 注意，时间上出现重叠的 2 份工作不能同时进行。
 *
 * 如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
 * 输出：120
 * 解释：
 * 我们选出第 1 份和第 4 份工作，
 * 时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。
 *
 *
 * 示例 2：
 *
 * ⁠
 *
 * 输入：startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit =
 * [20,20,100,70,60]
 * 输出：150
 * 解释：
 * 我们选择第 1，4，5 份工作。
 * 共获得报酬 150 = 20 + 70 + 60。
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入：startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
 * 输出：6
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
 * 1 <= startTime[i] < endTime[i] <= 10^9
 * 1 <= profit[i] <= 10^4
 *
 *
 */
package lc1235

import "sort"

// @lc code=start
// 动态规划+二分搜索
// dp[i] 表示按照结束时间排序后的前i个工作可获取的最大报酬
// 状态转移：dp[i] = max(dp[i-1], dp[j]+profit[i])
// 解释：j为endTime临近的上一个状态；到第i个工作时，可以不选择即dp[i-1]，或选择即最近的状态dp[j]+工作i的报酬
// 问：是否可以贪心？
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// 以endtime正排
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	// 避免处理-1，下标从1开始
	dp := make([]int, n+1)
	for i, job := range jobs {
		j := binarySearchRight(jobs, i, job[0])
		dp[i+1] = max(dp[i], dp[j+1]+job[2])
	}
	return dp[n]
}

// 返回job[0..right]之间，endTime<=startTime的最大下标
func binarySearchRight(jobs [][3]int, right, startTime int) int {
	left := -1
	for left+1 < right {
		mid := left + (right-left)/2
		if jobs[mid][1] <= startTime {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
