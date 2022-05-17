package main

import "fmt"

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//
//
// 示例 3:
//
//
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums 为 无重复元素 的 升序 排列数组
// -10⁴ <= target <= 10⁴
//
// Related Topics 数组 二分查找 👍 1547 👎 0

func main() {
	//value := searchInsert([]int{1, 3, 5, 6}, 5)
	//value := searchInsert([]int{1, 3, 5, 6}, 7)
	value := searchInsert([]int{1, 3, 5, 6}, 0)
	no35Print("%+v", value)
}

func no35Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		midVal := nums[mid]
		if midVal == target {
			return mid
		} else if midVal < target {
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}

	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		return end + 1
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
