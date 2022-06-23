package main

import "fmt"

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,3,2]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,1,0,1,100]
//输出：100
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//
//
// 注意：本题与主站 137 题相同：https://leetcode-cn.com/problems/single-number-ii/
// Related Topics 位运算 数组 👍 81 👎 0

func main() {
	noOfferII004Print("%+v", singleNumber([]int{2, 2, 3, 2}))
}

func noOfferII004Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	arr := make([]int, 64)
	ans := 0
	for i := 0; i < 64; i++ {
		total := 0
		for _, num := range nums {
			total += num >> i & 1
		}
		arr[i] = total % 3
	}

	for i := 0; i < 64; i++ {
		ans += arr[i] << i
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
