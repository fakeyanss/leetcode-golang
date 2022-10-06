/*
 * @lc app=leetcode.cn id=927 lang=golang
 *
 * [927] 三等分
 *
 * https://leetcode.cn/problems/three-equal-parts/description/
 *
 * algorithms
 * Hard (35.45%)
 * Likes:    172
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 41.2K
 * Testcase Example:  '[1,0,1,0,1]'
 *
 * 给定一个由 0 和 1 组成的数组 arr ，将数组分成  3 个非空的部分 ，使得所有这些部分表示相同的二进制值。
 *
 * 如果可以做到，请返回任何 [i, j]，其中 i+1 < j，这样一来：
 *
 *
 * arr[0], arr[1], ..., arr[i] 为第一部分；
 * arr[i + 1], arr[i + 2], ..., arr[j - 1] 为第二部分；
 * arr[j], arr[j + 1], ..., arr[arr.length - 1] 为第三部分。
 * 这三个部分所表示的二进制值相等。
 *
 *
 * 如果无法做到，就返回 [-1, -1]。
 *
 * 注意，在考虑每个部分所表示的二进制时，应当将其看作一个整体。例如，[1,1,0] 表示十进制中的 6，而不会是 3。此外，前导零也是被允许的，所以
 * [0,1,1] 和 [1,1] 表示相同的值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：arr = [1,0,1,0,1]
 * 输出：[0,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：arr = [1,1,0,1,1]
 * 输出：[-1,-1]
 *
 * 示例 3:
 *
 *
 * 输入：arr = [1,1,0,0,1]
 * 输出：[0,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 * 3 <= arr.length <= 3 * 10^4
 * arr[i] 是 0 或 1
 *
 *
 */
package lc927

// @lc code=start
func threeEqualParts(arr []int) []int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}
	partial := sum / 3
	// 数组的末尾是无法改变的，因此区间 [third, arr.length−1] 所表示的二进制值可以固定
	first, second, third, cur := 0, 0, 0, 0
	for i, val := range arr {
		if val == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == partial*2 {
				third = i
			}
			cur++
		}
	}

	// length 表示二进制值的长度
	length := len(arr) - third
	if first+length <= second && second+length <= third {
		i := 0
		for third+i < len(arr) {
			// 以third为基准，比较first和second的元素1和1之间的0是否排列相同
			if arr[i+first] != arr[i+third] || arr[i+second] != arr[i+third] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}

// @lc code=end
