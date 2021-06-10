/**
 * @Author: lzw5399
 * @Date: 2021/6/10 0:38
 * @Desc: 剑指 Offer 32 - II. 从上到下打印二叉树 II
 * @Desc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
 * @Desc: 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
 * @Desc: 例如: 给定二叉树: [3,9,20,null,null,15,7]
 * @Desc: 返回 [[3], [9,20], [15,7]]
 * @Desc: [思路]: 还是借助一个队列进行层序遍历，对每一层
 */
package main

import "fmt"

func main() {
	root := NewSampleTree()
	result := levelOrder(root)
	fmt.Println(result)
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
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
