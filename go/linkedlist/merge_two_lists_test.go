package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeTwoLists(t *testing.T) {
	var fnList = []func(list1 *ListNode, list2 *ListNode) *ListNode{
		MergeTwoLists, MergeTwoListsWithRecursion,
	}

	for _, fn := range fnList {
		l1 := fn(nil, nil)
		assert.Nil(t, l1)

		l2 := fn(nil, &ListNode{})
		assert.Equal(t, l2, &ListNode{})

		l3 := fn(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		}, &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 6,
				},
			},
		})
		assert.Equal(t, l3, &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
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
		})
	}
}
