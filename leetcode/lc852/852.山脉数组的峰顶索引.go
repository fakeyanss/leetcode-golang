/*
 * @lc app=leetcode.cn id=852 lang=golang
 * @lcpr version=30305
 *
 * [852] 山脉数组的峰顶索引
 *
 * https://leetcode.cn/problems/peak-index-in-a-mountain-array/description/
 *
 * algorithms
 * Medium (67.96%)
 * Likes:    471
 * Dislikes: 0
 * Total Accepted:    206.5K
 * Total Submissions: 303.8K
 * Testcase Example:  '[0,1,0]\n[0,2,1,0]\n[0,10,5,2]'
 *
 * 给定一个长度为 n 的整数 山脉 数组 arr ，其中的值递增到一个 峰值元素 然后递减。
 *
 * 返回峰值元素的下标。
 *
 * 你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [0,1,0]
 * 输出：1
 *
 *
 * 示例 2：
 *
 * 输入：arr = [0,2,1,0]
 * 输出：1
 *
 *
 * 示例 3：
 *
 * 输入：arr = [0,10,5,2]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= arr.length <= 10^5
 * 0 <= arr[i] <= 10^6
 * 题目数据 保证 arr 是一个山脉数组
 *
 *
 */
package lc852

// @lc code=start
// 思路：二分
func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > arr[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,2,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,10,5,2]\n
// @lcpr case=end

*/
