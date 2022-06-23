package main

import "fmt"

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
// Related Topics 位运算 数组 回溯 👍 1686 👎 0

func main() {
	set := []int{1, 3, 4}
	value := append([]int(nil), set...)
	no78Print("%+v", value)
	no78Print("%+v", set)
}

func no78Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) (ans [][]int) {
	var set []int
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(nums)-1]
		dfs(cur + 1)
	}
	dfs(0)

	//set := []int{}
	//var dfs func(int)
	//dfs = func(cur int) {
	//	if cur == len(nums) {
	//		ans = append(ans, append([]int(nil), set...))
	//		return
	//	}
	//	set = append(set, nums[cur])
	//	dfs(cur + 1)
	//	set = set[:len(set)-1]
	//	dfs(cur + 1)
	//}
	//dfs(0)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
