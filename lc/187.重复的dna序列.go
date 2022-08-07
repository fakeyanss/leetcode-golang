/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 *
 * https://leetcode.cn/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (52.92%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    110K
 * Total Submissions: 207.8K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * DNA序列 由一系列核苷酸组成，缩写为 'A', 'C', 'G' 和 'T'.。
 *
 *
 * 例如，"ACGAATTCCG" 是一个 DNA序列 。
 *
 *
 * 在研究 DNA 时，识别 DNA 中的重复序列非常有用。
 *
 * 给定一个表示 DNA序列 的字符串 s ，返回所有在 DNA 分子中出现不止一次的 长度为 10 的序列(子字符串)。你可以按 任意顺序
 * 返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 * 输出：["AAAAACCCCC","CCCCCAAAAA"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "AAAAAAAAAAAAA"
 * 输出：["AAAAAAAAAA"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 10^5
 * s[i]=='A'、'C'、'G' or 'T'
 *
 *
 */
package lc

import "math"

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	// 先把字符串转为4进制的数字
	nums := make([]int, len(s))
	for i := range s {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'G':
			nums[i] = 1
		case 'C':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}

	// 记录重复出现的hash值
	seen := make(map[int]int)
	// 记录重复出现的字符串结果
	res := make([]string, 0)

	// 数字位数
	L := 10
	// 进制
	R := 4
	// 存储 R^(L-1) 的结果
	RL := int(math.Pow(float64(R), float64(L-1)))
	// 维护滑动窗口中字符串的hash值
	windowHash := 0

	left, right := 0, 0
	for right < len(nums) {
		// 扩大窗口，移入字符，并维护窗口哈希值（在最低位添加数字）
		windowHash = R*windowHash + nums[right]
		right++

		// 当子串长度满足要求, 记录结果，缩小窗口
		if right-left == L {
			seen[windowHash]++
			// 存在重复出现多次（大于2）的情况，只在结果中添加一次
			if seen[windowHash] == 2 {
				res = append(res, s[left:right])
			}
			// 缩小窗口，移出字符，并维护窗口哈希值（删除最高位数字）
			windowHash = windowHash - nums[left]*RL
			left++
		}
	}
	return res
}

// @lc code=end
