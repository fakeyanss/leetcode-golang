/*
* @lc app=leetcode.cn id=221 lang=golang
* @lcpr version=30305
*
* [221] 最大正方形
*
* https://leetcode.cn/problems/maximal-square/description/
*
  - algorithms
  - Medium (51.78%)
  - Likes:    1862
  - Dislikes: 0
  - Total Accepted:    431.6K
  - Total Submissions: 833.3K
  - Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n' +
    '[["0","1"],["1","0"]]\n' +
    '[["0"]]'

*
* 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
*
*
*
* 示例 1：
*
* 输入：matrix =
* [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
* 输出：4
*
*
* 示例 2：
*
* 输入：matrix = [["0","1"],["1","0"]]
* 输出：1
*
*
* 示例 3：
*
* 输入：matrix = [["0"]]
* 输出：0
*
*
*
*
* 提示：
*
*
* m == matrix.length
* n == matrix[i].length
* 1 <= m, n <= 300
* matrix[i][j] 为 '0' 或 '1'
*
*
*/
package lc221

// @lc code=start
func maximalSquare(matrix [][]byte) int {
    // dp[i][j] 表示以matrix[i][j]为右下角的最大正方形边长
    // dp[i][j] = min(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])+1
    // 理解dp[i][j]为左、上、左上三个方向的最大正方形边长+1
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    res := 0
    for i := range m {
        for j := range n {
            if matrix[i][j] == '1' {
                // 兼容了dp[1][x]和dp[x][1]的情况
                dp[i+1][j+1] = min(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
                res = max(res, dp[i+1][j+1])
            }
        }
    }
    return res * res
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","1"],["1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

*/
