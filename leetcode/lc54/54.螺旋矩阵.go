/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=20004
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (52.17%)
 * Likes:    1811
 * Dislikes: 0
 * Total Accepted:    637.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */
package lc0054

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, m-1, 0, n-1

	var res []int
	for top <= bottom && left <= right {
		if top <= bottom {
			for col := left; col <= right; col++ {
				res = append(res, matrix[top][col])
			}
			top++
		}

		if left <= right {
			for row := top; row <= bottom; row++ {
				res = append(res, matrix[row][right])
			}
			right--
		}

		if top <= bottom {
			for col := right; col >= left; col-- {
				res = append(res, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				res = append(res, matrix[row][left])
			}
			left++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
