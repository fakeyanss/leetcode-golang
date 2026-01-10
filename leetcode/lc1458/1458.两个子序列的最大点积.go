/*
 * @lc app=leetcode.cn id=1458 lang=golang
 * @lcpr version=30305
 *
 * [1458] 两个子序列的最大点积
 *
 * https://leetcode.cn/problems/max-dot-product-of-two-subsequences/description/
 *
 * algorithms
 * Hard (54.58%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    18.9K
 * Total Submissions: 33.4K
 * Testcase Example:  '[2,1,-2,5]\n[3,0,-6]\n[3,-2]\n[2,-6,7]\n[-1,-1]\n[1,1]'
 *
 * 给你两个数组 nums1 和 nums2 。
 *
 * 请你返回 nums1 和 nums2 中两个长度相同的 非空 子序列的最大点积。
 *
 * 数组的非空子序列是通过删除原数组中某些元素（可能一个也不删除）后剩余数字组成的序列，但不能改变数字间相对顺序。比方说，[2,3,5] 是
 * [1,2,3,4,5] 的一个子序列而 [1,5,3] 不是。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [2,1,-2,5], nums2 = [3,0,-6]
 * 输出：18
 * 解释：从 nums1 中得到子序列 [2,-2] ，从 nums2 中得到子序列 [3,-6] 。
 * 它们的点积为 (2*3 + (-2)*(-6)) = 18 。
 *
 * 示例 2：
 *
 * 输入：nums1 = [3,-2], nums2 = [2,-6,7]
 * 输出：21
 * 解释：从 nums1 中得到子序列 [3] ，从 nums2 中得到子序列 [7] 。
 * 它们的点积为 (3*7) = 21 。
 *
 * 示例 3：
 *
 * 输入：nums1 = [-1,-1], nums2 = [1,1]
 * 输出：-1
 * 解释：从 nums1 中得到子序列 [-1] ，从 nums2 中得到子序列 [1] 。
 * 它们的点积为 -1 。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 500
 * -1000 <= nums1[i], nums2[i] <= 1000
 *
 *
 *
 *
 * 点积：
 *
 * 定义 a = [a1, a2,…, an] 和 b = [b1, b2,…, bn] 的点积为：
 *
 *
 *
 * 这里的 Σ 指示总和符号。
 *
 *
 */
package lc1458

import "math"

// @lc code=start
// 思路：DP
func maxDotProduct(nums1 []int, nums2 []int) int {
	// dp[i][j]表示nums[:i+1]和nums2[:j+1]的子序列最大点积
	// 需要选择的两个子序列长度相等，所以对每个元素需要枚举需要选或不选
	// 状态1: 选nums1[i]和nums2[j]，前面的元素可以选或不选, dp[i][j]=nums1[i]*num2[j]+dp[i-1][j-1]，dp[i-1][j-1]
	// 状态2: 只选nums1[i], dp[i][j]=dp[i][j-1]
	// 状态3: 只选nums2[j], dp[i][j]=dp[i-1][j]
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = math.MinInt
	}
	for j := range dp[0] {
		dp[0][j] = math.MinInt
	}

	for i, x := range nums1 {
		for j, y := range nums2 {
			// +1兼容i,j为0的情况
			dp[i+1][j+1] = max(max(dp[i][j], 0)+x*y, dp[i+1][j], dp[i][j+1])
		}
	}
	return dp[m][n]
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,-2,5]\n[3,0,-6]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2]\n[2,-6,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1,-1]\n[1,1]\n
// @lcpr case=end

*/
