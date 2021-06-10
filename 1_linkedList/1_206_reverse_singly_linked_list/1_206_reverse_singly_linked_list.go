/**
 * @Author: lzw5399
 * @Date: 2021/6/10 19:46
 * @Desc: 206. 反转链表
 * @Desc: https://leetcode-cn.com/problems/reverse-linked-list/
 * @Desc: 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 */
package main

import "fmt"

func main() {
	linkedNode := NewSampleLinkedList()
	reversedNode := reverseList(linkedNode)

	fmt.Println(reversedNode)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
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








//191
//191
//191
//1911
//19
//19
//
//19














