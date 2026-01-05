/*
 * @lc app=leetcode.cn id=60 lang=golang
 * @lcpr version=30305
 *
 * [60] 排列序列
 *
 * https://leetcode.cn/problems/permutation-sequence/description/
 *
 * algorithms
 * Hard (54.50%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    160.3K
 * Total Submissions: 293.9K
 * Testcase Example:  '3\n3\n4\n9\n3\n1'
 *
 * 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 3, k = 3
 * 输出："213"
 *
 *
 * 示例 2：
 *
 * 输入：n = 4, k = 9
 * 输出："2314"
 *
 *
 * 示例 3：
 *
 * 输入：n = 3, k = 1
 * 输出："123"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 9
 * 1 <= k <= n!
 *
 *
 */
package lc60

import (
	"strconv"
	"strings"
)

// @lc code=start
func getPermutation(n int, k int) string {
	var res string
	used := make([]bool, n+1)
	var path []string
	count := 0

	var dfs func() bool
	dfs = func() bool {
		if len(path) == n {
			count++
			if count == k {
				res = strings.Join(path, "")
				return true
			}
			return false
		}

		// 从1-n遍历回溯，天然字典序
		for i := 1; i <= n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, strconv.Itoa(i))
			if dfs() {
				return true // 提前结束
			}
			path = path[:len(path)-1]
			used[i] = false
		}
		return false
	}
	dfs()
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3\n3\n
// @lcpr case=end

// @lcpr case=start
// 4\n9\n
// @lcpr case=end

// @lcpr case=start
// 3\n1\n
// @lcpr case=end

*/
