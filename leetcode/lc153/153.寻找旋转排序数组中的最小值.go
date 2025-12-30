/*
 * @lc app=leetcode.cn id=153 lang=golang
 * @lcpr version=20004
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (58.06%)
 * Likes:    1190
 * Dislikes: 0
 * Total Accepted:    565K
 * Total Submissions: 971.8K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7]
 * 在变化后可能得到：
 *
 * 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
 * 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
 *
 *
 * 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2],
 * ..., a[n-2]] 。
 *
 * 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
 *
 * 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,4,5,1,2]
 * 输出：1
 * 解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [4,5,6,7,0,1,2]
 * 输出：0
 * 解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [11,13,15,17]
 * 输出：11
 * 解释：原数组为 [11,13,15,17] ，旋转 4 次得到输入数组。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 5000
 * -5000 <= nums[i] <= 5000
 * nums 中的所有整数 互不相同
 * nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转
 *
 *
 */
package lc153

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMin(nums []int) int {
	// 二分，每次查找到mid，观察左右是否有序，至少有一侧是有序的
	// l, r := 0, len(nums)-1
	// res := nums[0]
	// for l <= r {
	// 	mid := l + (r-l)/2
	// 	if nums[l] <= nums[mid] { // 左侧有序
	// 		res = min(res, nums[l])
	// 		l = mid + 1
	// 	} else {
	// 		res = min(res, nums[mid])
	// 		r = mid - 1
	// 	}
	// }
	// return res

	// 每次和nums[r]比较更简单
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid // 此时不知道mid是不是最小，所以要保留mid，丢掉右半边
		} else {
			l = mid + 1 // 此时mid已经不是最小，所以丢掉mid和左半边
		}
	}
	return nums[l] // 最后l和r相等
}

// @lc code=end

/*
// @lcpr case=start
// [3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [11,13,15,17]\n
// @lcpr case=end

*/
