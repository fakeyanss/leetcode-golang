/*
* @lc app=leetcode.cn id=295 lang=golang
*
* [295] 数据流的中位数
*
* https://leetcode.cn/problems/find-median-from-data-stream/description/
*
  - algorithms
  - Hard (53.11%)
  - Likes:    767
  - Dislikes: 0
  - Total Accepted:    96.7K
  - Total Submissions: 182K
  - Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
    '[[],[1],[2],[],[3],[]]'

*
* 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
*
* 例如，
*
* [2,3,4] 的中位数是 3
*
* [2,3] 的中位数是 (2 + 3) / 2 = 2.5
*
* 设计一个支持以下两种操作的数据结构：
*
*
* void addNum(int num) - 从数据流中添加一个整数到数据结构中。
* double findMedian() - 返回目前所有元素的中位数。
*
*
* 示例：
*
* addNum(1)
* addNum(2)
* findMedian() -> 1.5
* addNum(3)
* findMedian() -> 2
*
* 进阶:
*
*
* 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
* 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
*
*
*/
package lc295

import (
    "container/heap"
)

// @lc code=start
type MedianFinder struct {
    minHeap hp
    maxHeap hp
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(int)) }
func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

func Constructor() MedianFinder {
    return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
    // 保证minHeap和maxHeap的元素个数最多相差1，且minHeap的所有元素大于等于maxHeap的所有元素
    if mf.minHeap.Len() == mf.maxHeap.Len() {
        heap.Push(&mf.minHeap, num)
        heap.Push(&mf.maxHeap, -(heap.Pop(&mf.minHeap).(int)))
    } else {
        heap.Push(&mf.maxHeap, -num)
        heap.Push(&mf.minHeap, -(heap.Pop(&mf.maxHeap).(int)))
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    if mf.minHeap.Len() == mf.maxHeap.Len() {
        return float64(mf.minHeap[0]-mf.maxHeap[0]) / 2.0
    } else {
        return float64(-mf.maxHeap[0])
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
