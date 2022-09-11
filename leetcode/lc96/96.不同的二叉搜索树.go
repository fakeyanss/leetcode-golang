/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode.cn/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (70.69%)
 * Likes:    1926
 * Dislikes: 0
 * Total Accepted:    292.4K
 * Total Submissions: 413.7K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：5
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package lc96

// @lc code=start
func numTrees(n int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var count func(lo, hi int) int
	count = func(lo, hi int) int {
		if lo > hi {
			return 1
		}
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}

		res := 0
		for mid := lo; mid <= hi; mid++ {
			left := count(lo, mid-1)
			right := count(mid+1, hi)
			res += left * right
		}

		memo[lo][hi] = res
		return res
	}

	return count(1, n)
}

// @lc code=end
