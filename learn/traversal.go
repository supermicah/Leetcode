package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (n *TreeNode) String() string {
	return fmt.Sprintf("{val: %d, left: %s, right: %s}", n.Val, n.Left, n.Right)
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	/***
	         1
	    2         3
	  4   5     6   7
	*/
	println("------------------")
	println(fmt.Sprintf("preorderRecursiveTraversal: %+v", preorderRecursiveTraversal(root)))
	println(fmt.Sprintf("inorderRecursiveTraversal: %+v", inorderRecursiveTraversal(root)))
	println(fmt.Sprintf("postorderRecursiveTraversal: %+v", postorderRecursiveTraversal(root)))
	println("")
	println(fmt.Sprintf("preorderTraversal: %+v", preorderTraversal(root)))
	println(fmt.Sprintf("inorderTraversal: %+v", inorderTraversal(root)))
	println(fmt.Sprintf("postorderTraversal: %+v", postorderTraversal(root)))

	println("")
	println("------------------")
	println(fmt.Sprintf("preorderDFSRecursiveTraversal: %+v", preorderDFSRecursiveTraversal(root)))

	println("")
	println("------------------")
	println(fmt.Sprintf("preorderBFSTraversal: %+v", preorderBFSTraversal(root)))

}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}

	return result
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		result []int
		stack  []*TreeNode
	)

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right

	}

	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		result    []int
		stack     []*TreeNode
		lastVisit *TreeNode
	)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}

func preorderRecursiveTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int

	result = append(result, root.Val)
	result = append(result, preorderRecursiveTraversal(root.Left)...)
	result = append(result, preorderRecursiveTraversal(root.Right)...)
	return result
}

func inorderRecursiveTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, inorderRecursiveTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderRecursiveTraversal(root.Right)...)

	return result
}

func postorderRecursiveTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	result = append(result, postorderRecursiveTraversal(root.Left)...)
	result = append(result, postorderRecursiveTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}

func preorderDFSRecursiveTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}
	result = append(result, root.Val)
	result = append(result, preorderDFSRecursiveTraversal(root.Left)...)
	result = append(result, preorderDFSRecursiveTraversal(root.Right)...)

	return result
}

func preorderBFSTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		result []int
		stack  []*TreeNode
	)

	stack = append(stack, root)
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			result = append(result, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return result
}
