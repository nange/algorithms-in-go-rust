package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DetectCycle(t *testing.T) {
	ret1 := DetectCycle(nil)
	assert.Nil(t, ret1)

	ret2 := DetectCycle(&ListNode{
		Val:  1,
		Next: &ListNode{},
	})
	assert.Nil(t, ret2)

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	node1 := &ListNode{
		Val: 3,
	}
	node2 := &ListNode{
		Val: 4,
	}
	head.Next = node1
	node1.Next = node2
	node2.Next = node1

	ret3 := DetectCycle(head)
	assert.Equal(t, ret3.Val, 3)
}
