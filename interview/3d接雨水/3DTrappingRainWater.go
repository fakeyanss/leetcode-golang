package main

import (
	"container/heap"
	"fmt"
)

// 题目定义：
// 给定一个 m x n 的非负整数矩阵 heightMap，表示一个三维地形的高度（矩阵的每个元素 heightMap[i][j] 代表坐标 (i,j) 处的高度）。下雨后，雨水会被困在地形的低洼处，计算总共能接住多少单位的雨水。
// 关键说明：
// 1. 雨水会被困在地形的低洼处，且无法流出 heightMap 的边界（最外层的所有单元格视为 “边缘”，无法储存雨水）。
// 2. 任意单元格 (i, j) 能储存的最大雨水量，由其周围 “边界” 的最低高度减去该单元格自身的高度决定。若计算结果为负数，则该单元格储存 0 单位雨水（即无积水）。
// 3. 地形视为实心结构，雨水只能沿着上下左右四个正方向在单元格之间流动，无法穿透地形本身。
// 示例：
// 输入：
// plaintext
// heightMap = [
//   [1,4,3,1,3,2],
//   [3,2,1,3,2,4],
//   [2,3,3,2,3,1]
// ]
// 输出：4解释：下雨后共能接住 4 单位雨水，低洼处的积水分布符合 “周围最低边界决定储水量” 的规则。

// 解法
// 核心思路：最小堆（优先队列）+ 广度优先搜索，从地形外围边界开始，动态维护当前最低边界，逐步向内扩展计算每个位置的储水量。

// 定义优先队列的元素结构：存储高度、行、列
type Cell struct {
	height int
	row    int
	col    int
}

type MinHeap []Cell

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].height < h[j].height
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 核心函数：计算三维接雨水总量
func trapRainWater(heightMap [][]int) int {
	rows, cols := len(heightMap), len(heightMap[0])
	// 行列长度<=2时无法接雨水
	if rows <= 2 || cols <= 2 {
		return 0
	}

	// 标记访问过的位置
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 初始化最小堆，将上下左右边界加入
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
				heap.Push(h, Cell{height: heightMap[i][j], row: i, col: j})
				visited[i][j] = true
			}
		}
	}

	// 上下左右4个方向
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// 总接水量
	totalWater := 0
	// 从外向内遍历，每次取最矮的位置
	for h.Len() > 0 {
		cur := heap.Pop(h).(Cell)
		for _, dir := range directions {
			nextRow := cur.row + dir[0]
			nextCol := cur.col + dir[1]
			// 检查新的位置是否合法
			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && !visited[nextRow][nextCol] {
				visited[nextRow][nextCol] = true
				// 计算新的位置可以接的雨水量，新位置更高则接水量为0
				water := max(0, cur.height-heightMap[nextRow][nextCol])
				totalWater += water
				// 新位置作为边界，加入时的高度取当前边界高度和新位置高度的max
				nextHeight := max(cur.height, heightMap[nextRow][nextCol])
				heap.Push(h, Cell{height: nextHeight, row: nextRow, col: nextCol})
			}
		}
	}

	return totalWater
}

func main() {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println("总储水量：", trapRainWater(heightMap)) // 输出：4
}
