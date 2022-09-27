/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 *
 * https://leetcode.cn/problems/super-egg-drop/description/
 *
 * algorithms
 * Hard (30.22%)
 * Likes:    875
 * Dislikes: 0
 * Total Accepted:    70.3K
 * Total Submissions: 232.6K
 * Testcase Example:  '1\n2'
 *
 * 给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
 *
 * 已知存在楼层 f ，满足 0  ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
 *
 * 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1
 * ）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
 *
 * 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
 *
 *
 * 示例 1：
 *
 *
 * 输入：k = 1, n = 2
 * 输出：2
 * 解释：
 * 鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
 * 否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
 * 如果它没碎，那么肯定能得出 f = 2 。
 * 因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
 *
 *
 * 示例 2：
 *
 *
 * 输入：k = 2, n = 6
 * 输出：3
 *
 *
 * 示例 3：
 *
 *
 * 输入：k = 3, n = 14
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */
package lc887

import (
	"math"
	"strconv"
)

// @lc code=start
func superEggDrop(k int, n int) int {
	// return solution1(k, n)
	return solution2(k, n)
}

// dp(k, n) = min(max(dp(k-1, i-1), dp(k, n-i)), ...) + 1, i<=i<=n
// dp(k, n) 表示当前k个鸡蛋，有n层楼时，确定楼层f的最小操作次数
func solution1(k, n int) int {
	memo := make(map[string]int)
	var dp func(int, int) int
	dp = func(k, n int) int {
		// base case
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}

		key := strconv.Itoa(k) + "," + strconv.Itoa(n)
		if val, ok := memo[key]; ok {
			return val
		}

		res := math.MaxInt
		// // 穷举所有的可能性，在每一层楼都尝试扔鸡蛋一次
		// for i := 1; i <= n; i++ {
		// 	res = minInt(res,
		// 		1+maxInt(
		// 			dp(k-1, i-1), // 鸡蛋在i层碎了, 说明目标楼层 0<f<=i-1
		// 			dp(k, n-i),   // 鸡蛋在i层没碎, 说明目标楼层 i<=f<n-1, 且鸡蛋可以继续使用
		// 		))
		// }

		// 二分法代替穷举
		lo, hi := 1, n
		for lo <= hi {
			mid := lo + (hi-lo)/2
			broken := dp(k-1, mid-1)  // 碎了
			notBroken := dp(k, n-mid) // 没碎
			if broken > notBroken {
				hi = mid - 1
				res = minInt(res, broken+1)
			} else {
				lo = mid + 1
				res = minInt(res, notBroken+1)
			}
		}

		memo[key] = res
		return res
	}
	return dp(k, n)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[k][m] = n 表示，当前有 k 个鸡蛋，可以尝试扔 m 次鸡蛋，这个条件下最坏情况下最多能确切测试一栋 n 层的楼
func solution2(k, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}

// @lc code=end
