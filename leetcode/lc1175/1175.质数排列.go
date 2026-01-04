/*
 * @lc app=leetcode.cn id=1175 lang=golang
 * @lcpr version=30305
 *
 * [1175] 质数排列
 *
 * https://leetcode.cn/problems/prime-arrangements/description/
 *
 * algorithms
 * Easy (57.39%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    37.4K
 * Total Submissions: 65.2K
 * Testcase Example:  '5\n100'
 *
 * 请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。
 *
 * 让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。
 *
 * 由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 5
 * 输出：12
 * 解释：举个例子，[1,2,5,4,3] 是一个有效的排列，但 [5,2,3,4,1] 不是，因为在第二种情况里质数 5 被错误地放在索引为 1
 * 的位置上。
 *
 *
 * 示例 2：
 *
 * 输入：n = 100
 * 输出：682289015
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 100
 *
 *
 */
package lc1175

// @lc code=start
func numPrimeArrangements(n int) int {
	primeCnt := 0
	if n < 2 {
		primeCnt = 0
	}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCnt++
		}
	}
	// 总排列数 = 质数排列数 * 合数排列数
	return (factorial(primeCnt) * factorial(n-primeCnt)) % 1000000007
}

// 判断质数
func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// 求阶层，即排列个数，结果取模
func factorial(k int) int {
	res := 1
	for i := 1; i <= k; i++ {
		res = (res * i) % 1000000007
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 100\n
// @lcpr case=end

*/
