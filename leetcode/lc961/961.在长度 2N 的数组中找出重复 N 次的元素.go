/*
 * @lc app=leetcode.cn id=961 lang=golang
 * @lcpr version=30305
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 *
 * https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/description/
 *
 * algorithms
 * Easy (70.69%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    96.6K
 * Total Submissions: 136.1K
 * Testcase Example:  '[1,2,3,3]\n[2,1,2,5,3,2]\n[5,1,5,2,5,3,5,4]'
 *
 * 给你一个整数数组 nums ，该数组具有以下属性：
 *
 *
 *
 *
 * nums.length == 2 * n.
 * nums 包含 n + 1 个 不同的 元素
 * nums 中恰有一个元素重复 n 次
 *
 *
 * 找出并返回重复了 n 次的那个元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,2,3,3]
 * 输出：3
 *
 *
 * 示例 2：
 *
 * 输入：nums = [2,1,2,5,3,2]
 * 输出：2
 *
 *
 * 示例 3：
 *
 * 输入：nums = [5,1,5,2,5,3,5,4]
 * 输出：5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 5000
 * nums.length == 2 * n
 * 0 <= nums[i] <= 10^4
 * nums 由 n + 1 个 不同的 元素组成，且其中一个元素恰好重复 n 次
 *
 *
 */
package lc961

// @lc code=start
func repeatedNTimes(nums []int) int {
	// 数组长度2n，目标元素出现n次，且有n+1个不同元素，所以其他元素都只出现了1次
	// 摩尔投票
	// 除掉nums[0]后，剩余2n-1个数，可以进行摩尔投票，找到众数，和leetcode 169相同
	res, cnt := 0, 0
	for _, v := range nums[1:] {
		if v == nums[0] {
			return v
		}
		if cnt == 0 {
			res = v
		}
		if res == v {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,2,5,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,5,2,5,3,5,4]\n
// @lcpr case=end

*/
