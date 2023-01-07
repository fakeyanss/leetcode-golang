/*
 * @lc app=leetcode.cn id=1803 lang=golang
 *
 * [1803] 统计异或值在范围内的数对有多少
 *
 * https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/description/
 *
 * algorithms
 * Hard (55.60%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    14.7K
 * Total Submissions: 26.4K
 * Testcase Example:  '[1,4,2,7]\n2\n6'
 *
 * 给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数：low 和 high ，请返回 漂亮数对 的数目。
 *
 * 漂亮数对 是一个形如 (i, j) 的数对，其中 0 <= i < j < nums.length 且 low <= (nums[i] XOR
 * nums[j]) <= high 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,4,2,7], low = 2, high = 6
 * 输出：6
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * ⁠   - (0, 1): nums[0] XOR nums[1] = 5
 * ⁠   - (0, 2): nums[0] XOR nums[2] = 3
 * ⁠   - (0, 3): nums[0] XOR nums[3] = 6
 * ⁠   - (1, 2): nums[1] XOR nums[2] = 6
 * ⁠   - (1, 3): nums[1] XOR nums[3] = 3
 * ⁠   - (2, 3): nums[2] XOR nums[3] = 5
 *
 *
 * 示例 2：
 *
 * 输入：nums = [9,8,4,2,1], low = 5, high = 14
 * 输出：8
 * 解释：所有漂亮数对 (i, j) 列出如下：
 * - (0, 2): nums[0] XOR nums[2] = 13
 * - (0, 3): nums[0] XOR nums[3] = 11
 * - (0, 4): nums[0] XOR nums[4] = 8
 * - (1, 2): nums[1] XOR nums[2] = 12
 * - (1, 3): nums[1] XOR nums[3] = 10
 * - (1, 4): nums[1] XOR nums[4] = 9
 * - (2, 3): nums[2] XOR nums[3] = 6
 * - (2, 4): nums[2] XOR nums[4] = 5
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] <= 2 * 10^4
 * 1 <= low <= high <= 2 * 10^4
 *
 *
 */
package lc1803

// @lc code=start

// 参考题解：https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solution/zi-duan-shu-by-lxk1203-l0ee/

// 落于区间[low, high]的个数 ➡ 落于区间[0, high]的个数 - 落于区间[0, low-1]的个数

// 判断数a是否小于等于数b，我们从高位往低位依次判断，假设当前为i：
// bi为0时：如果ai为1，a肯定大于b；如果ai也为0，继续往下一位判断；
// bi为1时：如果ai为0，a肯定小于b；如果ai也为1，继续往下一位判断；
func countPairs(nums []int, low int, high int) int {
	return f(nums, high) - f(nums, low-1)
}

const HIGH_BIT = 14 // nums[i]<=2*10^4，二进制不超过15位

func f(nums []int, x int) int {
	root := &trie{}
	n, res := len(nums), 0
	for i := 1; i < n; i++ {
		// 组合遍历所有的数对:
		// 向trie存入nums[i-1], 再使用nums[i]与nums[0...i-1]做组合
		root.insert(nums[i-1])         // 将前一个数加入trie
		res += root.search(nums[i], x) // 统计[0,i]中异或结果<=x的数对个数
	}
	return res
}

type trie struct {
	// children[0]为左子树，表示当前位置出现0
	// children[1]为右子树，表示当前位置出现1
	children [2]*trie
	// 表示有多少个数字以根结点到该节点路径为前缀
	cnt int
}

func (t *trie) insert(num int) {
	cur := t
	// 从高位往低位遍历
	for k := HIGH_BIT; k >= 0; k-- {
		bit := (num >> k) & 1 // 取第k位的值
		if cur.children[bit] == nil {
			cur.children[bit] = &trie{}
		}
		cur = cur.children[bit]
		cur.cnt++ // 只要有bit出现，以这个路径为前缀的数就增加一个
	}
}

// 对于一个新的数num，从高位(14位)进行遍历。假设遍历到第i位，num第i位的值为n_i, 目标值x的第i位的值为x_i:
// - x_i = 0，那么对于第i位，n_i与字典树中的数的异或结果必须为0，继续判断下一位才有可能符合条件；因此应该往字典树中的n_i路径上走（两个相同的值异或结果为0，所以应该看字典树中n_i的路径是否存在）。如果路径不存在，即没有对应的数；
// - x_i = 1，那么对于第i位，n_i与字典树中的数的异或结果是什么都可以符合条件：
// -- 如果异或结果为1，就应该往下一位看。由于a^b=c可以得到a^c=b，因此通过n_i^1得到字典树中与n_i异或结果为1的路径；
// -- 如果异或结果为0，不需要再看后面的位，已经符合条件。那么应该得到所有前缀为当前路径的数有多少个。
func (t *trie) search(num, x int) int {
	cur := t
	sum := 0
	// 从高位往低位遍历
	for k := HIGH_BIT; k >= 0; k-- {
		bit := (num >> k) & 1    // 取第k位的值
		if ((x >> k) & 1) != 0 { // x的第k位为1
			// 对异或结果为0的处理
			// 两个相同的值，异或结果为0，应该看trie中bit的路径是否存在
			if cur.children[bit] != nil {
				// 路径存在，说明bit与trie中异或结果为0，此时异或结果<x，直接累加
				sum += cur.children[bit].cnt
			}
			// 对异或结果为1的处理
			// 由于a^b=c可以得到a^c=b，因此通过bit^1得到trie中与bit异或结果为1的路径
			// 如果路径不存在，不需要往下走了，直接返回当前的sum
			if cur.children[bit^1] == nil {
				return sum
			}
			// 路径存在，往下一位走
			cur = cur.children[bit^1]
		} else { // x的第k位为0
			// 对于第i位，bit与trie中的数的异或结果必须为0，继续判断下一位才有可能符合条件；
			// 如果路径不存在，不需要往下走了，直接返回当前的sum
			if cur.children[bit] == nil {
				return sum
			}
			// 路径存在，往下一位走
			cur = cur.children[bit]
		}

	}
	sum += cur.cnt // 遍历到最后，最后一个节点的cnt也要累加
	return sum
}

// @lc code=end
