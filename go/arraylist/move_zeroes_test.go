package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, nums)
}
