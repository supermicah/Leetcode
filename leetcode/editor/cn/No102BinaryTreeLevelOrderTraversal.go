package main

import "fmt"

//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
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
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1340 👎 0

func main() {
	value := 1
	no102Print("%+v", value)
}

func no102Print(format string, params ...interface{}) {
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
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		result [][]int
		stack  []*TreeNode
	)

	stack = append(stack, root)

	for len(stack) > 0 {
		l := len(stack)
		var ans []int
		for i := 0; i < l; i++ {
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
		if len(ans) > 0 {
			result = append(result, ans)
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
