/*
 * @lc app=leetcode.cn id=1390 lang=golang
 * @lcpr version=30305
 *
 * [1390] 四因数
 *
 * https://leetcode.cn/problems/four-divisors/description/
 *
 * algorithms
 * Medium (42.14%)
 * Likes:    46
 * Dislikes: 0
 * Total Accepted:    15.4K
 * Total Submissions: 33.9K
 * Testcase Example:  '[21,4,7]\n[21,21]\n[1,2,3,4,5]'
 *
 * 给你一个整数数组 nums，请你返回该数组中恰有四个因数的这些整数的各因数之和。如果数组中不存在满足题意的整数，则返回 0 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [21,4,7]
 * 输出：32
 * 解释：
 * 21 有 4 个因数：1, 3, 7, 21
 * 4 有 3 个因数：1, 2, 4
 * 7 有 2 个因数：1, 7
 * 答案仅为 21 的所有因数的和。
 *
 *
 * 示例 2:
 *
 * 输入: nums = [21,21]
 * 输出: 64
 *
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,4,5]
 * 输出: 0
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * 1 <= nums[i] <= 10^5
 *
 *
 */
package lc1390

// @lc code=start
// func sumFourDivisors(nums []int) int {
// 	var res int
// 	for _, v := range nums {
// 		factorCnt, factorSum := 0, 0
// 		for i := 1; i <= v/i; i++ {
// 			if v%i == 0 {
// 				factorCnt++
// 				factorSum += i
// 				if i != v/i { // 如果两个因数不相等，可以将另一个因数也加入结果
// 					factorCnt++
// 					factorSum += v / i
// 				}
// 			}
// 		}
// 		if factorCnt == 4 {
// 			res += factorSum
// 		}
// 	}
// 	return res
// }
const cnt = 100001

var factorCnt, factorSum [cnt]int

func init() {
	for i := 1; i < cnt; i++ {
		for j := i; j < cnt; j += i {
			// i是j的因数
			factorCnt[j]++
			factorSum[j] += i
		}
	}
}

func sumFourDivisors(nums []int) int {
	var res int
	for _, x := range nums {
		if factorCnt[x] == 4 {
			res += factorSum[x]
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [21,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [21,21]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/
