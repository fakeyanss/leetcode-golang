/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.54%)
 * Likes:    1861
 * Dislikes: 0
 * Total Accepted:    731.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4], k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6], k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
package lc

import "math/rand"

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// (快排)快速选择

	n := len(nums)
	// 打乱顺序，避免最坏情况
	var shuffle = func(nums []int) {
		for i := 0; i < n; i++ {
			r := rand.Intn(n)
			nums[i], nums[r] = nums[r], nums[i]
		}
	}
	shuffle(nums)

	var partition = func(nums []int, lo, hi int) int {
		pivot := nums[lo]
		i, j := lo+1, hi
		for i <= j {
			for i < hi && nums[i] <= pivot {
				i++
			}
			for j > lo && nums[j] > pivot {
				j--
			}

			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[lo], nums[j] = nums[j], nums[lo]
		return j
	}

	lo, hi := 0, n-1
	// 第k大 -> 第n-k小
	k = n - k
	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
}

// @lc code=end
