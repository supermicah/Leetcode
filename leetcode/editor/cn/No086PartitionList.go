package main

import "fmt"

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
//
// Related Topics 链表 双指针 👍 562 👎 0

func main() {
	var head *ListNode
	for _, val := range []int{1, 4, 3, 2, 5, 2} { //1,2,3,4,5
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
	value := partition(head, 3)
	no86Print("%+v", value)
}

func no86Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
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

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	head = dummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
			continue
		}

		tail.Next = head.Next
		tail = tail.Next
		head.Next = head.Next.Next
	}
	tail.Next = nil
	head.Next = tailDummy.Next

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
