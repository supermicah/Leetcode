package main

import "fmt"

//给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//
//
//
//
// 示例 1:
//
//
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
//
// 示例 2:
//
//
//输入: numRows = 1
//输出: [[1]]
//
//
//
//
// 提示:
//
//
// 1 <= numRows <= 30
//
// Related Topics 数组 动态规划 👍 778 👎 0

func main() {
	no118Print("%+v", generate(10))
}

func no118Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	m1 := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		if i == 0 {
			m1[i] = row
			continue
		}
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = m1[i-1][j-1] + m1[i-1][j]
		}
		m1[i] = row
	}
	return m1
}

//leetcode submit region end(Prohibit modification and deletion)
