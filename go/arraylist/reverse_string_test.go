package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	s := []byte("hello world")
	ReverseString(s)
	assert.Equal(t, []byte("dlrow olleh"), s)
}
