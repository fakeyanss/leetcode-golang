/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 *
 * https://leetcode.cn/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (61.67%)
 * Likes:    384
 * Dislikes: 0
 * Total Accepted:    88.2K
 * Total Submissions: 143K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 *
 * 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
 * firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
 *
 * 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 * 解释：
 * 航班编号        1   2   3   4   5
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       20  20
 * 预订记录 3 ：       25  25  25  25
 * 总座位数：      10  55  45  25  25
 * 因此，answer = [10,55,45,25,25]
 *
 *
 * 示例 2：
 *
 *
 * 输入：bookings = [[1,2,10],[2,2,15]], n = 2
 * 输出：[10,25]
 * 解释：
 * 航班编号        1   2
 * 预订记录 1 ：   10  10
 * 预订记录 2 ：       15
 * 总座位数：      10  25
 * 因此，answer = [10,25]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 2 * 10^4
 * 1 <= bookings.length <= 2 * 10^4
 * bookings[i].length == 3
 * 1 <= firsti <= lasti <= n
 * 1 <= seatsi <= 10^4
 *
 *
 */
package lc

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	diffArr := Constructor1109(nums)

	for _, booking := range bookings {
		i := booking[0] - 1
		j := booking[1] - 1
		val := booking[2]
		diffArr.increment(i, j, val)
	}

	return diffArr.result()
}

// 差分数组
type DiffArr struct {
	diff []int
}

// 输入一个初始数组，区间操作将在这个数组上进行
func Constructor1109(nums []int) DiffArr {
	length := len(nums)
	diff := make([]int, length)
	diff[0] = nums[0]
	for i := 1; i < length; i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return DiffArr{diff}
}

// 给闭区间 [i, j] 增加 val（可以是负数）
func (d DiffArr) increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

// 返回结果数组
func (d DiffArr) result() []int {
	length := len(d.diff)
	res := make([]int, length)
	res[0] = d.diff[0]
	for i := 1; i < length; i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

// @lc code=end
