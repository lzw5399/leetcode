/**
 * @Author: lzw5399
 * @Date: 2021/6/10 0:06
 * @Desc: 剑指 Offer 32 - I. 从上到下打印二叉树
 * @Desc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
 * @Desc: 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印
 * @Desc: 例如: 给定二叉树: [3,9,20,null,null,15,7]
 * @Desc: 返回 [3,9,20,15,7]
 * @Desc: [思路]: 借助一个队列进行层序遍历
 */
package main

import "fmt"

func main() {
	tree := NewSampleTree()
	result := levelOrder(tree)
	fmt.Println(result)
}

func levelOrder(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// pop一个
		node := queue[0]
		result = append(result, node.Val)

		queue = queue[1:]

		// 加入左右子树(如果有)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
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
