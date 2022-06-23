package main

import (
	"fmt"
	"strconv"
)

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
// 注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
//
//
//
// 示例 1:
//
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//
// 示例 2:
//
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 200
// num1 和 num2 只能由数字组成。
// num1 和 num2 都不包含任何前导零，除了数字0本身。
//
// Related Topics 数学 字符串 模拟 👍 959 👎 0

func main() {
	no43Print("%+v", multiply("99", "111"))
	no43Print("%+v", multiply("123", "456"))
	no43Print("%+v", multiply("10000", "10000"))
}

func no43Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//	ans := "0"
//	m, n := len(num1), len(num2)
//	for i := n - 1; i >= 0; i-- {
//		curr := ""
//		add := 0
//		for j := n - 1; j > i; j-- {
//			curr += "0"
//		}
//		y := int(num2[i] - '0')
//		for j := m - 1; j >= 0; j-- {
//			x := int(num1[j] - '0')
//			product := x*y + add
//			curr = strconv.Itoa(product%10) + curr
//			add = product / 10
//		}
//		for ; add != 0; add /= 10 {
//			curr = strconv.Itoa(add%10) + curr
//		}
//		ans = addStrings(ans, curr)
//	}
//	return ans
//}
//
//func addStrings(num1, num2 string) string {
//	i, j := len(num1)-1, len(num2)-1
//	add := 0
//	ans := ""
//	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
//		x, y := 0, 0
//		if i >= 0 {
//			x = int(num1[i] - '0')
//		}
//		if j >= 0 {
//			y = int(num2[j] - '0')
//		}
//		result := x + y + add
//		ans = strconv.Itoa(result%10) + ans
//		add = result / 10
//	}
//	return ans
//}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var (
		ans string
	)
	for i := 0; i < len(num2); i++ {
		n := int(num2[len(num2)-1-i] - '0')
		if n == 0 {
			continue
		}
		multi := ""
		for j := 0; j < i; j++ {
			multi += "0"
		}

		s := multiString(num1, n)
		ans = addString(ans, s+multi)

	}
	return ans
}

func addString(num1 string, num2 string) string {
	var (
		l1, l2 = len(num1) - 1, len(num2) - 1
		c      int
		ans    string
	)

	for ; l1 >= 0 || l2 >= 0 || c > 0; l1, l2 = l1-1, l2-1 {
		x, y := 0, 0
		if l1 >= 0 {
			x = int(num1[l1] - '0')
		}
		if l2 >= 0 {
			y = int(num2[l2] - '0')
		}
		m := x + y + c
		c = m / 10
		ans = strconv.Itoa(m%10) + ans
	}

	return ans
}

func multiString(num1 string, num int) string {
	var (
		ans string
		c   int
	)

	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		m := n1*num + c
		c = m / 10
		ans = strconv.Itoa(m%10) + ans
	}
	if c > 0 {
		ans = strconv.Itoa(c) + ans
	}

	return ans
}

//
//
//
//
//
//
//
//
//
//

//
//
//
//
//

//
//
//
//
//

//
//
//
//
//

//func multiply1(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//
//	var ans string
//
//	for i := 0; i < len(num2); i++ {
//		multi := ""
//		for j := 0; j < i; j++ {
//			multi += "0"
//		}
//		n2 := num2[len(num2)-1-i] - '0'
//		ans = addString(ans, multiString(num1, int(n2))+multi)
//	}
//
//	return ans
//}
//
//func multi1(num1 string, num2 int) string {
//	carry := 0
//	ans := ""
//	for i := 0; i < len(num1); i++ {
//		n1 := num1[len(num1)-1-i] - '0'
//		m := int(n1)*num2 + carry
//		if m > 9 {
//			carry = m / 10
//		} else {
//			carry = 0
//		}
//		m = m % 10
//		ans = strconv.Itoa(m) + ans
//	}
//	if carry > 0 {
//		ans = strconv.Itoa(carry) + ans
//	}
//
//	return ans
//}
//
//func add(num1 string, num2 string) string {
//	carry := uint8(0)
//	ans := ""
//	for i := 0; i < len(num1) || i < len(num2); i++ {
//
//		n1 := uint8(0)
//		if i < len(num1) {
//			n1 = num1[len(num1)-1-i] - '0'
//		}
//		n2 := uint8(0)
//		if i < len(num2) {
//			n2 = num2[len(num2)-1-i] - '0'
//		}
//
//		m := n1 + n2 + carry
//		if m > 9 {
//			carry = uint8(1)
//		} else {
//			carry = 0
//		}
//		m = m % 10
//		ans = strconv.Itoa(int(m)) + ans
//	}
//
//	if carry > 0 {
//		ans = "1" + ans
//	}
//	return ans
//}
//
//func addString(num1 string, num2 string) string {
//	var (
//		ans string
//		c   int
//	)
//
//	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || c != 0; i, j = i-1, j-1 {
//		x, y := 0, 0
//		if i >= 0 {
//			x = int(num1[i] - '0')
//		}
//		if j >= 0 {
//			y = int(num2[j] - '0')
//		}
//
//		num := x + y + c
//		c = num / 10
//		ans = strconv.Itoa(num%10) + ans
//	}
//	return ans
//}
//
//func multiString(num1 string, num2 int) string {
//	var (
//		ans string
//		c   int
//	)
//
//	for i := len(num1) - 1; i >= 0 || c != 0; i-- {
//		x := 0
//		if i >= 0 {
//			x = int(num1[i] - '0')
//		}
//
//		num := x*num2 + c
//		c = num / 10
//		ans = strconv.Itoa(num%10) + ans
//	}
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)
