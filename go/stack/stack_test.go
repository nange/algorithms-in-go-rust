package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New()
	if stack.Size() != 0 || !stack.IsEmpty() {
		t.Errorf("stack#Size() or stack#IsEmpty() failed. new stack size should be 0.")
	}

	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("stack#Push() failed.")
	}

	p := stack.Pop()
	if v, ok := p.(int); !ok || v != 2 || stack.Size() != 1 {
		t.Errorf("stack#Pop() failed.")
	}

	p2 := stack.Peek()
	if v, ok := p2.(int); !ok || v != 1 || stack.Size() != 1 {
		t.Errorf("stack#Peek() failed.")
	}

	stack.Clear()
	if !stack.IsEmpty() {
		t.Errorf("stack#Clear() failed.")
	}

}
