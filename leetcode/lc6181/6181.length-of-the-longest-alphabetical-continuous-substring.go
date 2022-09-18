// https://leetcode.cn/problems/length-of-the-longest-alphabetical-continuous-substring/

package lc6181

func longestContinuousSubstring(s string) int {
	start, ans := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1]+1 {
			ans = maxInt(ans, i-start)
			start = i
		}
	}
	return maxInt(ans, len(s)-start)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
