/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.54%)
 * Likes:    1861
 * Dislikes: 0
 * Total Accepted:    731.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4], k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6], k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
package lc215

import "math/rand"

// @lc code=start
// 最小堆
// func findKthLargest(nums []int, k int) int {
// 	h := &hp{}
// 	heap.Init(h)
// 	for _, v := range nums {
// 		heap.Push(h, v)
// 		if h.Len() > k {
// 			heap.Pop(h)
// 		}
// 	}
// 	return heap.Pop(h).(int)
// }

// type hp []int

// func (h hp) Len() int           { return len(h) }
// func (h hp) Less(i, j int) bool { return h[i] < h[j] }
// func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(x any)        { *h = append(*h, x.(int)) }
// func (h *hp) Pop() any          { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }

// (快排)快速选择
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k // 第k大的元素idx是n-k
	left, right := 0, n-1
	for {
		i := partition(nums, left, right)
		if i == target {
			return nums[i]
		}
		if i > target {
			right = i - 1
		} else {
			left = i + 1
		}
	}
}

// 在子数组 [left, right] 中随机选择一个基准元素 pivot
// 根据 pivot 重新排列子数组 [left, right]
// 重新排列后，<= pivot 的元素都在 pivot 的左侧，>= pivot 的元素都在 pivot 的右侧
// 返回 pivot 在重新排列后的 nums 中的下标
// 特别地，如果子数组的所有元素都等于 pivot，我们会返回子数组的中心下标，避免退化
func partition(nums []int, left, right int) int {
	// 1. 在子数组 [left, right] 中随机选择一个基准元素 pivot
	i := left + rand.Intn(right-left+1)
	pivot := nums[i]
	nums[i], nums[left] = nums[left], nums[i]

	// 2. 相向双指针遍历子数组 [left + 1, right]
	// 循环不变量：在循环过程中，子数组的数据分布始终如下图
	// [ pivot | <=pivot | 尚未遍历 | >=pivot ]
	//   ^                 ^     ^         ^
	//   left              i     j         right
	i, j := left+1, right
	for {
		for i <= j && nums[i] < pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	// 循环结束后
	// [ pivot | <=pivot | >=pivot ]
	//   ^             ^   ^     ^
	//   left          j   i     right

	// 3. 把 pivot 与 nums[j] 交换，完成划分（partition）
	nums[left], nums[j] = nums[j], nums[left]

	// 交换后
	// [ <=pivot | pivot | >=pivot ]
	//               ^
	//               j

	return j
}

// @lc code=end
