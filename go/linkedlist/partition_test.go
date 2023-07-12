package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	ret1 := Partition(nil, 0)
	assert.Nil(t, ret1)

	ret2 := Partition(&ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: -1,
				}},
		}}, 1)
	assert.Equal(t, ret2, &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: -1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	})
}
