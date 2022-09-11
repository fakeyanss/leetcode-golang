/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 *
 * https://leetcode.cn/problems/range-sum-query-2d-immutable/description/
 *
 * algorithms
 * Medium (59.67%)
 * Likes:    419
 * Dislikes: 0
 * Total Accepted:    99.6K
 * Total Submissions: 166.9K
 * Testcase Example:  '["NumMatrix","sumRegion","sumRegion","sumRegion"]\n' +
  '[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]'
 *
 * 给定一个二维矩阵 matrix，以下类型的多个请求：
 *
 *
 * 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。
 *
 *
 * 实现 NumMatrix 类：
 *
 *
 * NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
 * int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1)
 * 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入:
 * ["NumMatrix","sumRegion","sumRegion","sumRegion"]
 *
 * [[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
 * 输出:
 * [null, 8, 11, 12]
 *
 * 解释:
 * NumMatrix numMatrix = new
 * NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
 * numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
 * numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
 * numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 200
 * -10^5 <= matrix[i][j] <= 10^5
 * 0 <= row1 <= row2 < m
 * 0 <= col1 <= col2 < n
 * 最多调用 10^4 次 sumRegion 方法
 *
 *
*/
package lc304

// @lc code=start
type NumMatrix struct {
    // preSum[i][j] 记录 matrix 中子矩阵 [0, 0, i-1, j-1] 的元素和
    preSum [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    preSum := make([][]int, m+1)
    for i := range preSum {
        preSum[i] = make([]int, n+1)
    }
    for i := range matrix {
        for j := range matrix[i] {
            preSum[i+1][j+1] = matrix[i][j] + preSum[i][j+1] + preSum[i+1][j] - preSum[i][j]
        }
    }
    return NumMatrix{preSum}
}

func (nu *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return nu.preSum[row2+1][col2+1] - nu.preSum[row2+1][col1] - nu.preSum[row1][col2+1] + nu.preSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
