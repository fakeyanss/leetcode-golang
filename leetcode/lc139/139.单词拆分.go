/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=20004
 *
 * [139] 单词拆分
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (57.02%)
 * Likes:    2653
 * Dislikes: 0
 * Total Accepted:    706.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
 *
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 *
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 * 注意，你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s 和 wordDict[i] 仅由小写英文字母组成
 * wordDict 中的所有字符串 互不相同
 *
 *
 */
package lc139

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp[i] 表示s[:i+1]是否可以被拆分
	// dp[i] = word && dp[j], word为j到i之间的单词
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	wordMap := make(map[string]bool, len(wordDict)) // 方便查找word
	for _, v := range wordDict {
		wordMap[v] = true
	}

	for i := 1; i <= n; i++ {
		// 遍历j正序倒序皆可，倒序更快
		for j := i - 1; j >= 0; j-- { // j到i之间存在任一个单词满足拆分条件即可
			if dp[j] && wordMap[s[j:i]] {
				// 如果前 j 个字符可以被拆分，且 s[j:i] 在字典中
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n["leet", "code"]\n
// @lcpr case=end

// @lcpr case=start
// "applepenapple"\n["apple", "pen"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats", "dog", "sand", "and", "cat"]\n
// @lcpr case=end

*/
