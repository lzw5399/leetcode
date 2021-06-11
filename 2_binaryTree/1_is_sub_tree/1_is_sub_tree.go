/**
 * @Author: lzw5399
 * @Date: 2021/6/11 14:53
 * @Desc: 剑指 Offer 26. 树的子结构
 * @Desc: https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
 * @Desc: 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
 * @Desc: B是A的子结构， 即 A中有出现和B相同的结构和节点值。
 * @Desc: [思路]: 先序遍历 & 包含判断
 */
package main

import "fmt"

func main() {
	r := NewSampleTree()
	sub := NewSampleSubTree()
	isSub := isSubStructure(r, sub)
	fmt.Println(isSub)
}

// 思路: 先序遍历 & 包含判断
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}

	if recur(a, b) {
		return true
	}

	if isSubStructure(a.Left, b) || isSubStructure(a.Right, b) {
		return true
	}

	return false
}

func recur(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return recur(a.Left, b.Left) && recur(a.Right, b.Right)
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

func NewSampleSubTree() *TreeNode {
	rightChild := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	leftChild := &TreeNode{
		Val:   15,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   20,
		Left:  leftChild,
		Right: rightChild,
	}

	return root
}
