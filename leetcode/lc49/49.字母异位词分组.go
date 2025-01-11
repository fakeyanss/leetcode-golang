/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=20004
 *
 * [49] 字母异位词分组
 *
 * https://leetcode.cn/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (69.07%)
 * Likes:    2073
 * Dislikes: 0
 * Total Accepted:    903.9K
 * Total Submissions: 1.3M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
 *
 * 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
 *
 *
 *
 * 示例 1:
 *
 * 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
 *
 * 示例 2:
 *
 * 输入: strs = [""]
 * 输出: [[""]]
 *
 *
 * 示例 3:
 *
 * 输入: strs = ["a"]
 * 输出: [["a"]]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] 仅包含小写字母
 *
 *
 */
package lc49

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func groupAnagrams(strs []string) [][]string {
	encode := func(str string) string {
		count := make([]byte, 26)
		for _, c := range str {
			count[c-'a']++
		}
		return string(count)
	}

	codeToGroup := make(map[string][]string)
	for _, s := range strs {
		code := encode(s)
		codeToGroup[code] = append(codeToGroup[code], s)
	}

	res := make([][]string, 0, len(codeToGroup))
	for _, group := range codeToGroup {
		res = append(res, group)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["eat", "tea", "tan", "ate", "nat", "bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
