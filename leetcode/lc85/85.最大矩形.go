/*
* @lc app=leetcode.cn id=85 lang=golang
* @lcpr version=30305
*
* [85] 最大矩形
*
* https://leetcode.cn/problems/maximal-rectangle/description/
*
  - algorithms
  - Hard (56.16%)
  - Likes:    1799
  - Dislikes: 0
  - Total Accepted:    236K
  - Total Submissions: 417.4K
  - Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n' +
    '[["0"]]\n' +
    '[["1"]]'

*
* 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*
*
*
* 示例 1：
*
* 输入：matrix =
* [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
* 输出：6
* 解释：最大矩形如上图所示。
*
*
* 示例 2：
*
* 输入：matrix = [["0"]]
* 输出：0
*
*
* 示例 3：
*
* 输入：matrix = [["1"]]
* 输出：1
*
*
*
*
* 提示：
*
*
* rows == matrix.length
* cols == matrix[0].length
* 1 <= rows, cols <= 200
* matrix[i][j] 为 '0' 或 '1'
*
*
*/
package lc85

// 思路：单调栈
// @lc code=start
func maximalRectangle(matrix [][]byte) int {
    // 遍历每一行，将每一列累积的高度视作leetcode 84的柱状图
    heights := make([]int, len(matrix[0]))
    var res int
    for _, row := range matrix {
        for j, v := range row {
            if v == '1' {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }
        res = max(res, largestRectangleArea(heights))
    }
    return res
}

// leetcode 84 解法
func largestRectangleArea(heights []int) int {
    n := len(heights)
    left, right := make([]int, n), make([]int, n)
    var stk []int
    for i := 0; i < n; i++ {
        for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
            stk = stk[:len(stk)-1]
        }
        left[i] = -1
        if len(stk) > 0 {
            left[i] = stk[len(stk)-1]
        }
        stk = append(stk, i)
    }
    stk = []int{}
    for i := n - 1; i >= 0; i-- {
        for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
            stk = stk[:len(stk)-1]
        }
        right[i] = n
        if len(stk) > 0 {
            right[i] = stk[len(stk)-1]
        }
        stk = append(stk, i)
    }
    var res int
    for i, h := range heights {
        res = max(res, h*(right[i]-left[i]-1))
    }
    return res
}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1"]]\n
// @lcpr case=end

*/
