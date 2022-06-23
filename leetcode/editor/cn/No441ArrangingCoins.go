package main

import (
	"fmt"
)

//你总共有 n 枚硬币，并计划将它们按阶梯状排列。对于一个由 k 行组成的阶梯，其第 i 行必须正好有 i 枚硬币。阶梯的最后一行 可能 是不完整的。
//
// 给你一个数字 n ，计算并返回可形成 完整阶梯行 的总行数。
//
//
//
// 示例 1：
//
//
//输入：n = 5
//输出：2
//解释：因为第三行不完整，所以返回 2 。
//
//
// 示例 2：
//
//
//输入：n = 8
//输出：3
//解释：因为第四行不完整，所以返回 3 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 数学 二分查找 👍 223 👎 0

func main() {
	value := 1
	no441Print("%+v", value)
}

func no441Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum == n {
			return i
		} else if sum > n {
			return i - 1
		}
	}
	return 0

	//return sort.Search(n, func(x int) bool {
	//	x++
	//	return x*(x+1) > 2*n
	//})

}

//leetcode submit region end(Prohibit modification and deletion)
