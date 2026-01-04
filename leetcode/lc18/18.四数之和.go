/*
 * @lc app=leetcode.cn id=18 lang=golang
 * @lcpr version=30305
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.89%)
 * Likes:    2125
 * Dislikes: 0
 * Total Accepted:    757.9K
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0\n[2,2,2,2,2]\n8'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 *
 *
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * 你可以按 任意顺序 返回答案 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * 示例 2：
 *
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */
package lc18

import "slices"

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	var res [][]int
	n := len(nums)
	for i := range n - 3 { // 枚举第一个数
		x := nums[i]
		if i > 0 && x == nums[i-1] { // 跳过重复数
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ { // 枚举第二个数
			y := nums[j]
			if j > i+1 && y == nums[j-1] { // 跳过重复数
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target {
				continue
			}
			l, r := j+1, n-1 // 双指针枚举第三个和第四个数
			for l < r {
				sum := x + y + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{x, y, nums[l], nums[r]})
					for ; l < r && nums[l] == nums[l+1]; l++ {
					}
					for ; l < r && nums[r] == nums[r-1]; r-- {
					}
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,-1,0,-2,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n8\n
// @lcpr case=end

*/
