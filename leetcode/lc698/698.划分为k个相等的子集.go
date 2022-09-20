/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 *
 * https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/description/
 *
 * algorithms
 * Medium (41.41%)
 * Likes:    814
 * Dislikes: 0
 * Total Accepted:    84K
 * Total Submissions: 197.3K
 * Testcase Example:  '[4,3,2,3,5,2,1]\n4'
 *
 * 给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
 * 输出： True
 * 说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
 *
 * 示例 2:
 *
 *
 * 输入: nums = [1,2,3,4], k = 3
 * 输出: false
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= len(nums) <= 16
 * 0 < nums[i] < 10000
 * 每个元素的频率在 [1,4] 范围内
 *
 *
 */
package lc698

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k
	used := 0 // 位图，二进制每一位表示nums[i]是否使用
	memo := make(map[int]bool)
	// 从k号桶装到1号桶，初始从nums[0]开始选择
	return backtrack(memo, k, 0, nums, 0, used, target)
}

func backtrack(memo map[int]bool, k, bucket int, nums []int, start, used, target int) bool {
	// base case
	if k == 0 {
		// 所有桶装满，nums一定全用完
		return true
	}
	if bucket == target {
		// 当前k号桶已经装满，穷举装下一个桶
		res := backtrack(memo, k-1, 0, nums, 0, used, target)
		// 记忆结果
		memo[used] = res
		return res
	}

	// 记忆搜索
	if val, ok := memo[used]; ok {
		return val
	}

	for i := start; i < len(nums); i++ {
		// 剪枝
		if (used>>i)&1 == 1 {
			// nums[i]已经使用
			continue
		}
		if nums[i]+bucket > target {
			continue
		}

		// 做选择，第i位与1或计算，置为1，并给bucket装入nums[i]
		used |= 1 << i
		bucket += nums[i]
		// 穷举
		if backtrack(memo, k, bucket, nums, i+1, used, target) {
			return true
		}
		// 撤销选择，第i位与1异或，变为0
		used ^= 1 << i
		bucket -= nums[i]
	}

	return false
}

// @lc code=end
