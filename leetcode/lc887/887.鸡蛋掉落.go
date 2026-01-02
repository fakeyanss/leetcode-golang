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

// @lc code=start
func superEggDrop(k int, n int) int {
	// // dp[k][n] 表示k个鸡蛋n层楼，确定临界楼层的最坏情况的最小扔鸡蛋次数
	// // dp[k][n] = min(1 + max(dp[k][n-x], dp[k-1][x-1])), x是扔鸡蛋的楼层, 1<=x<=n
	// // 需要枚举所有可能的扔鸡蛋楼层，枚举加速可以用二分
	// dp := make([][]int, k+1)
	// for i := range dp {
	// 	dp[i] = make([]int, n+1)
	// 	dp[i][1] = 1
	// }
	// for j := range dp[1] {
	// 	dp[1][j] = j // 1个鸡蛋，n层楼得扔n次
	// }
	// for i := 2; i <= k; i++ {
	// 	for j := 2; j <= n; j++ {
	// 		dp[i][j] = j // 初始化最坏情况j层楼需要j次
	// 		// // 枚举所有楼层
	// 		// for x := 1; x <= j; x++ {
	// 		// 	broken := dp[i-1][x-1]
	// 		// 	notBroken := dp[i][j-x]
	// 		// 	dp[i][j] = min(dp[i][j], 1+max(broken, notBroken))
	// 		// }
	// 		// 二分加速
	// 		lo, hi := 1, j
	// 		for lo <= hi {
	// 			mid := lo + (hi-lo)/2
	// 			broken := dp[i-1][mid-1]
	// 			notBroken := dp[i][j-mid]
	// 			if broken < notBroken {
	// 				lo = mid + 1
	// 				dp[i][j] = min(dp[i][j], 1+notBroken)
	// 			} else {
	// 				hi = mid - 1
	// 				dp[i][j] = min(dp[i][j], 1+broken)
	// 			}
	// 		}
	// 	}
	// }
	// return dp[k][n]

	// dp[m][k]=n，表示k个鸡蛋扔m次，可以确定的最大的楼层数为n。
	// 转换问题为1<=m<=n，已知k，n，求m.
	// 也就是在递推计算过程中，当dp>=n时，就可以得到扔鸡蛋的次数为m。
	// 解释：扔出1枚鸡蛋，如果碎了，那么下面的楼层个数为dp[m-1][k-1]；
	// 如果没碎，那么上面的楼层个数位dp[m-1][k]，
	// 所以dp[k][m] = 1 + dp[k-1][m-1] + dp[k][m-1]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][1] = i
	}
	for j := range dp[1] {
		dp[1][j] = 1
	}

	var m int
	for m = 2; m <= n; m++ { // 遍历所有的扔鸡蛋次数
		for j := 2; j <= k; j++ {
			dp[m][j] = 1 + dp[m-1][j] + dp[m-1][j-1]
		}
		if dp[m][k] >= n {
			return m
		}
	}
	return n // 包含边界n=1 or k=1的情况
}

// @lc code=end
