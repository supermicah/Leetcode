package main

import (
	"fmt"
	"sort"
	"strings"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4648 👎 0

func main() {
	value := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(fmt.Sprintf("%+v", value))
}

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	var rsp [][]int
	if len(nums) < 3 {
		return rsp
	}
	sort.Ints(nums)
	keyMap := make(map[string]bool)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				iValue := nums[i]
				jValue := nums[j]
				kValue := nums[k]
				if iValue+jValue+kValue == 0 {
					key := sortKey(iValue, jValue, kValue)
					if _, ok := keyMap[key]; ok {
						continue
					}
					rsp = append(rsp, []int{iValue, jValue, kValue})
					keyMap[key] = true
				}
			}
		}

	}

	return rsp
}

func sortKey(key1, key2, key3 int) string {
	k1 := fmt.Sprintf("%d", key1)
	k2 := fmt.Sprintf("%d", key2)
	k3 := fmt.Sprintf("%d", key3)
	s := []string{k1, k2, k3}
	sort.Strings(s)
	return strings.Join(s, "_")
}

//leetcode submit region end(Prohibit modification and deletion)
