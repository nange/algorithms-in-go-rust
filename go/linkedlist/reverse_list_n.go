package linkedlist

func ReverseN(head *ListNode, n int) *ListNode {
	h, _ := helper(head, n)
	return h
}

func helper(node *ListNode, n int) (*ListNode, *ListNode) {
	if n <= 0 || (node == nil || node.Next == nil) { // 边界条件
		return node, node
	}

	if n == 1 { // 递归退出条件
		return node, node.Next
	}

	last, successor := helper(node.Next, n-1)

	node.Next.Next = node
	node.Next = successor

	return last, successor
}
