package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseN(t *testing.T) {
	ret1 := ReverseN(nil, 0)
	assert.Nil(t, ret1)

	ret2 := ReverseN(&ListNode{
		Val:  1,
		Next: nil,
	}, 2)
	assert.Equal(t, ret2, &ListNode{
		Val:  1,
		Next: nil,
	})

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ret3 := ReverseN(list, 3)
	assert.Equal(t, ret3, &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	})
}
