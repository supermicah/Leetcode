package main

import "fmt"

//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 672 👎 0

func main() {
	value := 1
	fmt.Println(fmt.Sprintf("%+v", value))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	//win := make(map[byte]int)
	//need := make(map[byte]int)
	//for i := 0; i < len(s2); i++ {
	//	need[s2[i]]++
	//}
	//
	//left := 0
	//right := 0
	//
	//// 匹配结果
	//match := 0
	//start := 0
	//end := 0
	//
	//for right < len(s1) {
	//
	//}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
