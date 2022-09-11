/*
 * @lc app=leetcode.cn id=1094 lang=golang
 *
 * [1094] 拼车
 *
 * https://leetcode.cn/problems/car-pooling/description/
 *
 * algorithms
 * Medium (53.93%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    52.6K
 * Total Submissions: 97.5K
 * Testcase Example:  '[[2,1,5],[3,3,7]]\n4'
 *
 * 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
 *
 * 给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i
 * 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
 *
 * 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 4
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：trips = [[2,1,5],[3,3,7]], capacity = 5
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= trips.length <= 1000
 * trips[i].length == 3
 * 1 <= numPassengersi <= 100
 * 0 <= fromi < toi <= 1000
 * 1 <= capacity <= 10^5
 *
 *
 */
package lc1094

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	diff := Constructor1094(nums)

	for _, trip := range trips {
		val := trip[0]
		i := trip[1]
		// 第 trip[2] 站乘客已经下车，
		// 即乘客在车上的区间是 [trip[1], trip[2] - 1]
		j := trip[2] - 1
		diff.increment(i, j, val)
	}

	for _, cnt := range diff.result() {
		if cnt > capacity {
			return false
		}
	}

	return true
}

// 差分数组
type Diff1094 struct {
	diff []int
}

// 输入一个初始数组，区间操作将在这个数组上进行
func Constructor1094(nums []int) Diff1094 {
	length := len(nums)
	diff := make([]int, length)
	diff[0] = nums[0]
	for i := 1; i < length; i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return Diff1094{diff}
}

// 给闭区间 [i,j] 增加 val
func (d Diff1094) increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d Diff1094) result() []int {
	length := len(d.diff)
	res := make([]int, length)
	res[0] = d.diff[0]
	for i := 1; i < length; i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}

// @lc code=end
