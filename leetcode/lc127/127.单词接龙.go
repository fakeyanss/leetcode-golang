/*
 * @lc app=leetcode.cn id=127 lang=golang
 * @lcpr version=20004
 *
 * [127] 单词接龙
 *
 * https://leetcode.cn/problems/word-ladder/description/
 *
 * algorithms
 * Hard (49.28%)
 * Likes:    1436
 * Dislikes: 0
 * Total Accepted:    236.5K
 * Total Submissions: 479.3K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 ->
 * s2 -> ... -> sk：
 *
 *
 * 每一对相邻的单词只差一个字母。
 * 对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
 * sk == endWord
 *
 *
 * 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
 * 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
 *
 *
 * 示例 1：
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log","cog"]
 * 输出：5
 * 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
 *
 *
 * 示例 2：
 *
 * 输入：beginWord = "hit", endWord = "cog", wordList =
 * ["hot","dot","dog","lot","log"]
 * 输出：0
 * 解释：endWord "cog" 不在字典中，所以无法进行转换。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= beginWord.length <= 10
 * endWord.length == beginWord.length
 * 1 <= wordList.length <= 5000
 * wordList[i].length == beginWord.length
 * beginWord、endWord 和 wordList[i] 由小写英文字母组成
 * beginWord != endWord
 * wordList 中的所有字符串 互不相同
 *
 *
 */
package lc127

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	words := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		words[word] = struct{}{}
	}
	if _, ok := words[endWord]; !ok {
		return 0
	}
	delete(words, beginWord)

	// bfs
	q := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	step := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			// 穷举 cur 修改一个字符能得到的单词
			cur := q[0]
			q = q[1:]
			chars := []rune(cur)
			for j := 0; j < len(chars); j++ {
				origin := chars[j]
				// 对每一位穷举 26 个字母
				for c := 'a'; c <= 'z'; c++ {
					if c == origin {
						continue
					}
					chars[j] = c
					newWord := string(chars)
					// 如果构成的新单词在 words 中，就是找到了一个可行的下一步
					_, ok1 := words[newWord]
					_, ok2 := visited[newWord]
					if ok1 && !ok2 {
						if newWord == endWord {
							return step + 1
						}
						q = append(q, newWord)
						visited[newWord] = struct{}{}
					}
				}
				chars[j] = origin // 恢复
			}
		}
		step++
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]\n
// @lcpr case=end

// @lcpr case=start
// "hit"\n"cog"\n["hot","dot","dog","lot","log"]\n
// @lcpr case=end

*/
