/*
 * @lc app=leetcode.cn id=16 lang=golang
 * @lcpr version=20003
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.79%)
 * Likes:    1675
 * Dislikes: 0
 * Total Accepted:    597.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 *
 * 返回这三个数的和。
 *
 * 假定每组输入只存在恰好一个解。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 * 解释：与 target 最接近的和是 0（0 + 0 + 0 = 0）。
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */
package lc16

import (
	"math"
	"sort"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	// 别忘了要先排序数组
	sort.Ints(nums)
	// 记录三数之和与目标值的偏差
	delta := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		// 固定 nums[i] 为三数之和中的第一个数，
		// 然后对 nums[i+1..] 搜索接近 target - nums[i] 的两数之和
		sum := nums[i] + twoSumClosest(nums, i+1, target-nums[i])
		if math.Abs(float64(delta)) > math.Abs(float64(target-sum)) {
			delta = target - sum
		}
	}
	return target - delta
}

// 在 nums[start..] 搜索最接近 target 的两数之和
func twoSumClosest(nums []int, start int, target int) int {
	lo, hi := start, len(nums)-1
	// 记录两数之和与目标值的偏差
	delta := math.MaxInt32
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if math.Abs(float64(delta)) > math.Abs(float64(target-sum)) {
			delta = target - sum
		}
		if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return target - delta
}

// func threeSumClosest(nums []int, target int) int {
//     sort.Ints(nums)
// 	return nSumTarget(nums, 3, 0, target)
// }

// // 排序+双指针问题
// func nSumTarget(nums []int, n, start, target int) int {
// 	res := math.MaxInt32

// 	// 基本边界条件
// 	if n < 2 || len(nums)-start < n {
// 		return res
// 	}

// 	// Base case: 2Sum
// 	if n == 2 {
// 		l, r := start, len(nums)-1
// 		for l < r {
// 			sum := nums[l] + nums[r]
// 			if sum == target {
//                 return res
// 			} else {
//                 if math.Abs(float64(res-target)) > math.Abs(float64(sum-target)) {
//                     res = sum
//                 }
//                 if sum < target {
//                     l++
//                 } else {
//                     r--
//                 }
// 			}
// 		}
// 		return res
// 	}

// 	// Recursive case: n > 2
// 	for i := start; i < len(nums); i++ {
// 		// 跳过重复元素
// 		if i > start && nums[i] == nums[i-1] {
// 			continue
// 		}
// 		subRes := nSumTarget(nums, n-1, i+1, target-nums[i])
// 		if math.Abs(float64(nums[i] + subRes - target)) < math.Abs(float64(res - target)) {
//             res = nums[i] + subRes
//         }
// 	}

// 	return res
// }

// @lc code=end

/*
// @lcpr case=start
// [-1,2,1,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n1\n
// @lcpr case=end

*/
