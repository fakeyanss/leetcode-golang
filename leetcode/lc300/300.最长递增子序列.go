/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (54.07%)
 * Likes:    2777
 * Dislikes: 0
 * Total Accepted:    610.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
 *
 * 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7]
 * 的子序列。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [10,9,2,5,3,7,101,18]
 * 输出：4
 * 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,3,2,3]
 * 输出：4
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [7,7,7,7,7,7,7]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2500
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
 *
 *
 */
package lc300

// @lc code=start
func lengthOfLIS(nums []int) int {
	// return dp(nums)
	return binarySearch(nums)
}

func dp(nums []int) int {
	// dp[i]表示以第i个元素结尾的最长子序列长度
	// dp[i] = max(dp[i], dp[j]+1), j < i, nums[j] < nums[i]
	n := len(nums)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}

	// 选出最大的dp[i]
	res := 0
	for _, v := range dp {
		res = maxInt(res, v)
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func binarySearch(nums []int) int {
	maxLen, n := 1, len(nums)
	if n == 0 {
		return 0
	}
	// tail[i] 表示长度为i的最长上升子序列的末尾元素的最小值
	tail := make([]int, n+1)
	tail[1] = nums[0]
	// 设当前已求出的最长上升子序列的长度为 len（初始时为 11），从前往后遍历数组 nums，并始终维护单调递增。
	// 在遍历到 nums[i] 时：
	// 如果 nums[i]>d[len] ，则直接加入到 tail 数组末尾，并更新 len=len+1；
	// 否则，在 tail 数组中二分查找，找到第一个比 nums[i] 小的数 tail[k] ，并更新 d[k+1]=nums[i]。
	for i := 1; i < n; i++ {
		if nums[i] > tail[maxLen] {
			maxLen++
			tail[maxLen] = nums[i]
		} else {
			l, r := 1, maxLen-1
			pos := 0 // 找不到时，直接更新tail[1]
			for l <= r {
				mid := l + (r-l)/2
				if tail[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			tail[pos+1] = nums[i]
		}
	}
	return maxLen
}

// @lc code=end
