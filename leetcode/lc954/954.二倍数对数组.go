/*
 * @lc app=leetcode.cn id=954 lang=golang
 * @lcpr version=20004
 *
 * [954] 二倍数对数组
 *
 * https://leetcode.cn/problems/array-of-doubled-pairs/description/
 *
 * algorithms
 * Medium (39.36%)
 * Likes:    199
 * Dislikes: 0
 * Total Accepted:    45.2K
 * Total Submissions: 114.8K
 * Testcase Example:  '[3,1,3,6]'
 *
 * 给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 *
 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [3,1,3,6]
 * 输出：false
 *
 *
 * 示例 2：
 *
 * 输入：arr = [2,1,2,6]
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：arr = [4,-2,2,-4]
 * 输出：true
 * 解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= arr.length <= 3 * 10^4
 * arr.length 是偶数
 * -10^5 <= arr[i] <= 10^5
 *
 *
 */
package lc954

import (
	"sort"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canReorderDoubled(arr []int) bool {
	if len(arr)%2 != 0 {
		return false
	}
	mp := make(map[int]int)
	for _, val := range arr {
		mp[val]++
	}

	sort.Ints(arr)
	lastNegative, firstPositive := -1, -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			lastNegative = i
		} else {
			firstPositive = i
			break
		}
	}
	if lastNegative != -1 {
		for i := lastNegative; i >= 0; i-- {
			if mp[arr[i]] > 0 && mp[arr[i]*2] > 0 {
				mp[arr[i]]--
				mp[arr[i]*2]--
			}
		}
	}
	if firstPositive != -1 {
		for i := firstPositive; i < len(arr); i++ {
			if mp[arr[i]] > 0 && mp[arr[i]*2] > 0 {
				mp[arr[i]]--
				mp[arr[i]*2]--
			}
		}
	}

	for _, val := range mp {
		if val != 0 {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,3,6]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,2,6]\n
// @lcpr case=end

// @lcpr case=start
// [4,-2,2,-4]\n
// @lcpr case=end

*/
