package main

import "fmt"

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
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
// Related Topics 链表 双指针 👍 886 👎 0

func main() {
	head := &ListNode{
		Val: 1,
	}

	//for _, val := range []int{2, 3, 3, 4, 4, 5} {//
	for _, val := range []int{1, 1, 2, 3} { // 1,1,1,2,3
		temp := &ListNode{Val: val}
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = temp
	}

	fmt.Println(fmt.Sprintf("test case: %+v", head))

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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	full := &ListNode{Val: 0}
	full.Next = head
	head = full

	rmVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			// 保留需要删除的val
			rmVal = head.Next.Val
			// 下个的值与需要删除的只相同，head的下一个，就是下下个
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			// 如果下一个和下下个不一样，则下一个可以返回
			head = head.Next
		}
	}

	return full.Next
}

//leetcode submit region end(Prohibit modification and deletion)
