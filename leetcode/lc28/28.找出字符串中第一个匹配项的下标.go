/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Medium (41.52%)
 * Likes:    1603
 * Dislikes: 0
 * Total Accepted:    741.3K
 * Total Submissions: 1.8M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0 ，所以返回 0 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */
package lc0028

// 本题即实现strStr()

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	// 暴力求解
	// p1, p2 := len(haystack), len(needle)
	// for i := 0; i <= p1-p2; i++ {
	// 	if haystack[i:i+p2] == needle {
	// 		return i
	// 	}
	// }
	// return -1

	// KMP, too hard
	k := construct(needle)
	return k.search(haystack)
}

type kmp struct {
	dp  [][]int
	pat string
}

func construct(pat string) kmp {
	m, n := len(pat), 256 // 256个ascii码
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	k := kmp{dp, pat}

	// base case
	dp[0][int(pat[0])] = 1
	// 影子状态初始为0
	x := 0
	for i := 1; i < m; i++ {
		// 遍历256个ascii码
		for j := 0; j < n; j++ {
			if int(pat[i]) == j {
				dp[i][j] = i + 1
			} else {
				dp[i][j] = dp[x][j]
			}
		}
		// 更新x为上一个与pat[i]有公共前缀的下标位置
		x = dp[x][int(pat[i])]
	}

	return k
}

func (k kmp) search(txt string) int {
	m, n := len(k.pat), len(txt)
	j := 0
	for i := 0; i < n; i++ {
		j = k.dp[j][int(txt[i])]
		if j == m {
			return i - j + 1
		}
	}
	return -1
}

// @lc code=end
