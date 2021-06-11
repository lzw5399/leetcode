/**
 * @Author: lzw5399
 * @Date: 2021/6/11 22:09
 * @Desc:
 */
package main

import "fmt"

func main() {
	r := NewSampleTree()
	sub := NewSampleSubTree()
	isSub := isSubStructure(r, sub)
	fmt.Println(isSub)
}

// 判断b是否是a的子树
// 先序遍历
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}

	if recur(a, b) {
		return true
	}

	return isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
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
