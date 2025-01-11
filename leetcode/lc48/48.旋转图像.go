/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=20004
 *
 * [48] 旋转图像
 *
 * https://leetcode.cn/problems/rotate-image/description/
 *
 * algorithms
 * Medium (77.35%)
 * Likes:    1970
 * Dislikes: 0
 * Total Accepted:    674.6K
 * Total Submissions: 870.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
 *
 * 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[7,4,1],[8,5,2],[9,6,3]]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
 * 输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length == matrix[i].length
 * 1 <= n <= 20
 * -1000 <= matrix[i][j] <= 1000
 *
 *
 *
 *
 */
package lc0048

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rotate(matrix [][]int) {
	// 方法1:原地旋转
	// 对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}

	// // 方法2:两次翻转
	// n := len(matrix)
	// // 先水平翻转
	// for i := 0; i < n/2; i++ {
	// 	matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	// }
	// // 再左上/右下对角线翻转
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < i; j++ {
	// 		matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
	// 	}
	// }
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/
