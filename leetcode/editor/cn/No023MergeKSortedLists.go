package main

import "fmt"

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2016 👎 0

func main() {

	no23Print("%+v", mergeKLists([]*ListNode{genNode([]int{1, 4, 5}), genNode([]int{1, 3, 4}), genNode([]int{2, 6})}))
	no23Print("%+v", mergeKLists([]*ListNode{nil, nil, nil}))
}

func no23Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%d, next: %s", l.Val, l.Next)
}

func genNode(nums []int) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for _, val := range nums {
		temp := &ListNode{Val: val}
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = temp
	}
	return dummy.Next
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: -100000}
	for _, l := range lists {
		dummy = mergeKListsSubMerge(dummy, l)
	}
	return dummy.Next
}

func mergeKListsSubMerge(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -100000}
	head := dummy
	for list1 != nil && list2 != nil {
		v1, v2 := list1.Val, list2.Val
		if v1 < v2 {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	for list1 != nil {
		head.Next = list1
		list1 = list1.Next
		head = head.Next
	}
	for list2 != nil {
		head.Next = list2
		list2 = list2.Next
		head = head.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
