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

// @lc code=start
//
// 排序+双指针问题
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

func nSumTarget(nums []int, n, start, target int) [][]int {
	ans := [][]int{}
	if n < 2 || len(nums) < n {
		return ans
	}
	if n == 2 {
		// 2Sum 是 base case
		lo, hi := start, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi]
			l, r := nums[lo], nums[hi]
			if sum < target {
				// 过滤重复值
				for ; lo < hi && nums[lo] == l; lo++ {
				}
			} else if sum > target {
				// 过滤重复值
				for ; lo < hi && nums[hi] == r; hi-- {
				}
			} else {
				ans = append(ans, []int{l, r})
				// 过滤重复值
				for ; lo < hi && nums[lo] == l; lo++ {
				}
				// 过滤重复值
				for ; lo < hi && nums[hi] == r; hi-- {
				}
			}
		}
	} else {
		// n > 2 时，递归计算 (n-1)Sum 的结果
		for i := start; i < len(nums); i++ {
			sub := nSumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				// (n-1)Sum 加上 nums[i] 就是 nSum
				arr = append(arr, nums[i])
				ans = append(ans, arr)
			}
			// 跳过重复值
			for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
			}
		}
	}
	return ans
}

// @lc code=end
