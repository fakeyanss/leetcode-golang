/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode.cn/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.67%)
 * Likes:    656
 * Dislikes: 0
 * Total Accepted:    437.1K
 * Total Submissions: 785.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */
package lc912

import "math/rand"

// @lc code=start
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	i := partition(nums)
	sortArray(nums[:i])
	sortArray(nums[i+1:])
	return nums
}

func partition(nums []int) int {
	// 1. 随机基准pivot
	n := len(nums)
	i := rand.Intn(n)
	pivot := nums[i]
	nums[i], nums[0] = nums[0], nums[i]

	// 2. 双指针遍历，小的放左边，大的放右边
	i, j := 1, n-1
	for {
		for i <= j && nums[i] < pivot { // 找到比pivot大的nums[i]
			i++
		}
		for i <= j && nums[j] > pivot { // 找到比pivot小的nums[j]
			j--
		}
		if i >= j { // 如果i，j相遇则已经交换完
			break
		}
		nums[i], nums[j] = nums[j], nums[i] // 交换
		i++
		j--
	}
	// 临界情况j可能为0，i可能为n
	// 此时j的位置是小于等于pivot的元素下标
	nums[0], nums[j] = nums[j], nums[0]
	return j
}

// @lc code=end
