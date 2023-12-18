package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s1 := "babad"
	res := LongestPalindrome(s1)
	assert.Equal(t, "bab", res)

	s2 := "cbbd"
	res = LongestPalindrome(s2)
	assert.Equal(t, "bb", res)
}
