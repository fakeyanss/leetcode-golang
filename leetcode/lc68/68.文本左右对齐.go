/*
 * @lc app=leetcode.cn id=68 lang=golang
 * @lcpr version=20004
 *
 * [68] 文本左右对齐
 *
 * https://leetcode.cn/problems/text-justification/description/
 *
 * algorithms
 * Hard (55.46%)
 * Likes:    451
 * Dislikes: 0
 * Total Accepted:    95.5K
 * Total Submissions: 172.1K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 *
 * 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth
 * 个字符。
 *
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 *
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 *
 * 注意:
 *
 *
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: words = ["This", "is", "an", "example", "of", "text", "justification."],
 * maxWidth = 16
 * 输出:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
 * 输出:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。
 * ⁠    第二行同样为左对齐，这是因为这行只包含一个单词。
 *
 *
 * 示例 3:
 *
 * 输入:words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth
 * = 20
 * 输出:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] 由小写英文字母和符号组成
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */
package lc68

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	var tmp []string
	wLen, wLenPadding := 0, 0
	for _, w := range words {
		if wLenPadding+len(w)-1 >= maxWidth {
			spaceNum := 0
			last := 0
			var line string
			if len(tmp) == 1 {
				spaceNum = maxWidth - len(tmp[0])
				line += tmp[0]
				for i := 0; i < spaceNum; i++ {
					line += " "
				}
			} else {
				spaceNum = (maxWidth - wLen) / (len(tmp) - 1)
				last = (maxWidth - wLen) % (len(tmp) - 1)
				for i, ww := range tmp {
					line += ww
					if i == len(tmp)-1 {
						continue
					}
					for j := 0; j < spaceNum; j++ {
						line += " "
					}
					if i < last {
						line += " "
					}
				}
			}

			res = append(res, line)

			wLen = 0
			wLenPadding = 0
			tmp = []string{}
		}
		wLen += len(w)
		wLenPadding += len(w) + 1
		tmp = append(tmp, w)
	}
	var line string
	for i, ww := range tmp {
		line += ww
		if i == len(tmp)-1 {
			ll := len(line)
			for i := 0; i < maxWidth-ll; i++ {
				line += " "
			}
		} else {
			line += " "
		}
	}
	res = append(res, line)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["This", "is", "an", "example", "of", "text", "justification."]\n16\n
// @lcpr case=end

// @lcpr case=start
// ["What","must","be","acknowledgment","shall","be"]\n16\n
// @lcpr case=end

// @lcpr case=start
// ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we"\n20\n
// @lcpr case=end

*/
