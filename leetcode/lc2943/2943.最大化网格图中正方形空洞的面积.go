/*
 * @lc app=leetcode.cn id=2943 lang=golang
 * @lcpr version=30305
 *
 * [2943] 最大化网格图中正方形空洞的面积
 *
 * https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/description/
 *
 * algorithms
 * Medium (40.64%)
 * Likes:    29
 * Dislikes: 0
 * Total Accepted:    11.9K
 * Total Submissions: 21.7K
 * Testcase Example:  '2\n1\n[2,3]\n[2]\n1\n1\n[2]\n[2]\n2\n3\n[2,3]\n[2,4]'
 *
 * 给你两个整数 n 和 m，以及两个整数数组 hBars 和 vBars。网格由 n + 2 条水平线和 m + 2 条竖直线组成，形成 1x1
 * 的单元格。网格中的线条从 1 开始编号。
 *
 * 你可以从 hBars 中 删除 一些水平线条，并从 vBars 中删除一些竖直线条。注意，其他线条是固定的，无法删除。
 *
 * 返回一个整数表示移除一些线条（可以不移除任何线条）后，网格中 正方形空洞的最大面积 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: n = 2, m = 1, hBars = [2,3], vBars = [2]
 *
 * 输出: 4
 *
 * 解释:
 *
 * 左侧图片展示了网格的初始状态。水平线是 [1,2,3,4]，竖直线是 [1,2,3]。
 *
 * 构造最大正方形空洞的一种方法是移除水平线 2 和竖直线 2。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入: n = 1, m = 1, hBars = [2], vBars = [2]
 *
 * 输出: 4
 *
 * 解释:
 *
 * 移除水平线 2 和竖直线 2，可以得到最大正方形空洞。
 *
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入: n = 2, m = 3, hBars = [2,3], vBars = [2,4]
 *
 * 输出: 4
 *
 * 解释:
 *
 * 构造最大正方形空洞的一种方法是移除水平线 3 和竖直线 4。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^9
 * 1 <= m <= 10^9
 * 1 <= hBars.length <= 100
 * 2 <= hBars[i] <= n + 1
 * 1 <= vBars.length <= 100
 * 2 <= vBars[i] <= m + 1
 * hBars 中所有值互不相同。
 * vBars 中所有值互不相同。
 *
 *
 */
package lc2943

// @lc code=start
// 思路：贪心
// 思考hBars和vBars的线条先全部删掉，再找最大的正方形边长，将多余删掉的线条恢复。
// 换个思路就是找连续的线条形成的最大长度+1，即正方形边长。
func maximizeSquareHoleArea(_ int, _ int, hBars []int, vBars []int) int {
	side := min(longestConsecutive(hBars), longestConsecutive(vBars)) + 1
	return side * side
}

// leetcode 128，最长连续序列
func longestConsecutive(nums []int) int {
	var res int
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		numsMap[v] = true
	}
	for v := range numsMap {
		if numsMap[v-1] {
			continue
		}
		curNum, curLen := v, 1
		for numsMap[curNum+1] {
			curNum++
			curLen++
		}
		res = max(res, curLen)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 2\n1\n[2,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n[2]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 2\n3\n[2,3]\n[2,4]\n
// @lcpr case=end

*/
