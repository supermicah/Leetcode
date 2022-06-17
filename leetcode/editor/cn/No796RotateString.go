package main

import (
	"fmt"
	"strings"
)

//给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。
//
// s 的 旋转操作 就是将 s 最左边的字符移动到最右边。
//
//
// 例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。
//
//
//
//
// 示例 1:
//
//
//输入: s = "abcde", goal = "cdeab"
//输出: true
//
//
// 示例 2:
//
//
//输入: s = "abcde", goal = "abced"
//输出: false
//
//
//
//
// 提示:
//
//
// 1 <= s.length, goal.length <= 100
// s 和 goal 由小写英文字母组成
//
// Related Topics 字符串 字符串匹配 👍 257 👎 0

func main() {
	value := rotateString("bbbacddceeb", "ceebbbbacdd")
	no796Print("%+v", value)
}

func no796Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(s+s, goal)
	//
	//	l := len(goal)
	//	if l == 0 {
	//		return true
	//	}
	//	if s == goal {
	//		return true
	//	}
	//next:
	//	for i := 0; i < l; i++ {
	//		for j := 0; j < l; j++ {
	//			if s[(i+j)%l] != goal[j] {
	//				continue next
	//			}
	//		}
	//		return true
	//	}
	//	return false
}

//leetcode submit region end(Prohibit modification and deletion)
