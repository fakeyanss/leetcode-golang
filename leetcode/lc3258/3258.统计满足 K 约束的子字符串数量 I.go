/*
 * @lc app=leetcode.cn id=3258 lang=golang
 * @lcpr version=20003
 *
 * [3258] 统计满足 K 约束的子字符串数量 I
 *
 * https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/description/
 *
 * algorithms
 * Easy (80.33%)
 * Likes:    41
 * Dislikes: 0
 * Total Accepted:    23.3K
 * Total Submissions: 27.8K
 * Testcase Example:  '"10101"\n1'
 *
 * 给你一个 二进制 字符串 s 和一个整数 k。
 *
 * 如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
 *
 *
 * 字符串中 0 的数量最多为 k。
 * 字符串中 1 的数量最多为 k。
 *
 *
 * 返回一个整数，表示 s 的所有满足 k 约束 的子字符串的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "10101", k = 1
 *
 * 输出：12
 *
 * 解释：
 *
 * s 的所有子字符串中，除了 "1010"、"10101" 和 "0101" 外，其余子字符串都满足 k 约束。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "1010101", k = 2
 *
 * 输出：25
 *
 * 解释：
 *
 * s 的所有子字符串中，除了长度大于 5 的子字符串外，其余子字符串都满足 k 约束。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "11111", k = 1
 *
 * 输出：15
 *
 * 解释：
 *
 * s 的所有子字符串都满足 k 约束。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 50
 * 1 <= k <= s.length
 * s[i] 是 '0' 或 '1'。
 *
 *
 */
package lc3258

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countKConstraintSubstrings2(s string, k int) int {
	n, res := len(s), 0
	for i := 0; i < n; i++ {
		count := [2]int{}
		for j := i; j < n; j++ {
			count[s[j]-'0']++
			if count[0] <= k || count[1] <= k {
				res++
			} else {
				break
			}
		}
	}
	return res
}

// https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/solutions/2884495/on-hua-dong-chuang-kou-pythonjavacgo-by-keubv/
func countKConstraintSubstrings(s string, k int) int {
	cnt := [2]int{}
	res, left := 0, 0
	for i, c := range s {
		// 字符 0 的 ASCII 值是偶数，字符 1 的 ASCII 值是奇数，所以可以用 ASCII 值 cmod2 得到对应的数字。这也等价于和 1 计算 AND。
		cnt[c&1]++
		for cnt[0] > k && cnt[1] > k {
			cnt[s[left]&1]--
			left++
		}
		res += i - left + 1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "10101"\n1\n
// @lcpr case=end

// @lcpr case=start
// "1010101"\n2\n
// @lcpr case=end

// @lcpr case=start
// "11111"\n1\n
// @lcpr case=end

*/
