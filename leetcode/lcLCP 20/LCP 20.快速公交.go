/*
 * @lc app=leetcode.cn id=LCP 20 lang=golang
 * @lcpr version=30305
 *
 * [LCP 20] 快速公交
 *
 * https://leetcode.cn/problems/meChtZ/description/
 *
 * algorithms
 * Hard (37.00%)
 * Likes:    47
 * Dislikes: 0
 * Total Accepted:    3.1K
 * Total Submissions: 8.5K
 * Testcase Example:  '31\n5\n3\n[6]\n[10]\n612\n4\n5\n[3,6,8,11,5,10,4]\n[4,7,6,3,7,6,4]'
 *
 * 小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：
 * - 小扣从 `x` 号站点移动至 `x + 1` 号站点需要花费的时间为 `inc`；
 * - 小扣从 `x` 号站点移动至 `x - 1` 号站点需要花费的时间为 `dec`。
 *
 * 现有 `m` 辆公交车，编号为 `0` 到 `m-1`。小扣也可以通过搭乘编号为 `i` 的公交车，从 `x` 号站点移动至 `jump[i]*x`
 * 号站点，耗时仅为 `cost[i]`。小扣可以搭乘任意编号的公交车且搭乘公交次数不限。
 *
 * 假定小扣起始站点记作 `0`，秋日市集站点记作 `target`，请返回小扣抵达秋日市集最少需要花费多少时间。由于数字较大，最终答案需要对
 * 1000000007 (1e9 + 7) 取模。
 *
 * 注意：小扣可在移动过程中到达编号大于 `target` 的站点。
 *
 * **示例 1：**
 * >输入：`target = 31, inc =  5, dec = 3, jump = [6], cost = [10]`
 * >
 * >输出：`33`
 * >
 * >解释：
 * >小扣步行至 1 号站点，花费时间为 5；
 * >小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10；
 * >小扣从 6 号站台步行至 5 号站台，花费时间为 3；
 * >小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10；
 * >小扣从 30 号站台步行至 31 号站台，花费时间为 5；
 * >最终小扣花费总时间为 33。
 *
 *
 * **示例 2：**
 * >输入：`target = 612, inc =  4, dec = 5, jump = [3,6,8,11,5,10,4], cost =
 * [4,7,6,3,7,6,4]`
 * >
 * >输出：`26`
 * >
 * >解释：
 * >小扣步行至 1 号站点，花费时间为 4；
 * >小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4；
 * >小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3；
 * >小扣从 33 号站台步行至 34 站台，花费时间为 4；
 * >小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4；
 * >小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7；
 * >最终小扣花费总时间为 26。
 *
 *
 * **提示：**
 * - `1 <= target <= 10^9`
 * - `1 <= jump.length, cost.length <= 10`
 * - `2 <= jump[i] <= 10^6`
 * - `1 <= inc, dec, cost[i] <= 10^6`
 */
package lcp20

// @lc code=start
func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	memo := map[int]int{}
	// dfs(cur) 表示从0到cur最小耗时
	var dfs func(int) int
	dfs = func(cur int) int {
		// 边界条件：cur ≤ 1时，直接步行到0，耗时cur*inc
		if cur <= 1 {
			return cur * inc
		}
		if v, ok := memo[cur]; ok {
			return v
		}
		res := cur * inc // 初始最大值
		// 枚举公交路线
		for i := range jump {
			j := jump[i]
			c := cost[i]
			mod := cur % j
			div := cur / j
			// 从子状态dfs(cur/j)步行前进到车站，即花费多次inc，再花费时间c坐车，可以到达dfs(cur)
			res = min(res, dfs(div)+mod*inc+c)
			// 从子状态dfs(cur/j+1)步行后退到车站，即花费多次dec，再花费时间c坐车，可以到达dfs(cur)
			// mod=0时二者相同，所以仅mod≠0时计算第二种方案，剪枝
			if mod != 0 {
				res = min(res, dfs(div+1)+(j-mod)*dec+c)
			}
		}
		memo[cur] = res
		return res
	}
	return dfs(target) % 1000000007
}

// @lc code=end

/*
// @lcpr case=start
// 31\n5\n3\n[6]\n[10]\n
// @lcpr case=end

// @lcpr case=start
// 612\n4\n5\n[3,6,8,11,5,10,4]\n[4,7,6,3,7,6,4]\n
// @lcpr case=end

*/
