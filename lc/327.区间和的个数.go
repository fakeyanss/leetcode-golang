/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 *
 * https://leetcode.cn/problems/count-of-range-sum/description/
 *
 * algorithms
 * Hard (41.26%)
 * Likes:    476
 * Dislikes: 0
 * Total Accepted:    33.8K
 * Total Submissions: 82K
 * Testcase Example:  '[-2,5,-1]\n-2\n2'
 *
 * 给你一个整数数组 nums 以及两个整数 lower 和 upper 。求数组中，值位于范围 [lower, upper] （包含 lower 和
 * upper）之内的 区间和的个数 。
 *
 * 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-2,5,-1], lower = -2, upper = 2
 * 输出：3
 * 解释：存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0], lower = 0, upper = 0
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 * -10^5
 * 题目数据保证答案是一个 32 位 的整数
 *
 *
 */
package lc

// @lc code=start
var lower327, upper327 int
var count327 int
var temp327 []int

func countRangeSum(nums []int, lower int, upper int) int {
	lower327 = lower
	upper327 = upper
	count327 = 0
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = nums[i] + preSum[i]
	}
	// 初始化辅助数组
	temp327 = make([]int, n+1)
	sort327(preSum, 0, n)
	return count327
}

func sort327(preSum []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort327(preSum, lo, mid)
	sort327(preSum, mid+1, hi)
	merge327(preSum, lo, mid, hi)
}

func merge327(preSum []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp327[i] = preSum[i]
	}

	// 维护左闭右开区间 [start, end) 中的元素落在 [lower, upper] 中
	start, end := mid+1, mid+1
	for i := lo; i <= mid; i++ {
		for start <= hi && preSum[start]-preSum[i] < lower327 {
			start++
		}
		for end <= hi && preSum[end]-preSum[i] <= upper327 {
			end++
		}
		count327 = count327 + end - start
	}

	// 合并
	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			preSum[p] = temp327[j]
			j++
		} else if j == hi+1 {
			preSum[p] = temp327[i]
			i++
		} else if temp327[i] > temp327[j] {
			preSum[p] = temp327[j]
			j++
		} else {
			preSum[p] = temp327[i]
			i++
		}
	}
}

// @lc code=end
