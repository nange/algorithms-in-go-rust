package linkedlist

func Partition(head *ListNode, x int) *ListNode {
	// 左节点用于存储小于x的值, 右节点用于存储大于x的值
	var dummyLeft, dummyRight = &ListNode{}, &ListNode{}
	// p1, p2分别用生成新的左右链表节点
	var p1, p2 = dummyLeft, dummyRight

	// p指针用于遍历head链表
	var p = head

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开新生成的p1,p2指针和原链表的连接关系
		tmp := p.Next
		p.Next = nil
		p = tmp
	}

	p1.Next = dummyRight.Next

	return dummyLeft.Next
}
