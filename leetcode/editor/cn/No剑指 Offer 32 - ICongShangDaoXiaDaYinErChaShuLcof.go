package main

import "fmt"

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
//
//
// 例如:
//给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回：
//
// [3,9,20,15,7]
//
//
//
//
// 提示：
//
//
// 节点总数 <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 202 👎 0

func main() {
	value := 1
	noOffer32IPrint("%+v", value)
}

func noOffer32IPrint(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		stack []*TreeNode
		ans   []int
	)
	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack)
		for i := 0; i < n; i++ {
			node := stack[0]
			stack = stack[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
