/*
 * @lc app=leetcode.cn id=3266 lang=golang
 * @lcpr version=20004
 *
 * [3266] K 次乘运算后的最终数组 II
 *
 * https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-ii/description/
 *
 * algorithms
 * Hard (23.70%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 16.2K
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
 * k 次操作以后，你需要将 nums 中每一个数值对 10^9 + 7 取余。
 *
 * 请你返回执行完 k 次乘运算以及取余运算之后，最终的 nums 数组。
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
 * 取余操作后
 * [8, 4, 6, 5, 6]
 *
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [100000,2000], k = 2, multiplier = 1000000
 *
 * 输出：[999999307,999999993]
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
 * [100000, 2000000000]
 *
 *
 * 2 次操作后
 * [100000000000, 2000000000]
 *
 *
 * 取余操作后
 * [999999307, 999999993]
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
 * 1 <= nums.length <= 10^4
 * 1 <= nums[i] <= 10^9
 * 1 <= k <= 10^9
 * 1 <= multiplier <= 10^6
 *
 *
 */
package lc3266

import (
	"container/heap"
	"sort"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }

func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

// 快速幂算法
func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3,5,6]\n5\n2\n
// @lcpr case=end

// @lcpr case=start
// [100000,2000]\n2\n1000000\n
// @lcpr case=end

*/
