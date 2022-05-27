package main

import (
	"fmt"
	"sort"
)

//给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[
//j]) 都应当满足：
//
// answer[i] % answer[j] == 0 ，或
// answer[j] % answer[i] == 0
//
//
// 如果存在多个有效解子集，返回其中任何一个均可。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[1,2]
//解释：[1,3] 也会被视为正确答案。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,4,8]
//输出：[1,2,4,8]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2 * 10⁹
// nums 中的所有整数 互不相同
//
// Related Topics 数组 数学 动态规划 排序 👍 442 👎 0

func main() {
	value := largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720})
	//value := largestDivisibleSubset([]int{1, 2, 3, 4, 6, 24})
	//value := largestDivisibleSubset([]int{1, 2, 3}) //期望结果:[1,2]
	fmt.Println(fmt.Sprintf("%+v", value))
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	dp := make([]int, l)
	for i := range dp {
		// 因为所有数都能被本身整除，所以这里默认使用1
		dp[i] = 1
	}

	// 如果i可以被j整除，那么i的深度应该在j+1的程度
	maxSize, maxVal := 1, 1
	for i := 1; i < l; i++ {
		for j, v := range nums[:i] {
			//println(fmt.Sprintf("i: %d, j: %d, dp[i]: %d, dp[j]: %d, maxSize: %d, maxVal: %d", i, j, dp[i], dp[j], maxSize, maxVal))
			if nums[i]%v == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	// 如果最大只有一层，那么返回第一个数就可以了
	if maxSize == 1 {
		return []int{nums[0]}
	}

	//
	var res []int
	for i := l - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}

	//println(fmt.Sprintf("maxSize: %d, maxVal: %d, res: %+v", maxSize, maxVal, res))
	sort.Ints(res)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
