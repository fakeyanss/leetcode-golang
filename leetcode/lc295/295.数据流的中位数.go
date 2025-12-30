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
    "sort"
)

// @lc code=start
type MedianFinder struct {
    left  hp // 入堆元素取相反数，即最大堆
    right hp // 最小堆
    // 保证 left 的大小和 right 的大小最多相差1
    // 保证 left 的所有元素都小于等于 right 的所有元素
}

func Constructor() MedianFinder {
    return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
    if mf.left.Len() == mf.right.Len() {
        heap.Push(&mf.right, num)
        heap.Push(&mf.left, -heap.Pop(&mf.right).(int))
    } else {
        heap.Push(&mf.left, -num)
        heap.Push(&mf.right, -heap.Pop(&mf.left).(int))
    }
}

func (mf *MedianFinder) FindMedian() float64 {
    if mf.left.Len() > mf.right.Len() {
        return float64(-mf.left.IntSlice[0])
    }
    return float64(mf.right.IntSlice[0]-mf.left.IntSlice[0]) / 2.0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any   { x := h.IntSlice[h.Len()-1]; h.IntSlice = h.IntSlice[:h.Len()-1]; return x }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
