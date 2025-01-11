/*
 * @lc app=leetcode.cn id=1705 lang=golang
 * @lcpr version=20004
 *
 * [1705] 吃苹果的最大数目
 *
 * https://leetcode.cn/problems/maximum-number-of-eaten-apples/description/
 *
 * algorithms
 * Medium (45.50%)
 * Likes:    241
 * Dislikes: 0
 * Total Accepted:    41.2K
 * Total Submissions: 87.1K
 * Testcase Example:  '[1,2,3,5,2]\n[3,2,1,4,2]'
 *
 * 有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，这些苹果将会在 days[i]
 * 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且
 * days[i] == 0 表示。
 *
 * 你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。
 *
 * 给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：apples = [1,2,3,5,2], days = [3,2,1,4,2]
 * 输出：7
 * 解释：你可以吃掉 7 个苹果：
 * - 第一天，你吃掉第一天长出来的苹果。
 * - 第二天，你吃掉一个第二天长出来的苹果。
 * - 第三天，你吃掉一个第二天长出来的苹果。过了这一天，第三天长出来的苹果就已经腐烂了。
 * - 第四天到第七天，你吃的都是第四天长出来的苹果。
 *
 *
 * 示例 2：
 *
 * 输入：apples = [3,0,0,0,0,2], days = [3,0,0,0,0,2]
 * 输出：5
 * 解释：你可以吃掉 5 个苹果：
 * - 第一天到第三天，你吃的都是第一天长出来的苹果。
 * - 第四天和第五天不吃苹果。
 * - 第六天和第七天，你吃的都是第六天长出来的苹果。
 *
 *
 *
 *
 * 提示：
 *
 *
 * apples.length == n
 * days.length == n
 * 1 <= n <= 2 * 10^4
 * 0 <= apples[i], days[i] <= 2 * 10^4
 * 只有在 apples[i] = 0 时，days[i] = 0 才成立
 *
 *
 */
package lc1705

import (
	"container/heap"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func eatenApples(apples []int, days []int) int {
	h := &H{}
	res := 0

	for day := 0; day < len(apples) || h.Len() > 0; day++ {
		if day < len(apples) {
			if apples[day] > 0 {
				// 添加当天的苹果到堆中，存储的是过期日和数量
				heap.Push(h, Apple{expire: day + days[day], count: apples[day]})
			}
		}

		// 每天尝试吃苹果
		for h.Len() > 0 && (*h)[0].expire <= day {
			heap.Pop(h) // 移除过期的苹果
		}

		if h.Len() > 0 {
			// 吃掉一颗苹果
			res++
			(*h)[0].count--
			if (*h)[0].count == 0 {
				heap.Pop(h)
			}
		}
	}

	return res
}

type Apple struct {
	expire int // 过期日期
	count  int // 苹果数量
}

type H []Apple

func (h H) Len() int           { return len(h) }
func (h H) Less(i, j int) bool { return h[i].expire < h[j].expire }
func (h H) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *H) Push(x any) {
	*h = append(*h, x.(Apple))
}
func (h *H) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,5,2]\n[3,2,1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,0,0,0,0,2]\n[3,0,0,0,0,2]\n
// @lcpr case=end

*/
