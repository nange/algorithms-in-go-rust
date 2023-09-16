package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveDuplicates(t *testing.T) {
	var res1 = RemoveDuplicates(nil)
	assert.Equal(t, 0, res1)

	var arr = []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	var res2 = RemoveDuplicates(arr)
	assert.Equal(t, 4, res2)
	assert.Equal(t, []int{1, 2, 3, 4}, arr[0:res2])
}
