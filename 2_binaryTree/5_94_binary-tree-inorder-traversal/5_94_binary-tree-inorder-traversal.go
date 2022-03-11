/**
 * @Author: lzw5399
 * @Date: 2021/6/10 19:33
 * @Desc: 94. 二叉树的中序遍历
 * @Desc: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * @Desc: 中序遍历二叉树
 * @Desc: 左根右
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	result := inorderTraversal(tree)
	fmt.Println(result)
}

func inorderTraversal(root *TreeNode) []int {
	var nums []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return nums
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
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
