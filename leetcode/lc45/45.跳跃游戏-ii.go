/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode.cn/problems/jump-game-ii/description/
 *
 * algorithms
 * Medium (45.16%)
 * Likes:    1821
 * Dislikes: 0
 * Total Accepted:    406.7K
 * Total Submissions: 900.5K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 *
 * 假设你总是可以到达数组的最后一个位置。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [2,3,0,1,4]
 * 输出: 2
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * 0
 *
 *
 */
package lc45

// @lc code=start
func jump(nums []int) int {
	// return dp(nums)
	return greedy(nums)
}

func dp(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = n
	}
	// 从索引 p 跳到最后一格，至少需要 f(nums, p) 步
	var f func([]int, int) int
	f = func(nums []int, p int) int {
		// base case
		if p >= n-1 {
			return 0
		}
		if memo[p] != n {
			return memo[p]
		}
		steps := nums[p]
		for i := 1; i <= steps; i++ {
			// 穷举每一个选择
			subProblem := f(nums, p+i)
			memo[p] = minInt(memo[p], subProblem+1)
		}
		return memo[p]
	}
	return f(nums, 0)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func greedy(nums []int) int {
	n := len(nums)
	end, farthest := 0, 0
	jumps := 0
	// i 和 end 标记了可以选择的跳跃步数，
	// farthest 标记了所有选择 [i..end] 中能够跳到的最远距离，
	// jumps 记录了跳跃次数
	for i := 0; i < n-1; i++ {
		farthest = maxInt(farthest, nums[i]+i)
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
