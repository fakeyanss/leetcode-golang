/*
 * @lc app=leetcode.cn id=1802 lang=golang
 *
 * [1802] 有界数组中指定下标处的最大值
 *
 * https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/
 *
 * algorithms
 * Medium (29.27%)
 * Likes:    144
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 52K
 * Testcase Example:  '4\n2\n6'
 *
 * 给你三个正整数 n、index 和 maxSum 。你需要构造一个同时满足下述所有条件的数组 nums（下标 从 0 开始 计数）：
 *
 *
 * nums.length == n
 * nums[i] 是 正整数 ，其中 0 <= i < n
 * abs(nums[i] - nums[i+1]) <= 1 ，其中 0 <= i < n-1
 * nums 中所有元素之和不超过 maxSum
 * nums[index] 的值被 最大化
 *
 *
 * 返回你所构造的数组中的 nums[index] 。
 *
 * 注意：abs(x) 等于 x 的前提是 x >= 0 ；否则，abs(x) 等于 -x 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 4, index = 2,  maxSum = 6
 * 输出：2
 * 解释：数组 [1,1,2,1] 和 [1,2,2,1] 满足所有条件。不存在其他在指定下标处具有更大值的有效数组。
 *
 *
 * 示例 2：
 *
 * 输入：n = 6, index = 1,  maxSum = 10
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= maxSum <= 10^9
 * 0 <= index < n
 *
 *
 */
package lc1802

// @lc code=start
// 二分查找
func maxValue(n int, index int, maxSum int) int {
	var res int
	l, r := 1, maxSum
	for l <= r {
		mid := l + (r-l)/2
		if isPossible(n, index, maxSum, mid) {
			res = max(res, mid)
			l = mid + 1 // let's maximize
		} else {
			r = mid - 1
		}
	}
	return res
}

func isPossible(n, index, maxSum, target int) bool {
	sum := target
	// For left and right we need to calculate the sum. We can't iterate over
	// the array and keep placing numbers. We havee to be faster than that.
	// For that, we'll compute the number of positions and the max value
	sum += compute(index, target-1) // left
	if sum > maxSum {
		return false
	}
	sum += compute(n-index-1, target-1)
	if sum > maxSum {
		return false
	}
	return true
}

// compute the sum for the given number of positions with the given maximum value
func compute(positions, maxVal int) int {
	if positions >= maxVal {
		return arithSum(maxVal) + (positions - maxVal) // filling in remaing by 1s
	}
	return arithSum(maxVal) - arithSum(maxVal-positions)
}

// arithSum computes the arithmetic progression from 1..n
func arithSum(n int) int {
	return n * (n + 1) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
