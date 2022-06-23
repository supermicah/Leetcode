package main

import "fmt"

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
//
//
//
// 进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,3,2,5]
//输出：[3,5]
//解释：[5, 3] 也是有效的答案。
//
//
// 示例 2：
//
//
//输入：nums = [-1,0]
//输出：[-1,0]
//
//
// 示例 3：
//
//
//输入：nums = [0,1]
//输出：[1,0]
//
//
// 提示：
//
//
// 2 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// 除两个只出现一次的整数外，nums 中的其他数字都出现两次
//
// Related Topics 位运算 数组 👍 609 👎 0

func main() {
	value := singleNumber([]int{1, 2, 1, 3, 2, 5})
	//value := singleNumber([]int{0, 0, 1, 2})
	fmt.Println(fmt.Sprintf("%+v", value))
}

//leetcode submit region begin(Prohibit modification and deletion)
func singleNumber3(nums []int) []int {
	length := len(nums)
	if length == 2 {
		return nums
	}
	total := 0
	for _, num := range nums {
		total = total ^ num
	}

	lastOne := 0
	for i := 0; i < 32; i++ {
		if (total>>i)&1 == 1 {
			lastOne = 1 << i
			break
		}
	}
	var one, two int
	for _, num := range nums {
		if num&lastOne == lastOne {
			one ^= num
		} else {
			two ^= num
		}
	}
	return []int{one, two}

}

//leetcode submit region end(Prohibit modification and deletion)
