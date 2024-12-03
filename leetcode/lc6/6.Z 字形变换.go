/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=20003
 *
 * [6] Z 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (53.32%)
 * Likes:    2407
 * Dislikes: 0
 * Total Accepted:    738.9K
 * Total Submissions: 1.4M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1 <= numRows <= 1000
 *
 *
 */
package lc6

import "bytes"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r > n {
		return s
	}
	// 二维数组保存，r行c列
	t := r + r - 2 // 每个周期的字符个数
	// 每个周期要用r-1列，总列数为n/t*(r-1)或n/t*(r-1)+1
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]byte, c)
	}

	x, y := 0, 0
	for i := range s {
		mat[x][y] = s[i]
		if i%t < r-1 {
			x++ // 向下
		} else {
			y++
			x-- // 斜向上
		}
	}

	ans := []byte{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if mat[i][j] != 0 {
				ans = append(ans, mat[i][j])
			}
		}
	}

	return string(ans)
}

func convert_optimized(s string, numRows int) string {
	r := numRows
	if r == 1 || r > len(s) {
		return s
	}
	t := 2*r - 2
	mat := make([][]byte, r)
	x := 0
	for i := range s {
		mat[x] = append(mat[x], s[i])
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

// @lc code=end

/*
// @lcpr case=start
// "PAYPALISHIRING"\n3\n
// @lcpr case=end

// @lcpr case=start
// "PAYPALISHIRING"\n4\n
// @lcpr case=end

// @lcpr case=start
// "A"\n1\n
// @lcpr case=end

*/
