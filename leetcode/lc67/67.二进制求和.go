/*
 * @lc app=leetcode.cn id=67 lang=golang
 * @lcpr version=20004
 *
 * [67] 二进制求和
 *
 * https://leetcode.cn/problems/add-binary/description/
 *
 * algorithms
 * Easy (53.57%)
 * Likes:    1260
 * Dislikes: 0
 * Total Accepted:    437.5K
 * Total Submissions: 816.7K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
 *
 *
 *
 * 示例 1：
 *
 * 输入:a = "11", b = "1"
 * 输出："100"
 *
 * 示例 2：
 *
 * 输入：a = "1010", b = "1011"
 * 输出："10101"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= a.length, b.length <= 10^4
 * a 和 b 仅由字符 '0' 或 '1' 组成
 * 字符串如果不是 "0" ，就不含前导零
 *
 *
 */

package lc67

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func addBinary(a string, b string) string {
	m, n := len(a)-1, len(b)-1
	res := make([]byte, max(m, n)+1+1)
	k := len(res) - 1
	flag := 0
	for m >= 0 && n >= 0 {
		sum := int(a[m]-'0') + int(b[n]-'0') + flag
		flag = 0
		if sum >= 2 {
			flag = 1
			if sum == 2 {
				res[k] = '0'
			}
			if sum == 3 {
				res[k] = '1'
			}
		} else {
			if sum == 1 {
				res[k] = '1'
			}
			if sum == 0 {
				res[k] = '0'
			}
		}
		m--
		n--
		k--
	}
	if m < 0 && n >= 0 {
		for n >= 0 {
			sum := int(b[n]-'0') + flag
			flag = 0
			if sum == 2 {
				res[k] = '0'
				flag = 1
			} else if sum == 1 {
				res[k] = '1'
			} else {
				res[k] = '0'
			}
			n--
			k--
		}
	}
	if n < 0 && m >= 0 {
		for m >= 0 {
			sum := int(a[m]-'0') + flag
			flag = 0
			if sum == 2 {
				res[k] = '0'
				flag = 1
			} else if sum == 1 {
				res[k] = '1'
			} else {
				res[k] = '0'
			}
			m--
			k--
		}
	}
	if flag == 1 {
		res[0] = '1'
	} else {
		res = res[1:]
	}

	return string(res)
}

// @lc code=end

/*
// @lcpr case=start
// "11"\n"1"\n
// @lcpr case=end

// @lcpr case=start
// "1010"\n"1011"\n
// @lcpr case=end

*/
