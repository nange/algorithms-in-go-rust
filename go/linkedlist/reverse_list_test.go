package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseList(t *testing.T) {
	ret1 := ReverseList(nil)
	assert.Nil(t, ret1)

	ret2 := ReverseList(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	assert.Equal(t, ret2, &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	})
}
