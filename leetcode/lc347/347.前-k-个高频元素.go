/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode.cn/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.33%)
 * Likes:    1397
 * Dislikes: 0
 * Total Accepted:    372.1K
 * Total Submissions: 587.5K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * k 的取值范围是 [1, 数组中不相同的元素的个数]
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
 *
 *
 *
 *
 * 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
 *
 */
package lc347

import "container/heap"

// @lc code=start

// 二叉堆基本应用
type Entry struct {
	val  int
	freq int
}

type EntryHeap []Entry

func (h EntryHeap) Len() int {
	return len(h)
}

func (h EntryHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h EntryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EntryHeap) Push(x interface{}) {
	*h = append(*h, x.(Entry))
}

func (h *EntryHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// 题目数据保证前k个高频元素的频次不会相同
func topKFrequent(nums []int, k int) []int {
	occurrences := make(map[int]int)
	for _, val := range nums {
		occurrences[val]++
	}
	h := &EntryHeap{}
	heap.Init(h)

	for key, val := range occurrences {
		heap.Push(h, Entry{val: key, freq: val})
		if h.Len() > k {
			// 只需要前k大
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).(Entry).val
	}
	return res
}

// @lc code=end
