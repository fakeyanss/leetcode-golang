/*
* @lc app=leetcode.cn id=1845 lang=golang
*
* [1845] 座位预约管理系统
*
* https://leetcode.cn/problems/seat-reservation-manager/description/
*
  - algorithms
  - Medium (47.66%)
  - Likes:    16
  - Dislikes: 0
  - Total Accepted:    13.8K
  - Total Submissions: 28.9K
  - Testcase Example:  '["SeatManager","reserve","reserve","unreserve","reserve","reserve","reserve","reserve","unreserve"]\n' +
    '[[5],[],[],[2],[],[],[],[],[5]]'

*
* 请你设计一个管理 n 个座位预约的系统，座位编号从 1 到 n 。
*
* 请你实现 SeatManager 类：
*
*
* SeatManager(int n) 初始化一个 SeatManager 对象，它管理从 1 到 n 编号的 n
* 个座位。所有座位初始都是可预约的。
* int reserve() 返回可以预约座位的 最小编号 ，此座位变为不可预约。
* void unreserve(int seatNumber) 将给定编号 seatNumber 对应的座位变成可以预约。
*
*
*
*
* 示例 1：
*
* 输入：
* ["SeatManager", "reserve", "reserve", "unreserve", "reserve", "reserve",
* "reserve", "reserve", "unreserve"]
* [[5], [], [], [2], [], [], [], [], [5]]
* 输出：
* [null, 1, 2, null, 2, 3, 4, 5, null]
*
* 解释：
* SeatManager seatManager = new SeatManager(5); // 初始化 SeatManager ，有 5 个座位。
* seatManager.reserve();    // 所有座位都可以预约，所以返回最小编号的座位，也就是 1 。
* seatManager.reserve();    // 可以预约的座位为 [2,3,4,5] ，返回最小编号的座位，也就是 2 。
* seatManager.unreserve(2); // 将座位 2 变为可以预约，现在可预约的座位为 [2,3,4,5] 。
* seatManager.reserve();    // 可以预约的座位为 [2,3,4,5] ，返回最小编号的座位，也就是 2 。
* seatManager.reserve();    // 可以预约的座位为 [3,4,5] ，返回最小编号的座位，也就是 3 。
* seatManager.reserve();    // 可以预约的座位为 [4,5] ，返回最小编号的座位，也就是 4 。
* seatManager.reserve();    // 唯一可以预约的是座位 5 ，所以返回 5 。
* seatManager.unreserve(5); // 将座位 5 变为可以预约，现在可预约的座位为 [5] 。
*
*
*
*
* 提示：
*
*
* 1 <= n <= 10^5
* 1 <= seatNumber <= n
* 每一次对 reserve 的调用，题目保证至少存在一个可以预约的座位。
* 每一次对 unreserve 的调用，题目保证 seatNumber 在调用函数前都是被预约状态。
* 对 reserve 和 unreserve 的调用 总共 不超过 10^5 次。
*
*
*/
package lc1845

import "container/heap"

// @lc code=start

// 优先级队列（二叉堆）解决
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

type SeatManager struct {
	h *IntHeap
}

func Constructor(n int) SeatManager {
	h := &IntHeap{}
	for i := 0; i < n; i++ {
		heap.Push(h, i+1)
	}
	return SeatManager{h}
}

func (sm *SeatManager) Reserve() int {
	return heap.Pop(sm.h).(int)
}

func (sm *SeatManager) Unreserve(seatNumber int) {
	heap.Push(sm.h, seatNumber)
}

/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
// @lc code=end
