/**
 * @Author: lzw5399
 * @Date: 2021/6/12 20:56
 * @Desc: 236. 二叉树的最近公共祖先
 * @Desc: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 * @Desc: 求最近的公共祖先，可以是他自己
 */
package main

import "fmt"

func main() {
	root, p, q := NewSampleTree()
	PrintTree(root)
	result := lowestCommonAncestor(root, p, q)
	PrintTree(result)
}

// 最近的公共祖先
// 思路: 二叉树的后序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewSampleTree() (*TreeNode, *TreeNode, *TreeNode) {
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

	return root, leftChildLeftChild, leftChildRightChild
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
