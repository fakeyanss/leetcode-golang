/*
* @lc app=leetcode.cn id=1895 lang=golang
* @lcpr version=30305
*
* [1895] 最大的幻方
*
* https://leetcode.cn/problems/largest-magic-square/description/
*
  - algorithms
  - Medium (59.84%)
  - Likes:    34
  - Dislikes: 0
  - Total Accepted:    8K
  - Total Submissions: 12.1K
  - Testcase Example:  '[[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]\n' +
    '[[5,1,3,1],[9,3,3,1],[1,3,3,8]]'

*
* 一个 k x k 的 幻方 指的是一个 k x k 填满整数的方格阵，且每一行、每一列以及两条对角线的和 全部相等 。幻方中的整数 不需要互不相同
* 。显然，每个 1 x 1 的方格都是一个幻方。
*
* 给你一个 m x n 的整数矩阵 grid ，请你返回矩阵中 最大幻方 的 尺寸 （即边长 k）。
*
*
*
* 示例 1：
*
* 输入：grid = [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]
* 输出：3
* 解释：最大幻方尺寸为 3 。
* 每一行，每一列以及两条对角线的和都等于 12 。
* - 每一行的和：5+1+6 = 5+4+3 = 2+7+3 = 12
* - 每一列的和：5+5+2 = 1+4+7 = 6+3+3 = 12
* - 对角线的和：5+4+3 = 6+4+2 = 12
*
*
* 示例 2：
*
* 输入：grid = [[5,1,3,1],[9,3,3,1],[1,3,3,8]]
* 输出：2
*
*
*
*
* 提示：
*
*
* m == grid.length
* n == grid[i].length
* 1 <= m, n <= 50
* 1 <= grid[i][j] <= 10^6
*
*
*/
package lc1895

// @lc code=start
// 思路：枚举，前缀和
func largestMagicSquare(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    rowSum := make([][]int, m) // 每一行的前缀和
    for i := range m {
        rowSum[i] = make([]int, n)
        rowSum[i][0] = grid[i][0]
        for j := 1; j < n; j++ {
            rowSum[i][j] = rowSum[i][j-1] + grid[i][j]
        }
    }
    colSum := make([][]int, m) // 每一列的前缀和
    for i := range m {
        colSum[i] = make([]int, n)
    }
    for j := range n {
        colSum[0][j] = grid[0][j]
        for i := 1; i < m; i++ {
            colSum[i][j] = colSum[i-1][j] + grid[i][j]
        }
    }
    // 枚举所有正方形边长
    for edge := min(m, n); edge >= 2; edge-- {
        // 枚举正方形左上角(i,j)
        for i := 0; i+edge <= m; i++ {
            for j := 0; j+edge <= n; j++ {
                baseSum := rowSum[i][j+edge-1] // 正方形第一行的和
                if j > 0 {
                    baseSum -= rowSum[i][j-1]
                }

                check := true
                // 检查每一行
                for ii := i + 1; ii < i+edge; ii++ {
                    sum := rowSum[ii][j+edge-1]
                    if j > 0 {
                        sum -= rowSum[ii][j-1]
                    }
                    if sum != baseSum {
                        check = false
                        break
                    }
                }
                if !check {
                    continue
                }
                // 检查每一列
                for jj := j; jj < j+edge; jj++ {
                    sum := colSum[i+edge-1][jj]
                    if i > 0 {
                        sum -= colSum[i-1][jj]
                    }
                    if sum != baseSum {
                        check = false
                        break
                    }
                }
                if !check {
                    continue
                }
                // 检查对角线
                d1, d2 := 0, 0
                for k := 0; k < edge; k++ {
                    d1 += grid[i+k][j+k]
                    d2 += grid[i+k][j+edge-1-k]
                }
                if d1 == baseSum && d2 == baseSum {
                    return edge // 因为edge是从大到小遍历，第一个满足条件的幻方就是结果
                }
            }
        }
    }
    return 1
}

// @lc code=end

/*
// @lcpr case=start
// [[7,1,4,5,6],[2,5,1,6,4],[1,5,4,3,2],[1,2,7,3,4]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,3,1],[9,3,3,1],[1,3,3,8]]\n
// @lcpr case=end

*/
