package main

// main 字节面试题，数组A中给定可以使用的1~9的数，返回由A数组中的元素组成的小于n的最大数。
// 例如A={1, 2, 4, 9}，x=2533，返回2499
func main() {

	println(getMax([]int{1, 2, 4, 9}, 2543)) // 2499
	println(getMax([]int{1, 2, 4, 9}, 2000)) // 1999
	println(getMax([]int{1, 2, 4, 9}, 1001)) // 999
}

func getMax(nums []int, target int) int {

	low := make(map[int]int)
	n := len(nums)

	for i, j, lo := 0, 0, -1; i < 10; i++ {
		if j < n && nums[j] <= i {
			lo = nums[j]
			j++
		}
		low[i] = lo
	}

	r := 0
	k := target - 1
	max := 0
	c := 0
	ans := 0
	multi := 1
	for k > 0 {
		r = k % 10
		k = k / 10
		max = max*10 + nums[n-1]

		if v := low[r+c]; v == -1 || v == 0 {
			ans = max
			c = -1
		} else if v == r+c {
			ans = v*multi + ans
			c = 0
		} else {
			ans = v*multi + max/10 // TODO 这里要记住。。重点
			c = 0
		}
		multi = multi * 10
	}

	if c == -1 {
		ans = ans % (multi / 10)
	}

	return ans
}

//func getMax(nums []int, target int) int {
//	k := target - 1
//	var (
//		// n nums 的长度；ans 表示结果；c 表示进位；multi 表示倍数
//		// low表示 小于i的最大数字
//		n     = len(nums)
//		low   = make(map[int]int)
//		multi = 1
//		c     = 0
//		ans   = 0
//	)
//
//	// 整理出来low的map
//	for i, j, lo := 0, 0, -1; i < 10; i++ {
//		if j <= n-1 && nums[j] <= i {
//			lo = nums[j]
//			j = j + 1
//		}
//		low[i] = lo
//	}
//
//	// r表示当前的余数；max表示当前可遍历的最大数；multi表示当前位数
//	r := 0
//	max := 0
//	for k > 0 {
//		r = k % 10
//		k = k / 10
//		max = max*10 + nums[n-1]
//
//		if v := low[r+c]; v == -1 || v == 0 { // 当前位数不在nums中，且无更小的，则取最大数值，进位为-1
//			ans = max
//			c = -1
//		} else if v == r+c { // 当前值相同
//			ans = v*multi + ans
//			c = 0
//		} else { // 当前值没有，则取最大值
//			ans = v*multi + max/10
//			c = 0
//		}
//		multi = multi * 10
//	}
//
//	if c == -1 {
//		ans = ans % (multi / 10)
//	}
//
//	return ans
//}
