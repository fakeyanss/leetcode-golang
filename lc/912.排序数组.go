/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode.cn/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.67%)
 * Likes:    656
 * Dislikes: 0
 * Total Accepted:    437.1K
 * Total Submissions: 785.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */
package lc

import "math/rand"

// @lc code=start
func sortArray(nums []int) []int {
	return quickSort(nums)
}

func mergeSort(nums []int) []int {
	n := len(nums)
	temp := make([]int, n)

	var merge func([]int, int, int, int)
	merge = func(nums []int, lo, mid, hi int) {
		for i := lo; i <= hi; i++ {
			temp[i] = nums[i]
		}
		i, j := lo, mid+1
		for p := lo; p <= hi; p++ {
			if i == mid+1 {
				// 左半边数组已全部被合并
				nums[p] = temp[j]
				j++
			} else if j == hi+1 {
				// 右半边数组已全部被合并
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

	var sort func([]int, int, int)
	sort = func(nums []int, lo, hi int) {
		if lo == hi {
			// 单个元素不用排序
			return
		}
		mid := lo + (hi-lo)/2
		sort(nums, lo, mid)
		sort(nums, mid+1, hi)
		merge(nums, lo, mid, hi)
	}

	sort(nums, 0, n-1)
	return nums
}

func quickSort(nums []int) []int {
	var shuffle = func(nums []int) {
		n := len(nums)
		for i := 0; i < n; i++ {
			r := rand.Intn(n)
			nums[i], nums[r] = nums[r], nums[i]
		}
	}
	// 为了避免出现耗时的极端情况，先随机打乱
	shuffle(nums)

	var partition = func(nums []int, lo, hi int) int {
		pivot := nums[lo]
		// i, j 定义为开区间，同时定义：[lo, i) <= pivot；(j, hi] > pivot
		i, j := lo+1, hi
		for i <= j {
			for i < hi && nums[i] <= pivot {
				// 结束时恰好 nums[i] > pivot
				i++
			}
			for j > lo && nums[j] > pivot {
				// 结束时恰好 nums[j] <= pivot
				j--
			}
			// 此时 [lo, i) <= pivot && (j, hi] > pivot

			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 将 pivot 放到合适的位置，即 pivot 左边元素较小，右边元素较大
		nums[lo], nums[j] = nums[j], nums[lo]
		return j
	}

	var sort func([]int, int, int)
	sort = func(nums []int, lo, hi int) {
		if lo >= hi {
			return
		}

		p := partition(nums, lo, hi)
		sort(nums, lo, p-1)
		sort(nums, p+1, hi)
	}
	sort(nums, 0, len(nums)-1)
	return nums
}

// @lc code=end
