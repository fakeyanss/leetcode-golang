/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (52.04%)
 * Likes:    2085
 * Dislikes: 0
 * Total Accepted:    573.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
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
 * 1
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1
 *
 *
 */
package lc6

import "bytes"

// @lc code=start

// 结果用二维矩阵存储，r行c列
// 考虑Z字形排列，拆解为周期性的动作，每个周期的字符个数为 t=r(竖向)+r-2(斜向)=2*r-2,
// 每个周期占 1+r-2=r-1 列，总列数 c=(n + t -1)/t * (r-1)
// 用 (x,y) 记录当前的矩阵下标，起始为 (0,0)，先纵向移动，再斜向上移动。即：
// 当 i%t<r-1 时，x++; 否则 x--, y++
// 再考虑矩阵中存在大量的空洞，可以不必提前分配每行长度，再使用时再添加。
func convert(s string, numRows int) string {
	return convert2(s, numRows)
}

func convert1(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r > n {
		return s
	}
	t := 2*r - 2
	c := (n + t - 1) / t * (r - 1)
	mat := make([][]byte, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i := range s {
		mat[x][y] = s[i]
		if i%t < r-1 {
			x++
		} else {
			x--
			y++
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

func convert2(s string, numRows int) string {
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
