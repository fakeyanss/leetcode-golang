/*
 * @lc app=leetcode.cn id=2975 lang=golang
 * @lcpr version=30305
 *
 * [2975] 移除栅栏得到的正方形田地的最大面积
 *
 * https://leetcode.cn/problems/maximum-square-area-by-removing-fences-from-a-field/description/
 *
 * algorithms
 * Medium (30.71%)
 * Likes:    21
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 23.9K
 * Testcase Example:  '4\n3\n[2,3]\n[2]\n6\n7\n[2]\n[4]'
 *
 * 有一个大型的 (m - 1) x (n - 1) 矩形田地，其两个对角分别是 (1, 1) 和 (m, n)
 * ，田地内部有一些水平栅栏和垂直栅栏，分别由数组 hFences 和 vFences 给出。
 *
 * 水平栅栏为坐标 (hFences[i], 1) 到 (hFences[i], n)，垂直栅栏为坐标 (1, vFences[i]) 到 (m,
 * vFences[i]) 。
 *
 * 返回通过 移除 一些栅栏（可能不移除）所能形成的最大面积的 正方形 田地的面积，或者如果无法形成正方形田地则返回 -1。
 *
 * 由于答案可能很大，所以请返回结果对 10^9 + 7 取余 后的值。
 *
 * 注意：田地外围两个水平栅栏（坐标 (1, 1) 到 (1, n) 和坐标 (m, 1) 到 (m, n) ）以及两个垂直栅栏（坐标 (1, 1) 到
 * (m, 1) 和坐标 (1, n) 到 (m, n) ）所包围。这些栅栏 不能 被移除。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：m = 4, n = 3, hFences = [2,3], vFences = [2]
 * 输出：4
 * 解释：移除位于 2 的水平栅栏和位于 2 的垂直栅栏将得到一个面积为 4 的正方形田地。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：m = 6, n = 7, hFences = [2], vFences = [4]
 * 输出：-1
 * 解释：可以证明无法通过移除栅栏形成正方形田地。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= m, n <= 10^9
 * 1 <= hFences.length, vFences.length <= 600
 * 1 < hFences[i] < m
 * 1 < vFences[i] < n
 * hFences 和 vFences 中的元素是唯一的。
 *
 *
 */
package lc2975

import "slices"

// @lc code=start
// 思路：枚举
func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	hSet := f(hFences, m)
	vSet := f(vFences, n)

	res := 0
	for x := range hSet {
		if vSet[x] {
			res = max(res, x)
		}
	}

	if res == 0 {
		return -1
	}
	return res * res % (1e9 + 7)
}

// 枚举所有栅栏间的边长
func f(a []int, mx int) map[int]bool {
	a = append(a, 1, mx)
	slices.Sort(a)

	set := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			set[a[j]-a[i]] = true
		}
	}
	return set
}

// @lc code=end

/*
// @lcpr case=start
// 4\n3\n[2,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 6\n7\n[2]\n[4]\n
// @lcpr case=end

*/
