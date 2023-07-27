package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetIntersectionNode(t *testing.T) {
	ret1 := GetIntersectionNode(nil, nil)
	assert.Nil(t, ret1)

	head1 := &ListNode{Val: 1}
	head2 := &ListNode{Val: 10}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 11}
	nodeInte := &ListNode{Val: 3, Next: &ListNode{Val: 4}}

	head1.Next = node1
	node1.Next = nodeInte
	head2.Next = node2
	node2.Next = nodeInte

	ret2 := GetIntersectionNode(head1, head2)
	assert.Equal(t, ret2, nodeInte)
}
