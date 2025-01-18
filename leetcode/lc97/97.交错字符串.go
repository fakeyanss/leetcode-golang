/*
 * @lc app=leetcode.cn id=97 lang=golang
 * @lcpr version=20004
 *
 * [97] 交错字符串
 *
 * https://leetcode.cn/problems/interleaving-string/description/
 *
 * algorithms
 * Medium (45.52%)
 * Likes:    1069
 * Dislikes: 0
 * Total Accepted:    172.5K
 * Total Submissions: 378.5K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
 *
 * 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
 *
 *
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m| <= 1
 * 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 +
 * ...
 *
 *
 * 注意：a + b 意味着字符串 a 和 b 连接。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * 输出：false
 *
 *
 * 示例 3：
 *
 * 输入：s1 = "", s2 = "", s3 = ""
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s1.length, s2.length <= 100
 * 0 <= s3.length <= 200
 * s1、s2、和 s3 都由小写英文字母组成
 *
 *
 *
 *
 * 进阶：您能否仅使用 O(s2.length) 额外的内存空间来解决它?
 *
 */
package lc97

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	// if len(s1)+len(s3) != len(s3) {
	// 	return false
	// }
	// i, j, k := 0, 0, 0
	// for k < len(s3) {
	// 	if s1[i] == s3[k] {
	// 		i++
	// 		k++
	// 	} else if s2[j] == s3[k] {
	// 		j++
	// 		k++
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true
	// 不能双指针遍历，至少要有回退的逻辑
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) bool // dp(i,j)表示s1[0..i-1]和s2[0..j-1]判断的结果
	dp = func(i, j int) bool {
		if i == 0 && j == 0 {
			return true
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}
		valid := false
		if i > 0 && s1[i-1] == s3[i+j-1] {
			valid = dp(i-1, j)
		}
		if j > 0 && s2[j-1] == s3[i+j-1] {
			valid = valid || dp(i, j-1)
		}
		if valid {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}
		return valid
	}

	return dp(m, n)
}

// @lc code=end

/*
// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbcbcac"\n
// @lcpr case=end

// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbbaccc"\n
// @lcpr case=end

// @lcpr case=start
// ""\n""\n""\n
// @lcpr case=end

*/
