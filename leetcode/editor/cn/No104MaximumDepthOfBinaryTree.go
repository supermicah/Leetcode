package main

import "fmt"

//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1258 👎 0

func main() {
	//root := &TreeNode{Val: 3}
	//root.Left = &TreeNode{Val: 9}
	//root.Right = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}

	value := maxDepth(root)
	no104Print("%+v", value)
}

func no104Print(format string, params ...interface{}) {
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDP := maxDepth(root.Left)
	rightDP := maxDepth(root.Right)
	if leftDP < rightDP {
		return rightDP + 1
	} else {
		return leftDP + 1
	}

	//var (
	//	stack []*TreeNode
	//	count int
	//)
	//count++
	//if root.Left == nil && root.Right == nil {
	//	return count
	//}
	//stack = append(stack, root)
	//for len(stack) > 0 {
	//	l := len(stack)
	//	for i := 0; i < l; i++ {
	//		node := stack[0]
	//		stack = stack[1:]
	//		if node.Left != nil {
	//			if node.Left.Left != nil || node.Left.Right != nil {
	//				stack = append(stack, node.Left)
	//			}
	//		}
	//		if node.Right != nil {
	//			if node.Right.Left != nil || node.Right.Right != nil {
	//				stack = append(stack, node.Right)
	//			}
	//		}
	//	}
	//	count++
	//}
	//return count
}

//leetcode submit region end(Prohibit modification and deletion)
