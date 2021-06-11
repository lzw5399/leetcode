/**
 * @Author: lzw5399
 * @Date: 2021/6/10 19:33
 * @Desc: 面试题 02.06. 回文链表
 * @Desc: https://leetcode-cn.com/problems/palindrome-linked-list-lcci/
 * @Desc: 编写一个函数，检查输入的链表是否是回文的。
 * @Desc: 输入： 1->2 输出： false
 * @Desc: 输入： 1->2->2->1 输出： true
 */
package main

import "fmt"

func main() {
	palindromeNode := NewPalindromeSampleNode()
	is := isPalindrome(palindromeNode)
	fmt.Println(is)

	printLinkedList(palindromeNode)
}

// 整个流程可以分为以下五个步骤：
// 1. 找到前半部分链表的尾节点 (快慢指针)
// 2. 反转后半部分链表
// 3. 判断是否回文
// 4. 恢复链表
// 5. 返回结果
func isPalindrome(node *ListNode) bool {
	if node == nil {
		return true
	}

	// 找到前后半的分界节点(快慢指针)
	fastPtr, slowPtr := node, node
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	// 2. 反转后半部分链表
	latterHalf := reverseList(slowPtr.Next)

	// 3. 比较前后是否相等
	result := true
	firstHalfNode := node
	latterHalfNode := latterHalf
	for result && latterHalfNode != nil {
		if firstHalfNode.Val != latterHalfNode.Val {
			result = false
		}
		firstHalfNode = firstHalfNode.Next
		latterHalfNode = latterHalfNode.Next
	}

	// 4. 恢复后半链表
	slowPtr.Next = reverseList(latterHalf)

	return result
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

//
//func isPalindromeByArr(node *ListNode) bool {
//	var arr []int
//	tempNode := node
//	for tempNode != nil {
//		arr = append(arr, tempNode.Val)
//		tempNode = tempNode.Next
//	}
//
//	for i := 0; i < len(arr)/2; i++ {
//		if arr[i] != arr[len(arr)-i-1] {
//			return false
//		}
//	}
//
//	return true
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewPalindromeSampleNode() *ListNode {
	node5 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  2,
		Next: node5,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node35 := &ListNode{
		Val:  3,
		Next: node3,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node35,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	return node1
}

func printLinkedList(node *ListNode) {
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