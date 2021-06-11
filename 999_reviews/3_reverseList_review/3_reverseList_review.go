/**
 * @Author: lzw5399
 * @Date: 2021/6/11 10:28
 * @Desc:
 */
package main

import "fmt"

func main() {
	linkedNode := NewSampleLinkedList()
	reversedNode := reverseList(linkedNode)

	fmt.Println(reversedNode)
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

func NewSampleLinkedList() *LinkedNode {
	node5 := &LinkedNode{
		Val:  5,
		Next: nil,
	}
	node4 := &LinkedNode{
		Val:  4,
		Next: node5,
	}
	node3 := &LinkedNode{
		Val:  3,
		Next: node4,
	}
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

