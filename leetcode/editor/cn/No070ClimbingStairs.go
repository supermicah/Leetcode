package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
// 示例 2：
//
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//
//
// 提示：
//
//
// 1 <= n <= 45
//
// Related Topics 记忆化搜索 数学 动态规划 👍 2479 👎 0

func main() {
	no70Print("%+v", climbStairs(4))
	no70Print("%+v", climbStairs(5))
}

func no70Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
var m = make(map[int]int)

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	p, q, r := 1, 2, 3
	for i := 4; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r

	return climbStairs(n-1) + climbStairs(n-2)

}

//leetcode submit region end(Prohibit modification and deletion)
