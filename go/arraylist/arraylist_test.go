package arraylist

import "testing"

func TestSize(t *testing.T) {
	list := NewArrayList()
	if list.Size() != 0 {
		t.Errorf("new arraylist's size should be zero.")
	}

	list.Append("hello")
	list.Append("world")
	if list.Size() != 2 {
		t.Errorf("error list size:%d, expected size:2.", list.Size())
	}
}

func TestGet(t *testing.T) {
	list := NewArrayList()
	list.Append("hello")
	list.Append("world")
	v, err := list.Get(1)
	if sv, ok := v.(string); !ok || sv != "world" || err != nil {
		t.Errorf("get for error value:%s, expected value:world.")
	}
}
