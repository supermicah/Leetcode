package main

import (
	"fmt"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 1697 👎 0

func main() {
	//value := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	value := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	no215Print("%+v", value)
}

func no215Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, start, end, k int) int {
	if start > end {
		return -1
	}
	p := quickSort0(nums, start, end)
	if p == len(nums)-k {
		return nums[p]
	} else if p < len(nums)-k {
		quickSort(nums, p+1, end, k)
	} else {
		quickSort(nums, start, p-1, k)
	}

	return nums[len(nums)-k]
}

func quickSort0(nums []int, start, end int) int {
	pIndex := start
	pValue := nums[pIndex]
	left := start
	right := end
	for left < right {
		for left < right && pValue <= nums[right] {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			pIndex = right
		}

		for left < right && nums[left] <= pValue {
			left++
		}

		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			pIndex = left
		}
	}
	return pIndex
}

//leetcode submit region end(Prohibit modification and deletion)
