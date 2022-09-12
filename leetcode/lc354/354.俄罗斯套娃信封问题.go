/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode.cn/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (40.26%)
 * Likes:    803
 * Dislikes: 0
 * Total Accepted:    91.9K
 * Total Submissions: 228.7K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
 *
 * 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 注意：不允许旋转信封。
 *
 *
 * 示例 1：
 *
 *
 * 输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出：3
 * 解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 * 示例 2：
 *
 *
 * 输入：envelopes = [[1,1],[1,1],[1,1]]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= envelopes.length <= 10^5
 * envelopes[i].length == 2
 * 1 <= wi, hi <= 10^5
 *
 *
 */
package lc354

import "sort"

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	// envelopes = [[w,h], [w,h]...]
	// 先对宽度 w 进行升序排序，如果遇到 w 相同的情况，则按照高度 h 降序排序
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	// 然后转化为对所有的 h 计算 LIS 的长度
	dp := make([]int, 0)
	for _, envelop := range envelopes {
		// 简化实现，直接用sort包的查询方法, 返回左侧边界; 自己实现可以复用lc300的LTS二分查找计算
		i := sort.SearchInts(dp, envelop[1])
		if i == len(dp) {
			dp = append(dp, envelop[1])
		} else {
			dp[i] = envelop[1]
		}
	}
	return len(dp)
}

// @lc code=end
