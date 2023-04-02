/*
 * @lc app=leetcode.cn id=1641 lang=golang
 *
 * [1641] 统计字典序元音字符串的数目
 *
 * https://leetcode.cn/problems/count-sorted-vowel-strings/description/
 *
 * algorithms
 * Medium (82.51%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    38.5K
 * Total Submissions: 46.7K
 * Testcase Example:  '1'
 *
 * 给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。
 *
 * 字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 1
 * 输出：5
 * 解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 2
 * 输出：15
 * 解释：仅由元音组成的 15 个字典序字符串为
 * ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
 * 注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 33
 * 输出：66045
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package lc1641

// @lc code=start

// a e i o u 转为 0 1 2 3 4，
// 结果字符串中前面的数不能大于后面的数，所以结果数由长度n和末尾的字母决定。
// 定义dp[n][m]为长度n，结尾为m的字符串的个数(0<=m<=4)，则
//   dp[n][m] = dp[n-1][0] + dp[n-1][1] + ... + dp[n-1][m-1]
// 因为dp[n]的计算只和dp[n-1]相关，且类似前缀和，所以可以降维优化空间复杂度
func countVowelStrings(n int) int {
	// return dp1(n)
	return dp2(n)
	// return math(n)
}

func dp1(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1 // 初始化
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	return dp[n][4]
}

func dp2(n int) int {
	dp := make([]int, 5)
	for i := 0; i < 5; i++ {
		dp[i] = 1
	}

	for i := 1; i <= n; i++ {
		// 计算长度为i的字符串的个数，每次直接原地修改
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[4]
}

// 数学组合
func math(n int) int {
	return (n + 1) * (n + 2) * (n + 3) * (n + 4) / 24
}

// @lc code=end
