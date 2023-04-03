/*
 * @lc app=leetcode.cn id=1053 lang=golang
 *
 * [1053] 交换一次的先前排列
 *
 * https://leetcode.cn/problems/previous-permutation-with-one-swap/description/
 *
 * algorithms
 * Medium (47.69%)
 * Likes:    118
 * Dislikes: 0
 * Total Accepted:    22.5K
 * Total Submissions: 47.1K
 * Testcase Example:  '[3,2,1]'
 *
 * 给你一个正整数数组 arr（可能存在重复的元素），请你返回可在 一次交换（交换两数字 arr[i] 和 arr[j] 的位置）后得到的、按字典序排列小于
 * arr 的最大排列。
 *
 * 如果无法这么操作，就请返回原数组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [3,2,1]
 * 输出：[3,1,2]
 * 解释：交换 2 和 1
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,1,5]
 * 输出：[1,1,5]
 * 解释：已经是最小排列
 *
 *
 * 示例 3：
 *
 *
 * 输入：arr = [1,9,4,6,7]
 * 输出：[1,7,4,6,9]
 * 解释：交换 9 和 7
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^4
 *
 *
 */
package lc1053

// @lc code=start
// 贪心算法。
// 要一次交换找到字典序排列小于 arr 的最大排列，也即是找到下标i和j，
// 让arr[i]>arr[j]，0<=i<j<n，且arr[j]=max(arr[i..n]).
// 换句话说，就是找到i尽可能靠右侧；找到j为下标i到n中的最大值，如果结果有多个，选择最左侧的j.
func prevPermOpt1(arr []int) []int {
	n := len(arr)
	// i从右往左，找到第一个
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			// 只用比较arr[i]和arr[i+1]，因为不会存在j>i+1使arr[i]>arr[j]，
			// 也就是arr[i]右侧数值是非递减的
			j := n - 1
			for arr[j] >= arr[i] || arr[j] == arr[j-1] {
				// 直到找到比arr[i]小的最大的arr[j]，如果有相等，找到左边界.
				// 这里也可以用二分查找左边界来优化
				j--
			}
			// 原地交换
			arr[i], arr[j] = arr[j], arr[i]
			break
		}
	}
	return arr
}

// @lc code=end
