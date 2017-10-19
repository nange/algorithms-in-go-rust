package binarytree

import (
	"log"
	"testing"
)

func TestBSTree(t *testing.T) {
	n := NewNode(1)
	m := NewNode(2)

	if n.Compare(m) != -1 || m.Compare(n) != 1 || n.Compare(n) != 0 {
		log.Println("n compare m:", n.Compare(m))
		t.Error("compare err")
		return
	}

	bst := NewTree(n)
	bst.Insert(4)
	bst.Insert(2)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(6)

	if bst.Size() != 6 {
		t.Errorf("tree size should be 6")
		return
	}

	five := bst.Find(5)
	if five.Value() != 5 {
		t.Errorf("node value should be 5")
		return
	}

	bst.Delete(5)
	if bst.Size() != 5 {
		t.Errorf("tree size should be 5")
		return
	}

	four := bst.Find(4)
	if four.Value() != 4 {
		t.Errorf("node value should be 4")
		return
	}

}
