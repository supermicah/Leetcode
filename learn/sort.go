package main

import "fmt"

func main() {

	println(fmt.Sprintf("mergeSort: \t{%+v}", mergeSort([]int{2, 1, 3, 32, 4, 5, 5, 3, 4})))
	println(fmt.Sprintf("quickSort: \t{%+v}", quickSort([]int{2, 1, 3, 32, 4, 5, 5, 3, 4})))
	println(fmt.Sprintf("bubbleSort: \t{%+v}", bubbleSort([]int{2, 1, 3, 32, 4, 5, 5, 3, 4})))
	println(fmt.Sprintf("selectionSort: \t{%+v}", selectionSort([]int{2, 1, 3, 32, 4, 5, 5, 3, 4})))

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
	return quickSort1(nums, 0, len(nums)-1)
}

func quickSort1(nums []int, start, end int) []int {
	if start >= end {
		return nums
	}
	pIndex := quickSortGetSpilt(nums, start, end)
	quickSort1(nums, start, pIndex-1)
	quickSort1(nums, pIndex+1, end)
	return nums
}

func quickSortGetSpilt(nums []int, start, end int) int {
	pIndex := start
	left := start
	right := end
	pValue := nums[pIndex]

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

func bubbleSort(nums []int) []int {
	// 需要循环的次数，因为最后一个没有对比项，所以只需要统计到最后一个的前面一项即可
	for i := 0; i < len(nums)-1; i++ {
		// 如果i不为0，则表示已经完成过几轮了，所以后面的数据已经是最大的了
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				minIndex = j
			}
		}

		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
