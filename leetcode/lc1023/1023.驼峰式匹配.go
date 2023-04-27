/*
 * @lc app=leetcode.cn id=1023 lang=golang
 *
 * [1023] 驼峰式匹配
 *
 * https://leetcode.cn/problems/camelcase-matching/description/
 *
 * algorithms
 * Medium (64.54%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    28.7K
 * Total Submissions: 44.4K
 * Testcase Example:  '["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FB"'
 *
 * 给你一个字符串数组 queries，和一个表示模式的字符串 pattern，请你返回一个布尔数组 answer 。只有在待查项 queries[i]
 * 与模式串 pattern 匹配时， answer[i] 才为 true，否则为 false。
 *
 * 如果可以将小写字母插入模式串 pattern 得到待查询项
 * query，那么待查询项与给定模式串匹配。可以在任何位置插入每个字符，也可以不插入字符。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FB"
 * 输出：[true,false,true,true,false]
 * 示例：
 * "FooBar" 可以这样生成："F" + "oo" + "B" + "ar"。
 * "FootBall" 可以这样生成："F" + "oot" + "B" + "all".
 * "FrameBuffer" 可以这样生成："F" + "rame" + "B" + "uffer".
 *
 * 示例 2：
 *
 *
 * 输入：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBa"
 * 输出：[true,false,true,false,false]
 * 解释：
 * "FooBar" 可以这样生成："Fo" + "o" + "Ba" + "r".
 * "FootBall" 可以这样生成："Fo" + "ot" + "Ba" + "ll".
 *
 *
 * 示例 3：
 *
 *
 * 输入：queries =
 * ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern =
 * "FoBaT"
 * 输出：[false,true,false,false,false]
 * 解释：
 * "FooBarTest" 可以这样生成："Fo" + "o" + "Ba" + "r" + "T" + "est".
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= pattern.length, queries.length <= 100
 * 1 <= queries[i].length <= 100
 * queries[i] 和 pattern 由英文字母组成
 *
 *
 */
package lc1023

import "unicode"

// @lc code=start
// 如果 pattern 是 queries[i] 的子序列，并且去掉 pattern 之后 queries[i] 的剩余部分都由小写字母构成，那么 queries[i] 就与 pattern 匹配
func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		ans[i] = true
		p, m := 0, len(pattern)
		for _, c := range queries[i] {
			if p < m && pattern[p] == byte(c) {
				p++
			} else if unicode.IsUpper(c) {
				ans[i] = false
				break
			}
		}
		if p < m {
			ans[i] = false
		}
	}
	return ans
}

// @lc code=end
