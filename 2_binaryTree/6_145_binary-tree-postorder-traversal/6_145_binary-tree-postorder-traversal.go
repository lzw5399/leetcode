/**
 * @Author: lzw5399
 * @Date: 2021/6/10 19:33
 * @Desc: 145. 二叉树的后序遍历
 * @Desc: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
 * @Desc: 后序遍历二叉树
 * @Desc: 左右根
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	result := postorderTraversal(tree)
	fmt.Println(result)
}

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		nums = append(nums, node.Val)
	}
	postOrder(root)
	return nums
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewSampleTree() *TreeNode {
	rightChildLeftChild := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	rightChildRightChild := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	leftChild := &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}

	rightChild := &TreeNode{
		Val:   20,
		Left:  rightChildLeftChild,
		Right: rightChildRightChild,
	}

	root := &TreeNode{
		Val:   3,
		Left:  leftChild,
		Right: rightChild,
	}

	return root
}
