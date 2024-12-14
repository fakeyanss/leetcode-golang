/*
 * @lc app=leetcode.cn id=3264 lang=golang
 * @lcpr version=20004
 *
 * [3264] K 次乘运算后的最终数组 I
 *
 * https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-i/description/
 *
 * algorithms
 * Easy (78.21%)
 * Likes:    18
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 14.7K
 * Testcase Example:  '[2,1,3,5,6]\n5\n2'
 *
 * 给你一个整数数组 nums ，一个整数 k  和一个整数 multiplier 。
 *
 * 你需要对 nums 执行 k 次操作，每次操作中：
 *
 *
 * 找到 nums 中的 最小 值 x ，如果存在多个最小值，选择最 前面 的一个。
 * 将 x 替换为 x * multiplier 。
 *
 *
 * 请你返回执行完 k 次乘运算之后，最终的 nums 数组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,1,3,5,6], k = 5, multiplier = 2
 *
 * 输出：[8,4,6,5,6]
 *
 * 解释：
 *
 *
 *
 *
 * 操作
 * 结果
 *
 *
 * 1 次操作后
 * [2, 2, 3, 5, 6]
 *
 *
 * 2 次操作后
 * [4, 2, 3, 5, 6]
 *
 *
 * 3 次操作后
 * [4, 4, 3, 5, 6]
 *
 *
 * 4 次操作后
 * [4, 4, 6, 5, 6]
 *
 *
 * 5 次操作后
 * [8, 4, 6, 5, 6]
 *
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2], k = 3, multiplier = 4
 *
 * 输出：[16,8]
 *
 * 解释：
 *
 *
 *
 *
 * 操作
 * 结果
 *
 *
 * 1 次操作后
 * [4, 2]
 *
 *
 * 2 次操作后
 * [4, 8]
 *
 *
 * 3 次操作后
 * [16, 8]
 *
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 * 1 <= k <= 10
 * 1 <= multiplier <= 5
 *
 *
 */
package lc3264

import "container/heap"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getFinalState(nums []int, k int, multiplier int) []int {
	res := make([]int, len(nums))
	h := &itemHeap{}
	heap.Init(h)
	for i, val := range nums {
		heap.Push(h, item{val, i})
	}
	for i := 0; i < k; i++ {
		it := heap.Pop(h).(item)
		it.val = it.val * multiplier
		heap.Push(h, it)
	}
	for h.Len() > 0 {
		x := heap.Pop(h).(item)
		res[x.idx] = x.val
	}
	return res
}

type item struct {
	val int
	idx int
}

type itemHeap []item

func (h itemHeap) Len() int {
	return len(h)
}

func (h itemHeap) Less(i, j int) bool {
	if h[i].val < h[j].val {
		return true
	} else if h[i].val == h[j].val {
		return h[i].idx < h[j].idx
	} else {
		return false
	}
}

func (h itemHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *itemHeap) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *itemHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3,5,6]\n5\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n3\n4\n
// @lcpr case=end

*/
