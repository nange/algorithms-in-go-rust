package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveNthFromEnd(t *testing.T) {
	ret1 := RemoveNthFromEnd(nil, 1)
	assert.Nil(t, ret1)

	lists := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	ret2 := RemoveNthFromEnd(lists, 1)
	assert.Equal(t, ret2, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
		},
	})

	lists2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	ret3 := RemoveNthFromEnd(lists2, 3)
	assert.Equal(t, ret3, &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
		},
	})
}
