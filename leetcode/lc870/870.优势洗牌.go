/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 *
 * https://leetcode.cn/problems/advantage-shuffle/description/
 *
 * algorithms
 * Medium (47.30%)
 * Likes:    194
 * Dislikes: 0
 * Total Accepted:    27.3K
 * Total Submissions: 57.8K
 * Testcase Example:  '[2,7,11,15]\n[1,10,4,11]'
 *
 * 给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums 的优势可以用满足 nums1[i] > nums2[i] 的索引 i
 * 的数目来描述。
 *
 * 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
 * 输出：[2,11,7,15]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
 * 输出：[24,32,8,12]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length <= 10^5
 * nums2.length == nums1.length
 * 0 <= nums1[i], nums2[i] <= 10^9
 *
 *
 */
package lc870

import (
	"sort"
)

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	b := make([][2]int, len(nums2))
	for i, n := range nums2 {
		b[i] = [2]int{n, i}
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})

	result := make([]int, len(nums1))
	left, right := 0, len(nums1)-1
	for i := len(b) - 1; i >= 0; i-- {
		val, idx := b[i][0], b[i][1]
		if nums1[right] > val {
			result[idx] = nums1[right]
			right--
		} else {
			result[idx] = nums1[left]
			left++
		}
	}
	return result
	// // nums1 升序排列
	// sort.Ints(nums1)

	// // nums2 降序排序，大顶堆
	// h := make(IntHeap, 0)
	// for i, val := range nums2 {
	// 	h = append(h, []int{i, val})
	// }
	// heap.Init(&h)

	// l, r := 0, len(nums1)-1
	// res := make([]int, len(nums1))
	// for h.Len() > 0 {
	// 	pair := h.Pop()
	// 	i, max := (pair.([]int))[0], (pair.([]int))[1]
	// 	if max < nums1[r] {
	// 		// 如果nums1[r]大于max，那就填入这个值
	// 		res[i] = nums1[r]
	// 		r--
	// 	} else {
	// 		// 否则用最小最混一下，田忌赛马
	// 		res[i] = nums1[l]
	// 		l++
	// 	}
	// }
	// return res
}

// // 大顶堆，实现heap的interface
// // h[x][0] 存储 index，h[x][1] 存储 value
// type IntHeap [][]int

// // 绑定Len方法
// func (h IntHeap) Len() int {
// 	return len(h)
// }

// // 绑定Less方法，这里用的是小于号，生成的是小根堆
// func (h IntHeap) Less(i, j int) bool {
// 	return h[i][1] > h[j][1]
// }

// // 绑定swap方法
// func (h IntHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// // pop、push 使用指针，因为改变了slice长度
// // 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	item := old[0]
// 	*h = old[1:n]
// 	return item
// }

// // 绑定push方法
// func (h *IntHeap) Push(x interface{}) {
// 	*h = append(*h, x.([]int))
// }

// @lc code=end
