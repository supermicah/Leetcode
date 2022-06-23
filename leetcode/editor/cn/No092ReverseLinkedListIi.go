package main

import "fmt"

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
// Related Topics 链表 👍 1274 👎 0

func main() {
	var head *ListNode

	//for _, val := range []int{3, 5} { //1,2,3,4,5
	for _, val := range []int{1, 2, 3, 4, 5} { //1,2,3,4,5
		temp := &ListNode{Val: val}
		if head == nil {
			head = &ListNode{Val: val}
			continue
		}
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = temp
	}
	fmt.Println(fmt.Sprintf("before: %+v", head))
	value := reverseBetween(head, 2, 4)
	fmt.Println(fmt.Sprintf("after: %+v", head))
	fmt.Println(fmt.Sprintf("%+v", value))
}

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func (l *ListNode) String() string {
//	return fmt.Sprintf("%d, next: %s", l.Val, l.Next)
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func (l *ListNode) String() string {
//	return fmt.Sprintf("%d, next: %s", l.Val, l.Next)
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	current := pre.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
