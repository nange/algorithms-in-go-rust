package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseKGroup(t *testing.T) {
	ret1 := ReverseKGroup(nil, 2)
	assert.Nil(t, ret1)

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	ret2 := ReverseKGroup(list, 2)
	assert.Equal(t, ret2, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	})
}
