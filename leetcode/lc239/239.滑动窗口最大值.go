/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode.cn/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (49.94%)
 * Likes:    1999
 * Dislikes: 0
 * Total Accepted:    380.9K
 * Total Submissions: 762.6K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回 滑动窗口中的最大值 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
 * 输出：[3,3,5,5,6,7]
 * 解释：
 * 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1], k = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 *
 *
 */
package lc239

// @lc code=start
type MonotonicQueue struct {
	q []int
}

func (mq *MonotonicQueue) push(n int) {
	que := mq.q
	// 将末尾小于n的元素全部删除
	for len(que) > 0 && que[len(que)-1] < n {
		que = que[:len(que)-1]
	}
	que = append(que, n)
	// 注意append在需要扩容时不返回原本的slice
	mq.q = que
}

func (mq *MonotonicQueue) max() int {
	return mq.q[0]
}

func (mq *MonotonicQueue) pop(n int) {
	q := mq.q
	if n == q[0] {
		q = q[1:]
	}
	mq.q = q
}

func maxSlidingWindow(nums []int, k int) []int {
	window := &MonotonicQueue{q: []int{}}
	res := []int{}

	for i := 0; i < len(nums); i++ {
		// 窗口扩大
		window.push(nums[i])
		if i < k-1 {
			continue
		}
		// 记录最大值
		res = append(res, window.max())
		// 窗口缩小
		window.pop(nums[i-k+1])
	}
	return res
}

// @lc code=end
