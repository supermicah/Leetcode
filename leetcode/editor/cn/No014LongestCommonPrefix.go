package main

import "fmt"

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
// Related Topics 字符串 👍 2305 👎 0

func main() {
	value := 1
	no14Print("%+v", value)
}

func no14Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 || len(s) == 0 {
			return ""
		}
		s = longestCommonPrefixIsSame(s, strs[i])
	}
	return s

}

func longestCommonPrefixIsSame(s, n string) string {
	i := 0
	for ; i < len(s) && i < len(n); i++ {
		if s[i] != n[i] {
			break
		}
	}
	return s[:i]
}

//leetcode submit region end(Prohibit modification and deletion)
