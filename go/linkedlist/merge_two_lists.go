package linkedlist

// Ref: https://leetcode.cn/problems/merge-two-sorted-lists/

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var p1 = list1
	var p2 = list2

	var dummy = &ListNode{}
	var p = dummy

	for {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next

		if p1 == nil || p2 == nil {
			break
		}
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

func MergeTwoListsWithRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		return &ListNode{
			Val:  list1.Val,
			Next: MergeTwoListsWithRecursion(list1.Next, list2),
		}
	} else {
		return &ListNode{
			Val:  list2.Val,
			Next: MergeTwoListsWithRecursion(list1, list2.Next),
		}
	}
}
