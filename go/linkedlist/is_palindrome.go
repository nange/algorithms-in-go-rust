package linkedlist

// IsPalindromeV1 V1版本采用递归方式，代码简单，空间复杂度为O(N)
func IsPalindromeV1(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	h := head

	var traverse func(node *ListNode) bool
	traverse = func(node *ListNode) bool {
		if node == nil {
			return true
		}

		eq := traverse(node.Next)
		eq = eq && h.Val == node.Val
		h = h.Next

		return eq
	}

	return traverse(head)
}

// IsPalindromeV2 不用递归，使空间复杂度变为O(1)
func IsPalindromeV2(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}

	middle := findMiddle(head)
	rev := reverseNode(middle)

	h := head
	for rev != nil {
		if rev.Val != h.Val {
			return false
		}
		rev = rev.Next
		h = h.Next
	}

	return true
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseNode(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	prev := (*ListNode)(nil)
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
