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
	nodeCount := 1
	for fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		if fastPtr.Next != nil {
			nodeCount++
			fastPtr = fastPtr.Next
		}
		if fastPtr.Next != nil {
			nodeCount++
			fastPtr = fastPtr.Next
		}
	}

	// 2. 反转后半部分链表
	result := true
	// 当为偶数时候，slowPtr在后半部分的开头
	// 奇数的时候，slowPtr在中间，所以要后移
	if nodeCount%2 == 1 {
		slowPtr = slowPtr.Next
	}
	latterHalf := reverseList(slowPtr, nil)

	// 3. 比较前后是否相等
	firstHalfNode := node
	latterHalfNode := latterHalf
	var latterPrevNode *ListNode // latterHalf的前一个节点
	for result && latterHalfNode != nil {
		if firstHalfNode.Val != latterHalfNode.Val {
			result = false
		}
		firstHalfNode = firstHalfNode.Next
		latterHalfNode = latterHalfNode.Next

		if latterHalfNode == nil && firstHalfNode.Next == latterHalfNode {
			latterPrevNode = firstHalfNode.Next
		} else if latterHalfNode == nil && firstHalfNode.Next.Next == latterHalfNode {
			latterPrevNode = firstHalfNode.Next.Next
		}
	}

	// 4. 恢复后半链表
	reverseList(latterHalf, latterPrevNode)

	return result
}

func reverseList(node *ListNode, prev *ListNode) *ListNode {
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
	//node35 := &ListNode{
	//	Val:  35,
	//	Next: node3,
	//}
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
