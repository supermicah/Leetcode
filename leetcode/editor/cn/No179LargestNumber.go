package main

import (
	"fmt"
	"sort"
	"strconv"
)

//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
//输入：nums = [10,2]
//输出："210"
//
// 示例 2：
//
//
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
// Related Topics 贪心 字符串 排序 👍 955 👎 0

func main() {
	value := 1
	no179Print("%+v", value)
}

func no179Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
func largestNumber(nums []int) string {

	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		// 把x的位数和y的最高位算出来，然后用x与y的最高位乘+y与另外一个比较
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})

	// 如果最大位是0， 则表示后面所有都是0，数字只能是0
	if nums[0] == 0 {
		return "0"
	}

	// 优化：刚开始用数组拼接，之后使用string构造方法一次性转为string
	var ans []byte
	for _, num := range nums {
		ans = append(ans, strconv.Itoa(num)...)
	}

	return string(ans)
}

func largestNumberSort(nums []int) []int {

	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
