package main

import (
	"fmt"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
// Related Topics 哈希表 字符串 滑动窗口 👍 7626 👎 0

func main() {
	value := lengthOfLongestSubstring("pwwkew")
	no3Print("%+v", value)
}

func no3Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	var (
		m      = make(map[byte]int)
		result int
		right  int
	)

	for left := 0; left < length; left++ {
		if left != 0 {
			delete(m, s[left-1])
		}

		for right < length && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}
		if right-left > result {
			result = right - left
		}
	}

	//var (
	//	eMap   = make(map[byte]int)
	//	result int
	//	right  = -1
	//)

	//for left := 0; left < length; left++ {
	//	if left != 0 {
	//		delete(eMap, s[left-1])
	//	}
	//	for right+1 < length && eMap[s[right+1]] == 0 {
	//		eMap[s[right+1]]++
	//		right++
	//	}
	//	// 取较大值
	//	if right-left+1 > result {
	//		result = right - left + 1
	//	}
	//}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
