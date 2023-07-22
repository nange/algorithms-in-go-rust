package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MiddleNode(t *testing.T) {
	ret1 := MiddleNode(nil)
	assert.Nil(t, ret1)

	ret2 := MiddleNode(&ListNode{
		Val:  1,
		Next: nil,
	})
	assert.Equal(t, ret2, &ListNode{Val: 1, Next: nil})

	ret3 := MiddleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	assert.Equal(t, ret3, &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}})

	ret4 := MiddleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	})
	assert.Equal(t, ret4, &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}})
}
