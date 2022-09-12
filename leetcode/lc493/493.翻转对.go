/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 *
 * https://leetcode.cn/problems/reverse-pairs/description/
 *
 * algorithms
 * Hard (35.75%)
 * Likes:    372
 * Dislikes: 0
 * Total Accepted:    34.3K
 * Total Submissions: 95.8K
 * Testcase Example:  '[1,3,2,3,1]'
 *
 * 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
 *
 * 你需要返回给定数组中的重要翻转对的数量。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,2,3,1]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,4,3,5,1]
 * 输出: 3
 *
 *
 * 注意:
 *
 *
 * 给定数组的长度不会超过50000。
 * 输入数组中的所有数字都在32位整数的表示范围内。
 *
 *
 */
package lc493

// @lc code=start

var temp []int
var count int

func reversePairs(nums []int) int {
	n := len(nums)
	temp = make([]int, n)
	count = 0
	sort(nums, 0, n-1)
	return count
}

func sort(nums []int, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort(nums, lo, mid)
	sort(nums, mid+1, hi)
	merge(nums, lo, mid, hi)
}

func merge(nums []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = nums[i]
	}

	// 对于左半边的每个 nums[i]，都去右半边寻找符合条件的元素
	// for i := lo; i <= mid; i++ {
	// 	for j := mid + 1; j <= hi; j++ {
	// 		if nums[i] > nums[j]*2 {
	// 			count++
	// 		}
	// 	}
	// }

	// 效率优化, 维护左闭右开区间 [mid+1, end) 中的元素乘 2 小于 nums[i]
	end := mid + 1
	for i := lo; i <= mid; i++ {
		for end <= hi && nums[i] > nums[end]*2 {
			end++
		}
		count += end - mid - 1
	}

	// 合并
	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			nums[p] = temp[j]
			j++
		} else if j == hi+1 {
			nums[p] = temp[i]
			i++
		} else if temp[i] > temp[j] {
			nums[p] = temp[j]
			j++
		} else {
			nums[p] = temp[i]
			i++
		}
	}
}

// @lc code=end
