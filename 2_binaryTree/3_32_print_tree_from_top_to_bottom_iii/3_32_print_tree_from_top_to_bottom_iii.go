/**
 * @Author: lzw5399
 * @Date: 2021/6/10 0:38
 * @Desc: 剑指 Offer 32 - III. 从上到下打印二叉树 III
 * @Desc: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
 * @Desc: 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
 * @Desc: 例如: 给定二叉树: [3,9,20,null,null,15,7]
 * @Desc: 返回 [[3], [20, 9], [15,7]]
 * @Desc: [思路]: 还是借助一个队列进行层序遍历
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
	level := 0
	for len(queue) != 0 {
		var temp []int

		length := len(queue)

		level++
		for i := 0; i < length; i++ {
			// 获取node
			node := queue[0]
			queue = queue[1:]

			// 偶数层prepend ！！！
			if level%2 == 0 {
				temp = append([]int{node.Val}, temp...)
			} else { // 奇数层append
				temp = append(temp, node.Val)
			}

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
