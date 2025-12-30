/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode.cn/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (69.57%)
 * Likes:    1339
 * Dislikes: 0
 * Total Accepted:    373K
 * Total Submissions: 536.2K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i
 * 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: temperatures = [73,74,75,71,69,72,76,73]
 * 输出: [1,1,4,2,1,1,0,0]
 *
 *
 * 示例 2:
 *
 *
 * 输入: temperatures = [30,40,50,60]
 * 输出: [1,1,1,0]
 *
 *
 * 示例 3:
 *
 *
 * 输入: temperatures = [30,60,90]
 * 输出: [1,1,0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= temperatures.length <= 10^5
 * 30 <= temperatures[i] <= 100
 *
 *
 */
package lc739

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stk []int // 存储温度的idx
	for i := range n {
		temp := temperatures[i]
		// 弹出栈顶所以比当前小的天数，弹出的值都可以更新结果
		for len(stk) > 0 && temp > temperatures[stk[len(stk)-1]] {
			prevIdx := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res[prevIdx] = i - prevIdx
		}
		stk = append(stk, i)
	}
	return res
}

// @lc code=end
