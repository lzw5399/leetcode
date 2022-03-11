/**
 * @Author: lzw5399
 * @Date: 2021/6/10 19:33
 * @Desc: 144. 二叉树的前序遍历
 * @Desc: https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
 * @Desc: 前序遍历二叉树
 * @Desc: 根左右
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	result := preOrderTraversal(tree)
	fmt.Println(result)
}

func preOrderTraversal(node *TreeNode) (nums []int) {
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		nums = append(nums, root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(node)
	return
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
