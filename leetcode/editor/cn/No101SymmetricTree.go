package main

import "fmt"

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1942 👎 0

func main() {

	root := sliceToTree([]int{1, 2, 2, 3, 4, 4, 3})
	//root := sliceToTree([]int{2, 3, 3, 4, 5, 5, 4, -100000, -100000, 8, 9, -100000, -100000, 9, 8})
	value := isSymmetric(root)
	no101Print("%+v", value)
}

func no101Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

func sliceToTree(treeValues []int) *TreeNode {
	if len(treeValues) == 0 {
		return nil
	}
	var (
		lastQueue []*TreeNode
		level     = 0
	)
	root := &TreeNode{Val: treeValues[0]}
	treeValues = treeValues[1:]
	lastQueue = append(lastQueue, root)
	for len(treeValues) > 0 {
		count := 1 << level
		for i := 0; i < count; i++ {
			node := lastQueue[0]
			lastQueue = lastQueue[1:]
			if len(treeValues) >= 2 {
				left := &TreeNode{Val: treeValues[0]}
				right := &TreeNode{Val: treeValues[1]}
				node.Left = left
				node.Right = right
				lastQueue = append(lastQueue, left)
				lastQueue = append(lastQueue, right)
				treeValues = treeValues[2:]
			}
		}
	}
	return root
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil || root.Left.Val != root.Right.Val {
		return false
	}

	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		l := len(queue)
		var ans []int
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil && node.Left.Val != -100000 {
				queue = append(queue, node.Left)
				ans = append(ans, node.Left.Val)
			} else {
				ans = append(ans, -100000)
			}
			if node.Right != nil && node.Right.Val != -100000 {
				queue = append(queue, node.Right)
				ans = append(ans, node.Right.Val)
			} else {
				ans = append(ans, -100000)
			}
		}
		left := 0
		right := len(ans) - 1
		for left < right {
			if ans[left] != ans[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return isSame(left.Left, right.Right) && isSame(right.Left, left.Right)
}
