package main

import (
	"fmt"
	"sort"
)

func main() {
	noSearchPrint("%+v", searchRange([]int{1, 2, 3, 3, 3, 3, 3, 5, 6, 7}, 3))
}
func searchRange(A []int, target int) []int {
	firstIndex := sort.SearchInts(A, target)
	lastIndex := sort.SearchInts(A, target+1)
	return []int{firstIndex, lastIndex}
}

func noSearchPrint(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}
