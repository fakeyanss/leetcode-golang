/*
 * @lc app=leetcode.cn id=2931 lang=golang
 * @lcpr version=20004
 *
 * [2931] 购买物品的最大开销
 *
 * https://leetcode.cn/problems/maximum-spending-after-buying-items/description/
 *
 * algorithms
 * Hard (67.32%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    8.8K
 * Total Submissions: 12.2K
 * Testcase Example:  '[[8,5,2],[6,4,1],[9,7,3]]'
 *
 * 给你一个下标从 0 开始大小为 m * n 的整数矩阵 values ，表示 m 个不同商店里 m * n 件不同的物品。每个商店有 n 件物品，第 i
 * 个商店的第 j 件物品的价值为 values[i][j] 。除此以外，第 i 个商店的物品已经按照价值非递增排好序了，也就是说对于所有 0 <= j <
 * n - 1 都有 values[i][j] >= values[i][j + 1] 。
 *
 * 每一天，你可以在一个商店里购买一件物品。具体来说，在第 d 天，你可以：
 *
 *
 * 选择商店 i 。
 * 购买数组中最右边的物品 j ，开销为 values[i][j] * d 。换句话说，选择该商店中还没购买过的物品中最大的下标 j ，并且花费
 * values[i][j] * d 去购买。
 *
 *
 * 注意，所有物品都视为不同的物品。比方说如果你已经从商店 1 购买了物品 0 ，你还可以在别的商店里购买其他商店的物品 0 。
 *
 * 请你返回购买所有 m * n 件物品需要的 最大开销 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：values = [[8,5,2],[6,4,1],[9,7,3]]
 * 输出：285
 * 解释：第一天，从商店 1 购买物品 2 ，开销为 values[1][2] * 1 = 1 。
 * 第二天，从商店 0 购买物品 2 ，开销为 values[0][2] * 2 = 4 。
 * 第三天，从商店 2 购买物品 2 ，开销为 values[2][2] * 3 = 9 。
 * 第四天，从商店 1 购买物品 1 ，开销为 values[1][1] * 4 = 16 。
 * 第五天，从商店 0 购买物品 1 ，开销为 values[0][1] * 5 = 25 。
 * 第六天，从商店 1 购买物品 0 ，开销为 values[1][0] * 6 = 36 。
 * 第七天，从商店 2 购买物品 1 ，开销为 values[2][1] * 7 = 49 。
 * 第八天，从商店 0 购买物品 0 ，开销为 values[0][0] * 8 = 64 。
 * 第九天，从商店 2 购买物品 0 ，开销为 values[2][0] * 9 = 81 。
 * 所以总开销为 285 。
 * 285 是购买所有 m * n 件物品的最大总开销。
 *
 *
 * 示例 2：
 *
 * 输入：values = [[10,8,6,4,2],[9,7,5,3,2]]
 * 输出：386
 * 解释：第一天，从商店 0 购买物品 4 ，开销为 values[0][4] * 1 = 2 。
 * 第二天，从商店 1 购买物品 4 ，开销为 values[1][4] * 2 = 4 。
 * 第三天，从商店 1 购买物品 3 ，开销为 values[1][3] * 3 = 9 。
 * 第四天，从商店 0 购买物品 3 ，开销为 values[0][3] * 4 = 16 。
 * 第五天，从商店 1 购买物品 2 ，开销为 values[1][2] * 5 = 25 。
 * 第六天，从商店 0 购买物品 2 ，开销为 values[0][2] * 6 = 36 。
 * 第七天，从商店 1 购买物品 1 ，开销为 values[1][1] * 7 = 49 。
 * 第八天，从商店 0 购买物品 1 ，开销为 values[0][1] * 8 = 64 。
 * 第九天，从商店 1 购买物品 0 ，开销为 values[1][0] * 9 = 81 。
 * 第十天，从商店 0 购买物品 0 ，开销为 values[0][0] * 10 = 100 。
 * 所以总开销为 386 。
 * 386 是购买所有 m * n 件物品的最大总开销。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m == values.length <= 10
 * 1 <= n == values[i].length <= 10^4
 * 1 <= values[i][j] <= 10^6
 * values[i] 按照非递增顺序排序。
 *
 *
 */
package lc2931

import (
	"container/heap"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSpending(values [][]int) int64 {
	var cost int64

	m := len(values)
	itemHeap := &ItemHeap{}
	heap.Init(itemHeap)
	for i := 0; i < m; i++ {
		heap.Push(itemHeap, Item{values[i][len(values[i])-1], i})
		values[i] = values[i][:len(values[i])-1]
	}

	d := 1
	for itemHeap.Len() > 0 {
		item := heap.Pop(itemHeap).(Item)
		cost += int64(item.value * d)
		d++
		row := item.idx
		if col := len(values[row]); col > 0 {
			heap.Push(itemHeap, Item{values[row][col-1], row})
			values[row] = values[row][:col-1]
		}
	}

	return cost
}

// ItemHeap 定义一个切片类型，实现 heap.Interface 接口
type ItemHeap []Item
type Item struct {
	value int
	idx   int
}

// Len 返回堆的长度
func (h ItemHeap) Len() int { return len(h) }

// Less 定义小顶堆，值小的优先级高
func (h ItemHeap) Less(i, j int) bool { return h[i].value < h[j].value }

// Swap 交换堆中两个元素的位置
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 将元素插入堆（接口要求使用指针接收器）
func (h *ItemHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

// Pop 移除并返回堆顶元素（接口要求使用指针接收器）
func (h *ItemHeap) Pop() any {
	// 从堆中移除最小元素
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// @lc code=end

/*
// @lcpr case=start
// [[8,5,2],[6,4,1],[9,7,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[10,8,6,4,2],[9,7,5,3,2]]\n
// @lcpr case=end

*/
