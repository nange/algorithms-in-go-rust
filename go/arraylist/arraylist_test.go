package arraylist

import "testing"

func TestSize(t *testing.T) {
	list := New()
	if list.Size() != 0 {
		t.Errorf("new arraylist's size should be zero.")
	}

	list.Append("hello")
	list.Append("world")
	if list.Size() != 2 {
		t.Errorf("arraylist#Size() failed. error list size:%d, expected size:2.", list.Size())
	}
}

func TestGet(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")
	v, err := list.Get(1)
	if sv, ok := v.(string); !ok || sv != "world" || err != nil {
		t.Errorf("arraylist#Get() failed. get for error value:%s, expected value:world.")
	}
}

func TestSet(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")
	err := list.Set(2, "nange")
	if err == nil {
		t.Errorf("arraylist#Set() failed. Index out of range. should have error.")
	}
	err2 := list.Set(1, "nange")
	if err2 != nil {
		t.Errorf("arraylist#Set() failed. Set method should not have error.")
	}
	newItem, _ := list.Get(1)
	if sv, ok := newItem.(string); !ok || sv != "nange" {
		t.Errorf("arraylist#Set() failed. Get a invalid value from list: %s,"+
			"excepted value is: nange.", sv)
	}
}

func TestInsert(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	err := list.Insert(0, 4)
	if err != nil {
		t.Errorf("arraylist#Insert() failed. Should not have error.")
	}
	first, _ := list.Get(0)
	second, _ := list.Get(1)
	third, _ := list.Get(2)
	four, _ := list.Get(3)
	if first.(int) != 4 || second.(int) != 1 || third.(int) != 2 ||
		four.(int) != 3 || list.Size() != 4 {
		t.Errorf("arraylist#Insert() failed. Got:%d, %d, %d, %d. Excepted:4, 1, 2, 3.",
			first.(int), second.(int), third.(int), four.(int))
	}
}

func TestAppend(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")
	list.Append("nange")

	if sv, _ := list.Get(2); sv != "nange" || list.Size() != 3 {
		t.Error("arraylist#Append() failed.")
	}
}

func TestRemove(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")

	err := list.Remove(1)
	if sv, _ := list.Get(0); sv != "hello" || err != nil || list.Size() != 1 {
		t.Error("arraylist#Remove() failed.")
	}
}

func TestClear(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")

	list.Clear()
	if list.Size() != 0 {
		t.Errorf("arraylist#Clear() failed. size: %d, excepted size: 0", list.Size())
	}
}

func TestString(t *testing.T) {
	list := New()
	list.Append("hello")
	list.Append("world")

	s := list.String()
	if s != "[hello world]" {
		t.Error("arraylist#Clear() failed.")
	}
}
