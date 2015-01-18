package queue

import "testing"

func TestAllMethod(t *testing.T) {
	q := New()
	if q.Size() != 0 || !q.IsEmpty() {
		t.Errorf("queue#Size() or queue#IsEmpty() failed. new queue size should be 0.")
	}

	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	if q.Size() != 3 || q.IsEmpty() {
		t.Errorf("queue#Enqueue() failed.")
	}

	d := q.Dequeue()
	if v, ok := d.(string); !ok || v != "1" {
		t.Errorf("queue#Dequeue() failed. v=%s", v)
	}

	f := q.Font()
	if v, ok := f.(string); !ok || v != "2" {
		t.Errorf("queue#Font() failed.")
	}

	e := q.End()

	if v, ok := e.(string); !ok || v != "3" {
		t.Errorf("queue#End() failed.")
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Errorf("queue#IsEmpty() failed.")
	}

}
