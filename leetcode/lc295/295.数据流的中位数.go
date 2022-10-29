/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode.cn/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (53.11%)
 * Likes:    767
 * Dislikes: 0
 * Total Accepted:    96.7K
 * Total Submissions: 182K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
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
// 用两个优先队列 queMax 和 queMin 分别记录大于中位数的数和小于等于中位数的数。
// 需要特别注意添加元素时，需要维护queMax和queMin的元素个数相差1，且保证queMax的堆顶不小于queMin的堆顶
type MedianFinder struct {
    queMin hp // 小顶堆, 存负数，pop时取相反数
    queMax hp // 大顶堆
}

func Constructor() MedianFinder {
    return MedianFinder{
        queMin: hp{},
        queMax: hp{},
    }
}

func (mf *MedianFinder) AddNum(num int) {
    minQ, maxQ := &mf.queMin, &mf.queMax
    if minQ.Len() >= maxQ.Len() {
        heap.Push(minQ, -num)
        heap.Push(maxQ, -heap.Pop(minQ).(int))
        return
    }
    heap.Push(maxQ, num)
    heap.Push(minQ, -heap.Pop(maxQ).(int))
}

func (mf *MedianFinder) FindMedian() float64 {
    minQ, maxQ := mf.queMin, mf.queMax
    // 如果元素不一样多，多的那个堆的堆顶元素就是中位数
    if minQ.Len() > maxQ.Len() {
        return float64(-minQ.IntSlice[0])
    } else if minQ.Len() < maxQ.Len() {
        return float64(maxQ.IntSlice[0])
    }
    // 如果元素一样多，两个堆堆顶元素的平均数是中位数
    return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct {
    sort.IntSlice // 默认大顶堆
}

func (h *hp) Push(v interface{}) {
    h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
    arr := h.IntSlice
    v := arr[len(arr)-1]
    h.IntSlice = arr[:len(arr)-1]
    return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
