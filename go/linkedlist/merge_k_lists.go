package linkedlist

import "container/heap"

func MergeKLists(lists []*ListNode) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{}
	// 指向虚拟头节点最后的节点，用于设置新节点
	p := dummy

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, head := range lists {
		// 将K个链表头节点加入优先级队列
		if head != nil {
			heap.Push(&pq, head)
		}
	}

	// 从优先级队列中逐个取出节点加入新链表中，
	// 并把新的链表头节点加入优先级队列中
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		p = p.Next
	}

	return dummy.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
