package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeKLists(t *testing.T) {
	ret1 := MergeKLists(nil)
	assert.Nil(t, ret1)
	ret2 := MergeKLists([]*ListNode{})
	assert.Nil(t, ret2)

	arrList := []*ListNode{
		{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	ret3 := MergeKLists(arrList)
	assert.Equal(t, ret3, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
								},
							},
						},
					},
				},
			},
		},
	})
}
