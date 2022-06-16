package main

import (
	"fmt"
	"math"
)

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
//输入：root = [2,1,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1597 👎 0

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}

	value := isValidBST(root)
	no98Print("%+v, org: %+v", value, root)
}

func no98Print(format string, params ...interface{}) {
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

func isValidBST(root *TreeNode) bool {
	return bst(root, math.MinInt, math.MaxInt)
}

func bst(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return bst(root.Left, lower, root.Val) && bst(root.Right, root.Val, upper)
}

//func isValidBST(root *TreeNode) bool {
//	inSearch := validBSTInSearch(root)
//	if len(inSearch) <= 1 {
//		return true
//	}
//
//	for i := 1; i < len(inSearch); i++ {
//		if inSearch[i-1] >= inSearch[i] {
//			return false
//		}
//	}
//	return true
//}
//
//func validBSTInSearch(root *TreeNode) []int {
//
//	if root == nil {
//		return nil
//	}
//
//	var (
//		ans []int
//	)
//	left := validBSTInSearch(root.Left)
//	ans = append(ans, left...)
//	ans = append(ans, root.Val)
//	right := validBSTInSearch(root.Right)
//	ans = append(ans, right...)
//
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)
