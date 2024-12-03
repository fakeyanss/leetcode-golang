/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (36.58%)
 * Likes:    5530
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 3.4M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
 * k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
 *
 * 你返回所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 解释：
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
 * 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
 * 注意，输出的顺序和三元组的顺序并不重要。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,1]
 * 输出：[]
 * 解释：唯一可能的三元组和不为 0 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0,0,0]
 * 输出：[[0,0,0]]
 * 解释：唯一可能的三元组和为 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */
package lc15

import "sort"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
//
// 排序+双指针问题
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	var res [][]int

	// 基本边界条件
	if n < 2 || len(nums)-start < n {
		return res
	}

	// Base case: 2Sum
	if n == 2 {
		l, r := start, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[l], nums[r]})
				l++
				r--
				// 跳过重复元素
				for ; l < r && nums[l] == nums[l-1]; l++ {
				}
				for ; l < r && nums[r] == nums[r+1]; r-- {
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
		return res
	}

	// Recursive case: n > 2
	for i := start; i < len(nums); i++ {
		// 跳过重复元素
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		subRes := nSumTarget(nums, n-1, i+1, target-nums[i])
		for _, subset := range subRes {
			// 将当前数字加入 (n-1) 的结果中
			res = append(res, append(subset, nums[i]))
		}
	}

	return res
}

// @lc code=end
