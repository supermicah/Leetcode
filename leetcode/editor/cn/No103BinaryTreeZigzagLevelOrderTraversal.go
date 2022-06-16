package main

import "fmt"

//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 广度优先搜索 二叉树 👍 641 👎 0

func main() {
	no103Print("%+v", zigzagLevelOrderReverse2([]int{1, 2}))
	no103Print("%+v", zigzagLevelOrderReverse2([]int{1, 2, 3}))
	no103Print("%+v", zigzagLevelOrderReverse2([]int{1, 2, 3, 4}))
	no103Print("%+v", zigzagLevelOrderReverse2([]int{1, 2, 3, 4, 5, 6}))
	no103Print("%+v", zigzagLevelOrderReverse2([]int{1, 2, 3, 4, 5, 6, 7}))
}

func no103Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func (t *TreeNode) String() string {
//	return fmt.Sprintf("[%d, left: %s, right: %s]", t.Val, t.Left, t.Right)
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result      [][]int
		stack       []*TreeNode
		needReverse bool
	)

	stack = append(stack, root)
	for len(stack) > 0 {
		l := len(stack)
		var subAns []int
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			subAns = append(subAns, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		if len(subAns) > 0 {
			if needReverse {
				zigzagLevelOrderReverse(subAns)
			}
			result = append(result, subAns)
		}
		needReverse = !needReverse
	}

	return result
}

func zigzagLevelOrderReverse(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for start, end := 0, len(nums)-1; start < end; start, end = start+1, end-1 {
		nums[start], nums[end] = nums[end], nums[start]
	}
}

func zigzagLevelOrderReverse2(subAns []int) []int {
	left := 0
	right := len(subAns) - 1
	for left < right {
		subAns[left], subAns[right] = subAns[right], subAns[left]
		left++
		right--
	}
	return subAns
}

//leetcode submit region end(Prohibit modification and deletion)
