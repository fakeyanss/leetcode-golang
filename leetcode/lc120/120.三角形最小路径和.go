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
	// // dp[i][j] 从(0,0)出发到(i,j)的最小路径和
	// // 答案为求到达最后一行所有位置的dp[m][j]的最小值
	// m := len(triangle)
	// dp := make([][]int, m)
	// for i := range dp {
	// 	dp[i] = make([]int, i+1)
	// }
	// dp[0][0] = triangle[0][0]
	// for i := 1; i < m; i++ {
	// 	dp[i][0] = dp[i-1][0] + triangle[i][0]
	// 	dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	// 	for j := 1; j < i; j++ {
	// 		dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
	// 	}
	// }
	// res := math.MaxInt
	// for j := range dp[m-1] {
	// 	res = min(res, dp[m-1][j])
	// }
	// return res

	// dp[i][j]表示从(i,j)到最后一行的最小路径和
	m := len(triangle)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}
	dp[m-1] = triangle[m-1]
	for i := m - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
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
