/**
 * @Author: lzw5399
 * @Date: 2021/6/12 22:31
 * @Desc: 143. 重排链表
 * @Desc: https://leetcode-cn.com/problems/reorder-list/
 * @Desc: 示例1: 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 * @Desc: 示例2: 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 */
package main

import "fmt"

func main() {
	list := NewSampleLinkedList()
	printLinkedList(list)

	reorderList(list)
	printLinkedList(list)
}

// 将后半的链表逆序，然后重新拼接
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 快慢指针 获取中间点
	slowPtr, quickPtr := head, head
	for quickPtr.Next != nil && quickPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		quickPtr = quickPtr.Next.Next
	}

	// 逆转后半链表
	firstHalfHead := head
	secondHalfHead := reverseList(slowPtr.Next)
	slowPtr.Next = nil

	// 合并两个
	firstHalfNode := firstHalfHead
	secondHalfNode := secondHalfHead
	for firstHalfNode != nil && secondHalfNode != nil {
		firstNext := firstHalfNode.Next
		firstHalfNode.Next = secondHalfNode

		secondNext := secondHalfNode.Next
		secondHalfNode.Next = firstNext

		firstHalfNode = firstNext
		secondHalfNode = secondNext
	}
}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	current := node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewSampleLinkedList() *ListNode {
	node5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  4,
		Next: node5,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	return node1
}

func printLinkedList(node *ListNode) {
	fmt.Println("---开始print链表值---")
	if node == nil {
		fmt.Println("链表为空")
	}
	tempNode := node
	for tempNode != nil {
		fmt.Printf(" %d -", tempNode.Val)
		tempNode = tempNode.Next
	}
	fmt.Println()
	fmt.Println("---打印链表值结束---")
}
