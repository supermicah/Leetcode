package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", twoSum([]int{3, 2, 4}, 6)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return recursion(l1, l2, 0)
}

func recursion(l1, l2 *ListNode, overflow int) *ListNode {
	if l1 == nil && l2 == nil {
		if overflow > 0 {
			return &ListNode{Val: overflow}
		}
		return nil
	}
	var v1, v2, o int
	var ln1, ln2 *ListNode
	if l1 != nil {
		v1 = l1.Val
		ln1 = l1.Next
	}
	if l2 != nil {
		v2 = l2.Val
		ln2 = l2.Next
	}
	if v1+v2+overflow > 9 {
		o = 1
	}
	next := recursion(ln1, ln2, o)
	val := (v1 + v2 + overflow) % 10
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		numMap[num] = i
	}

	for i := 0; i < len(nums); i++ {
		want := target - nums[i]
		if idx, ok := numMap[want]; ok && i != idx {
			return []int{i, idx}
		}
	}
	//for i := 0; i< len(nums); i++{
	//	for j := i+1; j<len(nums); j++{
	//		if nums[i]+nums[j] == target{
	//			return []int{i,j}
	//		}
	//	}
	//}
	return nil
}
