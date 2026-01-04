/*
 * @lc app=leetcode.cn id=2998 lang=golang
 * @lcpr version=30305
 *
 * [2998] 使 X 和 Y 相等的最少操作次数
 *
 * https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal/description/
 *
 * algorithms
 * Medium (50.15%)
 * Likes:    14
 * Dislikes: 0
 * Total Accepted:    5.4K
 * Total Submissions: 10.7K
 * Testcase Example:  '26\n1\n54\n2\n25\n30'
 *
 * 给你两个正整数 x 和 y 。
 *
 * 一次操作中，你可以执行以下四种操作之一：
 *
 *
 * 如果 x 是 11 的倍数，将 x 除以 11 。
 * 如果 x 是 5 的倍数，将 x 除以 5 。
 * 将 x 减 1 。
 * 将 x 加 1 。
 *
 *
 * 请你返回让 x 和 y 相等的 最少 操作次数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：x = 26, y = 1
 * 输出：3
 * 解释：我们可以通过以下操作将 26 变为 1 ：
 * 1. 将 x 减 1
 * 2. 将 x 除以 5
 * 3. 将 x 除以 5
 * 将 26 变为 1 最少需要 3 次操作。
 *
 *
 * 示例 2：
 *
 * 输入：x = 54, y = 2
 * 输出：4
 * 解释：我们可以通过以下操作将 54 变为 2 ：
 * 1. 将 x 加 1
 * 2. 将 x 除以 11
 * 3. 将 x 除以 5
 * 4. 将 x 加 1
 * 将 54 变为 2 最少需要 4 次操作。
 *
 *
 * 示例 3：
 *
 * 输入：x = 25, y = 30
 * 输出：5
 * 解释：我们可以通过以下操作将 25 变为 30 ：
 * 1. 将 x 加 1
 * 2. 将 x 加 1
 * 3. 将 x 加 1
 * 4. 将 x 加 1
 * 5. 将 x 加 1
 * 将 25 变为 30 最少需要 5 次操作。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= x, y <= 10^4
 *
 *
 */
package lc2998

// @lc code=start
func minimumOperationsToMakeEqual(x int, y int) int {
	memo := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if x <= y {
			return y - x
		}
		if v, ok := memo[x]; ok {
			return v
		}
		res := min(x-y, // -1
			dfs(x/11)+x%11+1,      // 先-1到11倍数再/11
			dfs(x/11+1)+11-x%11+1, // 先+1到11倍数再/11
			dfs(x/5)+x%5+1,
			dfs(x/5+1)+5-x%5+1)
		memo[x] = res
		return res
	}
	return dfs(x)
}

// @lc code=end

/*
// @lcpr case=start
// 26\n1\n
// @lcpr case=end

// @lcpr case=start
// 54\n2\n
// @lcpr case=end

// @lcpr case=start
// 25\n30\n
// @lcpr case=end

*/
