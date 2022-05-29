package main

import "fmt"

func main() {

	println(fmt.Sprintf("mergeSort: {%+v}", mergeSort([]int{2, 1, 32, 4, 5, 5, 3, 4})))

}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left []int, right []int) []int {
	l := 0
	r := 0
	var result []int
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func quickSort(nums []int) []int {
	return nil
}

func quickSort1(nums []int, start, end int) []int {

	return nil
}
func quickSortSwap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
