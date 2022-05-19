package main

import "fmt"

//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
//
//
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
// Related Topics 数组 二分查找 矩阵 👍 646 👎 0

func main() {
	//value := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 11)
	//value := searchMatrix([][]int{{1}}, 0)
	value := searchMatrix([][]int{{1, 3}}, 3)
	no74Print("%+v", value)
}

func no74Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}

	start := 0
	totalLen := len(matrix)
	colLen := len(matrix[0])
	end := totalLen*colLen - 1
	for start+1 < end {
		mid := start + (end-start)/2
		midVal := matrix[mid/colLen][mid%colLen]
		if midVal == target {
			return true
		} else if midVal < target {
			start = mid
		} else if midVal > target {
			end = mid
		}
	}
	if matrix[start/colLen][start%colLen] == target ||
		matrix[end/colLen][end%colLen] == target {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
