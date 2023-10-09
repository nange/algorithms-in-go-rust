package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	assert.Equal(t, 4, RemoveElement(nums, 4))

	val := 3
	res := RemoveElement(nums, val)
	assert.Equal(t, 2, res)
	assert.Equal(t, []int{2, 2}, nums[:res])
}
