/*
 * @lc app=leetcode.cn id=75 lang=golang
 * @lcpr version=30305
 *
 * [75] 颜色分类
 *
 * https://leetcode.cn/problems/sort-colors/description/
 *
 * algorithms
 * Medium (63.35%)
 * Likes:    2023
 * Dislikes: 0
 * Total Accepted:    894.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '[2,0,2,1,1,0]\n[2,0,1]'
 *
 * 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地 对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 *
 *
 *
 * 必须在不使用库内置的 sort 函数的情况下解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,0,2,1,1,0]
 * 输出：[0,0,1,1,2,2]
 *
 *
 * 示例 2：
 *
 * 输入：nums = [2,0,1]
 * 输出：[0,1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 300
 * nums[i] 为 0、1 或 2
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */
package lc75

// @lc code=start
func sortColors(nums []int) {
	// p0, p2 := 0, len(nums)-1 // p0,p2指向左右
	// p := 0                   // p遍历数组
	// for p <= p2 {
	// 	if nums[p] == 0 {
	// 		nums[p0], nums[p] = nums[p], nums[p0]
	// 		p0++
	// 		p++ // p左边一定都是0，所以这里可以p++
	// 	} else if nums[p] == 2 {
	// 		nums[p2], nums[p] = nums[p], nums[p2]
	// 		p2--
	// 	} else {
	// 		p++
	// 	}
	// }

	p0, p1 := 0, 0
	for i, x := range nums {
		nums[i] = 2
		if x <= 1 {
			nums[p1] = 1
			p1++
		}
		if x == 0 {
			nums[p0] = 0
			p0++
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [2,0,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [2,0,1]\n
// @lcpr case=end

*/
