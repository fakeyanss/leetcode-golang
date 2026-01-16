/*
 * @lc app=leetcode.cn id=3453 lang=golang
 * @lcpr version=30305
 *
 * [3453] 分割正方形 I
 *
 * https://leetcode.cn/problems/separate-squares-i/description/
 *
 * algorithms
 * Medium (44.91%)
 * Likes:    29
 * Dislikes: 0
 * Total Accepted:    10.5K
 * Total Submissions: 19.7K
 * Testcase Example:  '[[0,0,1],[2,2,1]]\n[[0,0,2],[1,1,1]]'
 *
 * 给你一个二维整数数组 squares ，其中 squares[i] = [xi, yi, li] 表示一个与 x
 * 轴平行的正方形的左下角坐标和正方形的边长。
 *
 * 找到一个最小的 y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 等于 该线以下正方形的总面积。
 *
 * 答案如果与实际答案的误差在 10^-5 以内，将视为正确答案。
 *
 * 注意：正方形 可能会 重叠。重叠区域应该被 多次计数 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： squares = [[0,0,1],[2,2,1]]
 *
 * 输出： 1.00000
 *
 * 解释：
 *
 *
 *
 * 任何在 y = 1 和 y = 2 之间的水平线都会有 1 平方单位的面积在其上方，1 平方单位的面积在其下方。最小的 y 坐标是 1。
 *
 *
 * 示例 2：
 *
 *
 * 输入： squares = [[0,0,2],[1,1,1]]
 *
 * 输出： 1.16667
 *
 * 解释：
 *
 *
 *
 * 面积如下：
 *
 *
 * 线下的面积：7/6 * 2 (红色) + 1/6 (蓝色) = 15/6 = 2.5。
 * 线上的面积：5/6 * 2 (红色) + 5/6 (蓝色) = 15/6 = 2.5。
 *
 *
 * 由于线以上和线以下的面积相等，输出为 7/6 = 1.16667。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= squares.length <= 5 * 10^4
 * squares[i] = [xi, yi, li]
 * squares[i].length == 3
 * 0 <= xi, yi <= 10^9
 * 1 <= li <= 10^9
 * 所有正方形的总面积不超过 10^12。
 *
 *
 */
package lc3453

import "math/bits"

// @lc code=start
// 思路：二分
func separateSquares(squares [][]int) float64 {
	totalArea, maxY := 0, 0
	for _, sq := range squares {
		l := sq[2]
		totalArea += l * l
		maxY = max(maxY, sq[1]+l)
	}

	check := func(y float64) bool {
		area := 0.
		for _, sq := range squares {
			yi := float64(sq[1])
			if yi < y { // 计算正放在在y之下的面积
				l := float64(sq[2])
				area += l * min(l, y-yi)
			}
		}
		return area >= float64(totalArea)/2
	}

	left, right := 0., float64(maxY)
	// 循环条件不用left和right比较，因为float有误差，直接计算最大循环次数。
	// 题目误差要求为10^-5，所以最大二分次数为log(maxY/10^-5)，
	// bits.Len()计算二进制位数，与求log2结果向上取整相同
	for range bits.Len(uint(maxY * 1e5)) {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return (left + right) / 2 // 取中点误差更小
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,1],[2,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,2],[1,1,1]]\n
// @lcpr case=end

*/
