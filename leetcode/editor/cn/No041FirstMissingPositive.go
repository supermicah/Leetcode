package main

import "fmt"

//给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,0]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [3,4,-1,1]
//输出：2
//
//
// 示例 3：
//
//
//输入：nums = [7,8,9,11,12]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// Related Topics 数组 哈希表 👍 1497 👎 0

func main() {
	no41Print("%+v", firstMissingPositive([]int{3, 4, -1, 1}))
	no41Print("%+v", firstMissingPositive([]int{1, 2, 0}))
	no41Print("%+v", firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func no41Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1

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
	//n := len(nums)
	//// 如果把负数改为n+1
	//for i := 0; i < n; i++ {
	//	if nums[i] <= 0 {
	//		nums[i] = n + 1
	//	}
	//}
	////
	//for i := 0; i < n; i++ {
	//	num := abs(nums[i])
	//	if num <= n {
	//		fmt.Println(num - 1)
	//		nums[num-1] = -abs(nums[num-1])
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	if nums[i] > 0 {
	//		return i + 1
	//	}
	//}
	//return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
