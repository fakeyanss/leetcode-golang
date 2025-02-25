/*
 * @lc app=leetcode.cn id=149 lang=golang
 * @lcpr version=20004
 *
 * [149] 直线上最多的点数
 *
 * https://leetcode.cn/problems/max-points-on-a-line/description/
 *
 * algorithms
 * Hard (41.43%)
 * Likes:    581
 * Dislikes: 0
 * Total Accepted:    104.7K
 * Total Submissions: 252.2K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
 *
 *
 *
 * 示例 1：
 *
 * 输入：points = [[1,1],[2,2],[3,3]]
 * 输出：3
 *
 *
 * 示例 2：
 *
 * 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= points.length <= 300
 * points[i].length == 2
 * -10^4 <= xi, yi <= 10^4
 * points 中的所有点 互不相同
 *
 *
 */
package lc149

import (
	"fmt"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxPoints(points [][]int) int {
	n := len(points)
	res := 1
	for i := 0; i < n; i++ {
		mapSlope := make(map[string]int)
		// 由当前点 i 发出的直线所经过的最多点数量
		imax := 0
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			a, b := x1-x2, y1-y2
			g := gcd(a, b)
			key := fmt.Sprintf("%d_%d", a/g, b/g)
			mapSlope[key]++
			imax = max(imax, mapSlope[key])
		}
		res = max(res, imax+1)
	}
	return res
}

// 最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1],[2,2],[3,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]\n
// @lcpr case=end

*/
