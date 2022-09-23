/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 *
 * https://leetcode.cn/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (52.80%)
 * Likes:    547
 * Dislikes: 0
 * Total Accepted:    104.3K
 * Total Submissions: 197.5K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 *
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 *
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 *
 * 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 *
 *
 * 示例 2:
 *
 *
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
 *
 *
 * 示例 3:
 *
 *
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：无法旋转到目标数字且不被锁定。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= deadends.length <= 500
 * deadends[i].length == 4
 * target.length == 4
 * target 不在 deadends 之中
 * target 和 deadends[i] 仅由若干位数字组成
 *
 *
 */
package lc752

// @lc code=start
func openLock(deadends []string, target string) int {
	// 双向bfs
	// 记录走过的密码, deadends也当做走过的密码，可以达到跳过的目的
	visited := make(map[string]bool)
	for _, s := range deadends {
		visited[s] = true
	}

	// 从起点开始bfs
	q1 := make(map[string]bool)

	q1["0000"] = true
	step := 0

	// 从终点开始bfs
	q2 := make(map[string]bool)
	q2[target] = true

	for len(q1) > 0 && len(q2) > 0 {
		// 先遍历小的一边
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		temp := make(map[string]bool)
		for cur := range q1 {
			if visited[cur] {
				continue
			}
			if q2[cur] {
				return step
			}

			visited[cur] = true

			// 将当前锁数字的下一个可能的数字加入集合
			for i := 0; i < 4; i++ {
				up := plus(cur, i)
				if !visited[up] {
					temp[up] = true
				}

				down := minus(cur, i)
				if !visited[down] {
					temp[down] = true
				}
			}
		}

		// 交换q1和q2
		q1 = q2
		q2 = temp

		step++
	}

	// 穷举完仍不能到达target
	return -1
}

// 将s[i]向上转动一次
func plus(s string, i int) string {
	b := []byte(s)
	if b[i] == '9' {
		b[i] = '0'
	} else {
		b[i] += 1
	}
	return string(b)
}

// 将s[i]向下转动一次
func minus(s string, i int) string {
	b := []byte(s)
	if b[i] == '0' {
		b[i] = '9'
	} else {
		b[i] -= 1
	}
	return string(b)
}

// @lc code=end
