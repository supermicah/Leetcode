package main

import "fmt"

//在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：
//
//
// 值 0 代表空单元格；
// 值 1 代表新鲜橘子；
// 值 2 代表腐烂的橘子。
//
//
// 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
//
// 返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
//输出：4
//
//
// 示例 2：
//
//
//输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
//输出：-1
//解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。
//
//
// 示例 3：
//
//
//输入：grid = [[0,2]]
//输出：0
//解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] 仅为 0、1 或 2
//
// Related Topics 广度优先搜索 数组 矩阵 👍 560 👎 0

func main() {
	value := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	no994Print("%+v", value)
}

func no994Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {

	var (
		goodCount int
		badQueue  [][2]int
		step      int
	)

	gridX := len(grid)
	gridY := len(grid[0])

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				goodCount++
			} else if grid[i][j] == 2 {
				badQueue = append(badQueue, [2]int{i, j})
			}
		}
	}

	for len(badQueue) > 0 {
		l := len(badQueue)
		for i := 0; i < l; i++ {
			bad := badQueue[0]      // 取出第一个
			badQueue = badQueue[1:] //删除第一个
			for _, idxArray := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				x := bad[0] + idxArray[0]
				y := bad[1] + idxArray[1]
				if x < 0 || x > gridX-1 || y < 0 || y > gridY-1 {
					continue
				}
				if grid[x][y] == 1 {
					goodCount--
					grid[x][y] = 2
					badQueue = append(badQueue, [2]int{x, y})
				}
			}
		}
		if len(badQueue) > 0 {
			step++
		}
	}
	if goodCount == 0 {
		return step
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
