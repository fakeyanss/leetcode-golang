/*
 * @lc app=leetcode.cn id=287 lang=golang
 * @lcpr version=30305
 *
 * [287] 寻找重复数
 *
 * https://leetcode.cn/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (67.07%)
 * Likes:    2715
 * Dislikes: 0
 * Total Accepted:    581.9K
 * Total Submissions: 866.9K
 * Testcase Example:  '[1,3,4,2,2]\n[3,1,3,4,2]\n[3,3,3,3,3]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
 *
 * 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
 *
 * 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,3,4,2,2]
 * 输出：2
 *
 *
 * 示例 2：
 *
 * 输入：nums = [3,1,3,4,2]
 * 输出：3
 *
 *
 * 示例 3 :
 *
 * 输入：nums = [3,3,3,3,3]
 * 输出：3
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^5
 * nums.length == n + 1
 * 1 <= nums[i] <= n
 * nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如何证明 nums 中至少存在一个重复的数字?
 * 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
 *
 *
 */
package lc287

// @lc code=start
func findDuplicate(nums []int) int {
	// 构建idx->nums[idx]的链表，环的入口就是重复元素。
	// 因为重复元素一定有多个节点指向它，即入度大于1，而题目数据是只有
	// 一个重复的数，所以[1,n]所有元素都存在，也就是说一定有环，所以可以
	// 用快慢指针找到环入口
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,4,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,3,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,3,3]\n
// @lcpr case=end

*/
