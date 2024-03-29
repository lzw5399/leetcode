/**
 * @Author: lzw5399
 * @Date: 2021/6/11 22:27
 * @Desc: 199. 二叉树的右视图
 * @Desc: https://leetcode-cn.com/problems/binary-tree-right-side-view/
 * @Desc: 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	PrintTree(tree)
	result := rightSideView(tree)
	fmt.Println(result)
}

// 走广度优先bfs的解法
// 一个队列，root入队
func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// 出队并且他的左右孩子入队
		length := len(queue)
		for i := 0; i < length; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// 最后一个
			if i == length-1 {
				result = append(result, node.Val)
			}
		}
	}

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewSampleTree() *TreeNode {
	leftChildRightChild := &TreeNode{
		Val:   88,
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
		Left:  nil,
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
