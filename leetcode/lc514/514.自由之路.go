/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 *
 * https://leetcode.cn/problems/freedom-trail/description/
 *
 * algorithms
 * Hard (50.89%)
 * Likes:    250
 * Dislikes: 0
 * Total Accepted:    24.9K
 * Total Submissions: 48.9K
 * Testcase Example:  '"godding"\n"gd"'
 *
 * 电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
 *
 * 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
 *
 * 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00
 * 方向对齐，然后按下中心按钮，以此逐个拼写完 key 中的所有字符。
 *
 * 旋转 ring 拼出 key 字符 key[i] 的阶段中：
 *
 *
 * 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00
 * 方向对齐，并且这个字符必须等于字符 key[i] 。
 * 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key
 * 的下一个字符（下一阶段）, 直至完成所有拼写。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 *
 *
 * 输入: ring = "godding", key = "gd"
 * 输出: 4
 * 解释:
 * ⁠对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
 * ⁠对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
 * ⁠当然, 我们还需要1步进行拼写。
 * ⁠因此最终的输出是 4。
 *
 *
 * 示例 2:
 *
 *
 * 输入: ring = "godding", key = "godding"
 * 输出: 13
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ring.length, key.length <= 100
 * ring 和 key 只包含小写英文字母
 * 保证 字符串 key 一定可以由字符串  ring 旋转拼出
 *
 *
 */
package lc514

import "math"

// @lc code=start
var charToIndex map[byte][]int
var memo [][]int

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	memo = make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	// 记录ring中字符的索引
	charToIndex = make(map[byte][]int)
	for i := 0; i < m; i++ {
		if _, ok := charToIndex[ring[i]]; !ok {
			charToIndex[ring[i]] = make([]int, 0)
		}
		charToIndex[ring[i]] = append(charToIndex[ring[i]], i)
	}

	return dp(ring, 0, key, 0)
}

// 「状态」就是「当前需要输入的字符」和「当前圆盘指针的位置」，也就是 i 和 j 两个变量
// 「选择」就是「如何拨动指针得到待输入的字符」，字符可能存在列表，且可以顺时针或逆时针
// dp定义：当圆盘指针指向 ring[i] 时，输入字符串 key[j..] 至少需要 dp(ring, i, key, j) 次操作
// 穷举，判断每次拨动的字母，字母不止一个时，应选择哪一个操作; 每次顺时针和逆时针方向，可以选择步数小的方向
func dp(ring string, i int, key string, j int) int {
	if j == len(key) {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	res := math.MaxInt
	for _, k := range charToIndex[key[j]] {
		// 拨动指针次数
		delta := int(math.Abs(float64(k - i)))
		// 顺时针或逆时针，选小的方向
		delta = minInt(delta, len(ring)-delta)
		// 将指针拨到 ring[k]，继续输入 key[j+1..]
		subProblem := dp(ring, k, key, j+1)
		// 选择「整体」操作次数最少的，加一是因为按动按钮也是一次操作
		res = minInt(res, subProblem+delta+1)
	}

	memo[i][j] = res
	return res
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
