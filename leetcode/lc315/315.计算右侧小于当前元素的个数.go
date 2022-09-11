/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode.cn/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (42.82%)
 * Likes:    870
 * Dislikes: 0
 * Total Accepted:    70K
 * Total Submissions: 163.4K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i]
 * 右侧小于 nums[i] 的元素的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,6,1]
 * 输出：[2,1,1,0]
 * 解释：
 * 5 的右侧有 2 个更小的元素 (2 和 1)
 * 2 的右侧仅有 1 个更小的元素 (1)
 * 6 的右侧有 1 个更小的元素 (1)
 * 1 的右侧有 0 个更小的元素
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-1]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [-1,-1]
 * 输出：[0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */
package lc315

// @lc code=start

// 记录数组的元素值与下标
type Pair struct {
	val int
	id  int
}

// 归并排序使用的临时数组
var temp []Pair

// 记录结果
var count []int

func countSmaller(nums []int) []int {
	n := len(nums)
	temp = make([]Pair, n)
	count = make([]int, n)

	// 记录元素原始的索引位置，以便在 count 数组中更新结果
	arr := make([]Pair, n)
	for i, val := range nums {
		arr[i] = Pair{val, i}
	}
	sort(arr, 0, n-1)
	return count
}

// 归并排序
func sort(arr []Pair, lo, hi int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	sort(arr, lo, mid)
	sort(arr, mid+1, hi)
	merge(arr, lo, mid, hi)
}

// 合并两个有序数组
func merge(arr []Pair, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		temp[i] = arr[i]
	}

	i, j := lo, mid+1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			arr[p] = temp[j]
			j++
		} else if j == hi+1 {
			arr[p] = temp[i]
			i++
			// 更新count数组, 此时右半边的值全部小于当前temp[i]
			// 多次归并时可能会重复计算某一个temp[i], 所以需要累加
			count[arr[p].id] += j - mid - 1
		} else if temp[i].val > temp[j].val {
			arr[p] = temp[j]
			j++
		} else {
			arr[p] = temp[i]
			i++
			// 更新count数组, 此时temp[mid+1]到temp[j]之间的值，都小于temp[i]
			count[arr[p].id] += j - mid - 1
		}
	}
}

// @lc code=end
