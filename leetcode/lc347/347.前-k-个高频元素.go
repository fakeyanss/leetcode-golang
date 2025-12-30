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

// @lc code=start
// 堆排序
// type Entry struct {
// 	val  int
// 	freq int
// }

// type EntryHeap []Entry

// func (h EntryHeap) Len() int           { return len(h) }
// func (h EntryHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
// func (h EntryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *EntryHeap) Push(x any)        { *h = append(*h, x.(Entry)) }
// func (h *EntryHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

// func topKFrequent(nums []int, k int) []int {
// 	cnt := make(map[int]int)
// 	for _, val := range nums {
// 		cnt[val]++
// 	}

// 	h := &EntryHeap{}
// 	heap.Init(h)
// 	for key, val := range cnt {
// 		heap.Push(h, Entry{val: key, freq: val})
// 		if h.Len() > k {
// 			heap.Pop(h) // 只需要前k大
// 		}
// 	}

// 	res := make([]int, k)
// 	for i := range k {
// 		res[k-i-1] = heap.Pop(h).(Entry).val
// 	}
// 	return res
// }

// 桶排序
func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	maxCnt := 0
	for _, val := range nums {
		cnt[val]++
		maxCnt = max(maxCnt, cnt[val])
	}

	bucket := make([][]int, maxCnt+1)
	for x, c := range cnt {
		bucket[c] = append(bucket[c], x)
	}

	// 倒序遍历 buckets，把出现次数前 k 大的元素加入答案
	res := make([]int, 0, k)
	for i := maxCnt; i >= 0 && len(res) < k; i-- {
		res = append(res, bucket[i]...)
	}
	return res[:k]
}

// @lc code=end
