/**
 * @Author: lzw5399
 * @Date: 2021/6/11 10:47
 * @Desc:
 */
package main

import "fmt"

func main() {
	palindromeNode := NewPalindromeSampleNode()
	is := isPalindrome(palindromeNode)
	fmt.Println(is)

	printLinkedList(palindromeNode)
}

// 判断是否是回文链表
// 1. 获取链表中间点
// 2. 反转后半链表
// 3. 判断前后半是否相等
// 4. 恢复后半链表
func isPalindrome(head *LinkedNode) bool {
	if head == nil {
		return true
	}

	// 确定中间点
	slowPtr, quickPtr := head, head
	for quickPtr.Next != nil && quickPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		quickPtr = quickPtr.Next.Next
	}

	// 反转后半
	secondHalfHead := reverseList(slowPtr.Next)

	// 比较
	result := true
	firstHalfNode := head
	secondHalfNode := secondHalfHead
	for result && secondHalfNode != nil {
		if firstHalfNode.Val != secondHalfNode.Val {
			result = false
		}
		firstHalfNode = firstHalfNode.Next
		secondHalfNode = secondHalfNode.Next
	}

	// 恢复后半
	slowPtr.Next = reverseList(secondHalfHead)

	return result
}

func reverseList(node *LinkedNode) *LinkedNode {
	var prev *LinkedNode
	current := node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

func NewPalindromeSampleNode() *LinkedNode {
	node5 := &LinkedNode{
		Val:  1,
		Next: nil,
	}
	node4 := &LinkedNode{
		Val:  2,
		Next: node5,
	}
	node3 := &LinkedNode{
		Val:  3,
		Next: node4,
	}
	//node35 := &LinkedNode{
	//	Val:  3,
	//	Next: node3,
	//}
	node2 := &LinkedNode{
		Val:  2,
		Next: node3,
	}
	node1 := &LinkedNode{
		Val:  1,
		Next: node2,
	}
	return node1
}

func printLinkedList(node *LinkedNode) {
	fmt.Println("开始print链表值")
	if node == nil {
		fmt.Println("链表为空")
	}
	tempNode := node
	for tempNode != nil {
		fmt.Printf(" %d -", tempNode.Val)
		tempNode = tempNode.Next
	}
}
