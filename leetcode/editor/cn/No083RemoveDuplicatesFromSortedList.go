package main

import "fmt"

//给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,1,2]
//输出：[1,2]
//
//
// 示例 2：
//
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围 [0, 300] 内
// -100 <= Node.val <= 100
// 题目数据保证链表已经按升序 排列
//
// Related Topics 链表 👍 787 👎 0

func main() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	value := deleteDuplicates(head)
	fmt.Println(fmt.Sprintf("%+v", value))
}

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

func deleteDuplicates1(head *ListNode) *ListNode {
	// 获取队列句柄
	current := head
	// 如果当前和下一个不存在，则直接返回（只有一个，是没有重复的）
	for current != nil && current.Next != nil {
		// 循环当前值与下一个值，如果当前值和下一个相同，则跳过下一个，采用下下个
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		// 切换句柄
		current = current.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
