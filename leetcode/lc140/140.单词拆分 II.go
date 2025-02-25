/*
 * @lc app=leetcode.cn id=140 lang=golang
 * @lcpr version=20004
 *
 * [140] 单词拆分 II
 *
 * https://leetcode.cn/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (59.84%)
 * Likes:    776
 * Dislikes: 0
 * Total Accepted:    107.1K
 * Total Submissions: 178.7K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序
 * 返回所有这些可能的句子。
 *
 * 注意：词典中的同一个单词可能在分段中被重复使用多次。
 *
 *
 *
 * 示例 1：
 *
 * 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
 * 输出:["cats and dog","cat sand dog"]
 *
 *
 * 示例 2：
 *
 * 输入:s = "pineapplepenapple", wordDict =
 * ["apple","pen","applepen","pine","pineapple"]
 * 输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
 * 输出:[]
 *
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 20
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 10
 * s 和 wordDict[i] 仅有小写英文字母组成
 * wordDict 中所有字符串都 不同
 *
 *
 */
package lc140

import "strings"

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	// 139是dp，140需要所有拆分结果，需要回溯
	var res []string
	var track []string

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(s) {
			res = append(res, strings.Join(track, " "))
			return
		}

		// 遍历所有单词，尝试匹配 s[i..] 的前缀
		for _, w := range wordDict {
			if i+len(w) > len(s) {
				continue
			}
			subStr := s[i : i+len(w)]
			if subStr != w { // 无法匹配，跳过
				continue
			}
			// 匹配成功
			track = append(track, w)
			backtrack(i + len(w))
			// 回退
			track = track[:len(track)-1]
		}
	}
	backtrack(0) // 判断 s[0..] 是否能够被拼出
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "catsanddog"\n["cat","cats","and","sand","dog"]\n
// @lcpr case=end

// @lcpr case=start
// "pineapplepenapple"\n["apple","pen","applepen","pine","pineapple"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats","dog","sand","and","cat"]\n
// @lcpr case=end

*/
