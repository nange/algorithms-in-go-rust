package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	// 第一步，让first节点先向前移动n+1步
	for i := 0; i < n+1; i++ {
		first = first.Next
		if first == nil && i < n { // 链表长度小于n，不需要删除任何节点
			return head
		}
	}

	// 第二步，让first和second节点同时向前移动，直到first节点到达链表的末尾
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 第三步，删除second节点的下一个节点
	second.Next = second.Next.Next

	return dummy.Next
}

type recursion struct {
	// 记录递归栈退出的次数，用于和n比较
	count int
}

func (r *recursion) RemoveNthFromEndWithRecursion(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = r.RemoveNthFromEndWithRecursion(head.Next, n)

	r.count++
	if r.count == n {
		return head.Next
	}

	return head
}
