package linkedlist

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	lenA := 0
	lenB := 0

	for p1 != nil {
		p1 = p1.Next
		lenA++
	}
	for p2 != nil {
		p2 = p2.Next
		lenB++
	}
	if lenA == 0 || lenB == 0 {
		return nil
	}

	p1 = headA
	p2 = headB
	if lenA >= lenB { // A 比 B 长，则A先走 lenA - lenB 步
		for i := 0; i < lenA-lenB; i++ {
			p1 = p1.Next
		}
	} else { // B 比 A 长，则B先走 lenB - lenA 步
		for i := 0; i < lenB-lenA; i++ {
			p2 = p2.Next
		}
	}
	// 再一起往后走，直到相遇
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}
