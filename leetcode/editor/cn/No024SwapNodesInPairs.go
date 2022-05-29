package main

import "fmt"

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// Related Topics 递归 链表 👍 1395 👎 0

func main() {
	var head *ListNode
	for _, val := range []int{1, 2, 3, 4, 5, 6} { //1,2,3,4,5
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
	value := swapPairs(head)
	no24Print("%+v", value)
}

func no24Print(format string, params ...interface{}) {
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
func swapPairs(head *ListNode) *ListNode {
	t := &ListNode{
		Val:  0,
		Next: head,
	}
	current := t
	for current != nil && current.Next != nil && current.Next.Next != nil {
		// 0, 1, 2, 3
		temp := current.Next.Next.Next        // 3单独拎出来
		current.Next.Next.Next = current.Next // 1 挂在到2的next 0，1，2，1
		current.Next = current.Next.Next      // 2 挂在到0上 0，2，1
		current.Next.Next.Next = temp         // 3 挂在到1上
		current = current.Next.Next
	}

	return t.Next
}

//leetcode submit region end(Prohibit modification and deletion)
