package main

import "fmt"

//给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
//。
//
//
//
// 示例 1：
//
//
//输入：left = 5, right = 7
//输出：4
//
//
// 示例 2：
//
//
//输入：left = 0, right = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：left = 1, right = 2147483647
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= left <= right <= 2³¹ - 1
//
// Related Topics 位运算 👍 378 👎 0

func main() {

	value := rangeBitwiseAnd(1, 5)
	fmt.Println(fmt.Sprintf("%+v", value))
}

//leetcode submit region begin(Prohibit modification and deletion)
func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left >>= 1
		right >>= 1
		count++
	}
	return right << count
}

//leetcode submit region end(Prohibit modification and deletion)
func rangeBitwiseAnd2(left int, right int) int {
	total := left
	for i := left + 1; i <= right && total != 0; i++ {
		total = total & i
	}
	return total
}
