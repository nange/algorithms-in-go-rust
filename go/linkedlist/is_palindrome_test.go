package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsPalindrome(t *testing.T) {
	fns := []func(node *ListNode) bool{IsPalindromeV1, IsPalindromeV2}
	for _, fn := range fns {
		ret1 := fn(nil)
		assert.False(t, ret1)

		ret2 := fn(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		})
		assert.True(t, ret2)

		ret3 := fn(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		})
		assert.False(t, ret3)

		ret4 := fn(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		})
		assert.True(t, ret4)

		ret5 := fn(&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		})
		assert.False(t, ret5)
	}
}
