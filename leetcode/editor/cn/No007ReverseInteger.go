package main

import (
	"fmt"
	"math"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−2³¹, 231 − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
//
//
// 示例 1：
//
//
//输入：x = 123
//输出：321
//
//
// 示例 2：
//
//
//输入：x = -123
//输出：-321
//
//
// 示例 3：
//
//
//输入：x = 120
//输出：21
//
//
// 示例 4：
//
//
//输入：x = 0
//输出：0
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
// Related Topics 数学 👍 3528 👎 0

func main() {
	no7Print("%+v", reverse(-123))
}

func no7Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	ans := 0
	for x != 0 {
		if x < math.MinInt32/10 || x > math.MaxInt32/10 {
			return 0
		}
		d := x % 10
		x = x / 10
		ans = ans*10 + d
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
