package linkedlist

func ReverseList(head *ListNode) *ListNode {
	// 定义递归退出条件
	if head == nil || head.Next == nil {
		return head
	}
	// 递归返回最后一个节点作为头结点返回
	last := ReverseList(head.Next)

	// 从倒数第二个节点开始，对其后面的一个节点执行链表反转
	head.Next.Next = head
	// 从倒数第二个节点开始，将自身的Next指针递归指向Nil
	//（中间节点其实不用这一步也可以，但是因为最后一个head节点必须要这样做，因此统一做最好）
	head.Next = nil

	return last
}
