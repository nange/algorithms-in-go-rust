package linkedlist

// 翻转[a, b)的元素
func reverse(a, b *ListNode) *ListNode {
	var pre, curr, next *ListNode = nil, a, a
	// 移动到b指针时，停止
	for curr != b {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	// 返回翻转后的头结点
	return pre
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到K的翻转区间
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil { // 不足k个节点时，直接返回head
			return head
		}
		b = b.Next
	}

	// 翻转前k个元素
	newHead := reverse(a, b)

	// 递归执行此操作
	a.Next = ReverseKGroup(b, k)

	return newHead
}
