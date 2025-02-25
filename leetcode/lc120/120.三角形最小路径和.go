/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=20004
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode.cn/problems/triangle/description/
 *
 * algorithms
 * Medium (69.13%)
 * Likes:    1409
 * Dislikes: 0
 * Total Accepted:    395.7K
 * Total Submissions: 572K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 *
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 * 示例 2：
 *
 * 输入：triangle = [[-10]]
 * 输出：-10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= triangle.length <= 200
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 <= triangle[i][j] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 *
 *
 */
package lc120

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumTotal(triangle [][]int) int {
	// n := len(triangle)
	// var dp func(i, j int) int // dp(i,j)表示triangle[i][j]出发到最后一行的的最小路径和
	// memo := make([][]int, n)
	// for i := range memo {
	// 	memo[i] = make([]int, i+1)
	// 	for j := range memo[i] {
	// 		memo[i][j] = 10001
	// 	}
	// }
	// dp = func(i, j int) int {
	// 	if i == n-1 {
	// 		return triangle[i][j]
	// 	}
	// 	p := &memo[i][j]
	// 	if *p != 10001 {
	// 		return *p
	// 	}
	// 	*p = min(dp(i+1, j), dp(i+1, j+1)) + triangle[i][j]
	// 	return *p
	// }
	// return dp(0, 0)

	n := len(triangle)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[n-1] = triangle[n-1]
	for i := n - 2; i >= 0; i-- {
		for j, x := range triangle[i] {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + x
		}
	}
	return dp[0][0]
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/
