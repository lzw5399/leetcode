package main

func main() {

}

func rightSideView(root *TreeNode) (nums []int) {
	if root == nil {
		return
	}

	// 基于bfs
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if i == length-1 {
				nums = append(nums, node.Val)
			}
		}
	}

	return
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
