package main

import "fmt"

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//
// Related Topics 递归 链表 👍 2510 👎 0

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
	fmt.Println(fmt.Sprintf("before: %+v", head))
	value := reverseList(head)
	fmt.Println(fmt.Sprintf("after: %+v", value))
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
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	cursor := head
	for cursor != nil {
		temp := cursor.Next
		cursor.Next = pre
		pre = cursor
		cursor = temp
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode

	for head != nil {
		// 下一个保存临时变量中
		temp := head.Next
		// 之前保存的，放到next上
		head.Next = pre
		// 把这个当成已经转换过的数据
		pre = head
		// 把head转移到下一个
		head = temp
	}

	return pre
}
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var listSlice []*ListNode
	listSlice = append(listSlice, head)
	for head.Next != nil {
		listSlice = append(listSlice, head.Next)
		head = head.Next
	}
	result := &ListNode{Val: 0}
	current := result

	for i := len(listSlice) - 1; i >= 0; i-- {
		node := listSlice[i]
		node.Next = nil
		current.Next = node
		current = current.Next
		fmt.Println(fmt.Sprintf("currentValue: %+v", current.Val))
	}

	return result.Next
}
