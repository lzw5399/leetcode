package main

func main() {

}

func levelOrder(root *TreeNode) (nums []int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// pop 第一个
		fir := queue[0]
		nums = append(nums, fir.Val)
		queue = queue[1:]

		// 加入左
		if fir.Left != nil {
			queue = append(queue, fir.Left)
		}

		// 加入右
		if fir.Right != nil {
			queue = append(queue, fir.Right)
		}
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
