package main

import "fmt"

//给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
//
//
//
// 示例 1：
//
//
//输入：s = "aaabb", k = 3
//输出：3
//解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
//
//
// 示例 2：
//
//
//输入：s = "ababbc", k = 2
//输出：5
//解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文字母组成
// 1 <= k <= 10⁵
//
// Related Topics 哈希表 字符串 分治 滑动窗口 👍 693 👎 0

func main() {
	value := longestSubstring("ababbc", 2)
	no395Print("%+v", value)
}

func no395Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func longestSubstring(s string, k int) (ans int) {
	for charNum := 1; charNum <= 26; charNum++ {
		cnt := [26]int{}
		total := 0
		lessK := 0
		l := 0

		for r, ch := range s {
			ch = ch - 'a'
			if cnt[ch] == 0 {
				total++
				lessK++
			}
			cnt[ch]++
			if cnt[ch] == k {
				lessK--
			}

			for total > charNum {
				lch := s[l] - 'a'
				if cnt[lch] == k {
					lessK++
				}
				cnt[lch]--
				if cnt[lch] == 0 {
					lessK--
					total--
				}
				l++
			}
			if lessK == 0 {
				ans = longestSubstringMax(ans, r-l+1)
			}
		}
	}

	return ans
	// 从一个字符-> 到26个字符
	//for t := 1; t <= 26; t++ {
	//	cnt := [26]int{}
	//	total := 0
	//	lessK := 0
	//	l := 0
	//	for r, ch := range s {
	//		ch -= 'a'
	//		if cnt[ch] == 0 {
	//			total++
	//			lessK++
	//		}
	//		cnt[ch]++
	//		if cnt[ch] == k {
	//			lessK--
	//		}
	//
	//		for total > t {
	//			ch := s[l] - 'a'
	//			if cnt[ch] == k {
	//				lessK++
	//			}
	//			cnt[ch]--
	//			if cnt[ch] == 0 {
	//				total--
	//				lessK--
	//			}
	//			l++
	//		}
	//		if lessK == 0 {
	//			ans = max(ans, r-l+1)
	//		}
	//	}
	//}
	//return ans
}

func longestSubstringMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
