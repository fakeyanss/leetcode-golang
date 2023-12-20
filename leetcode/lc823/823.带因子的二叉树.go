/*
 * @lc app=leetcode.cn id=823 lang=golang
 *
 * [823] 带因子的二叉树
 *
 * https://leetcode.cn/problems/binary-trees-with-factors/description/
 *
 * algorithms
 * Medium (43.58%)
 * Likes:    188
 * Dislikes: 0
 * Total Accepted:    19.5K
 * Total Submissions: 41.7K
 * Testcase Example:  '[2,4]'
 *
 * 给出一个含有不重复整数元素的数组 arr ，每个整数 arr[i] 均大于 1。
 *
 * 用这些整数来构建二叉树，每个整数可以使用任意次数。其中：每个非叶结点的值应等于它的两个子结点的值的乘积。
 *
 * 满足条件的二叉树一共有多少个？答案可能很大，返回 对 10^9 + 7 取余 的结果。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: arr = [2, 4]
 * 输出: 3
 * 解释: 可以得到这些二叉树: [2], [4], [4, 2, 2]
 *
 * 示例 2:
 *
 *
 * 输入: arr = [2, 4, 5, 10]
 * 输出: 7
 * 解释: 可以得到这些二叉树: [2], [4], [5], [10], [4, 2, 2], [10, 2, 5], [10, 5, 2].
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 1000
 * 2 <= arr[i] <= 10^9
 * arr 中的所有值 互不相同
 *
 *
 */
package lc823

import (
	"math"
	"sort"
)

// @lc code=start
func numFactoredBinaryTrees(arr []int) int {
	// key: root node value
	// value: number of binary tree
	dp := make(map[int]int)
	constant := (int)(math.Pow(10, 9) + 7)
	sort.Ints(arr)
	for i, curNum := range arr {
		// Case 1: cur_num as root with child nodes
		// scan each potential child node value
		for j := 0; j < i; j += 1 {
			factor := arr[j]
			quotient, remainder := curNum/factor, curNum%factor
			// current (factor, quotient) pair are feasible to be child nodes
			if remainder == 0 {
				dp[curNum] += dp[factor] * dp[quotient]
			}
		}

		// Case 2: cur_num as root without child nodes
		dp[curNum] += 1
	}

	totalCount := 0
	for _, count := range dp {
		totalCount += count
	}

	return totalCount % constant
}

// @lc code=end
