package main

import "fmt"

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//
// 输入为 非空 字符串且只包含数字 1 和 0。
//
//
//
// 示例 1:
//
// 输入: a = "11", b = "1"
//输出: "100"
//
// 示例 2:
//
// 输入: a = "1010", b = "1011"
//输出: "10101"
//
//
//
// 提示：
//
//
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
//
// Related Topics 位运算 数学 字符串 模拟 👍 832 👎 0

func main() {
	value := 1
	no67Print("%+v", value)
}

func no67Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	max := la
	if la < lb {
		max = lb
	}
	ansArr := make([]byte, max)
	c := byte(0)
	for i := 0; i < max; i++ {
		var ba, bb byte
		if i < la {
			ba = a[la-1-i] - '0'
		}
		if i < lb {
			bb = b[lb-1-i] - '0'
		}
		m := c + ba + bb
		ansArr[max-1-i] = m%2 + '0'
		c = m / 2
	}
	if c != 0 {
		ansArr = append([]byte{c + '0'}, ansArr...)
	}
	return string(ansArr)
}

//leetcode submit region end(Prohibit modification and deletion)
