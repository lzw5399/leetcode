/**
 * @Author: lzw5399
 * @Date: 2021/6/11 23:41
 * @Desc: 226. 翻转二叉树
 * @Desc: https://leetcode-cn.com/problems/invert-binary-tree/
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	PrintTree(tree)
	invertTree(tree)
	PrintTree(tree)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewSampleTree() *TreeNode {
	leftChildLeftChild := &TreeNode{
		Val:   99,
		Left:  nil,
		Right: nil,
	}
	leftChildRightChild := &TreeNode{
		Val:   66,
		Left:  nil,
		Right: nil,
	}

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
		Left:  leftChildLeftChild,
		Right: leftChildRightChild,
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

func PrintTree(root *TreeNode) {
	fmt.Println("---开始打印二叉树结构---")
	var result [][]int
	if root == nil {
		fmt.Println("---二叉树root节点为空---")
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var temp []int

		length := len(queue)
		for i := 0; i < length; i++ {
			// 获取node
			node := queue[0]
			queue = queue[1:]

			temp = append(temp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, temp)
	}

	for _, v := range result {
		fmt.Println(v)
	}
	fmt.Println("---打印二叉树结构完毕---")
}
