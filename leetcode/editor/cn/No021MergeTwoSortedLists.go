package main

import "fmt"

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// Related Topics 递归 链表 👍 2474 👎 0

func main() {
	value := 1
	no21Print("%+v", value)
}

func no21Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	n1 := list1
	n2 := list2
	cValue := -200
	dummy := &ListNode{Val: cValue}
	head := dummy
	for n1 != nil && n2 != nil {
		v1 := n1.Val
		v2 := n2.Val
		if v1 < v2 {
			head.Next = n1
			n1 = n1.Next
		} else {
			head.Next = n2
			n2 = n2.Next
		}
		head = head.Next
	}

	for n1 != nil {
		head.Next = n1
		n1 = n1.Next
		head = head.Next
	}

	for n2 != nil {
		head.Next = n2
		n2 = n2.Next
		head = head.Next
	}

	return dummy.Next

}

//leetcode submit region end(Prohibit modification and deletion)
