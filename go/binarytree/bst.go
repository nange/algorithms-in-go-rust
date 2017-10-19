package binarytree

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

func NewNode(i int) *Node {
	return &Node{value: i}
}

func (n *Node) Compare(m *Node) int {
	if n.value < m.value {
		return -1
	} else if n.value > m.value {
		return 1
	} else {
		return 0
	}
}

func (n *Node) Value() int {
	return n.value
}

type BSTree struct {
	head *Node
	size int
}

func NewTree(root *Node) *BSTree {
	if root == nil {
		return &BSTree{}
	}
	return &BSTree{head: root, size: 1}
}

func (bst *BSTree) Insert(i int) {
	n := &Node{value: i}
	if bst.head == nil {
		bst.head = n
		bst.size++
		return
	}

	h := bst.head

	for {
		if n.Compare(h) == -1 {
			if h.left != nil {
				h = h.left
				continue
			}
			h.left = n
			n.parent = h
		} else {
			if h.right != nil {
				h = h.right
				continue
			}
			h.right = n
			n.parent = h
		}
		break
	}
	bst.size++
}

func (bst *BSTree) Find(i int) *Node {
	h := bst.head
	n := &Node{value: i}

	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.right
		case 1:
			h = h.left
		case 0:
			return h
		default:
			panic("Node not found") // this should not happen
		}
	}
	return nil
}

func (bst *BSTree) Delete(i int) bool {
	h := bst.head
	n := &Node{value: i}
	for h != nil {
		switch n.Compare(h) {
		case -1:
			h = h.left
		case 1:
			h = h.right
		case 0:
			if h.left != nil {
				h.parent.left = h.left
			}
			if h.right != nil {
				h.parent.right = h.right
			}
			h = nil
			bst.size--
			return true
		}
	}
	return false
}

func (bst *BSTree) Size() int {
	return bst.size
}

func PreOrder(root *Node, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.value)
	PreOrder(root.left, ret)
	PreOrder(root.right, ret)
}

func InOrder(root *Node, ret *[]int) {
	if root == nil {
		return
	}
	InOrder(root.left, ret)
	*ret = append(*ret, root.value)
	InOrder(root.right, ret)
}

func PostOrder(root *Node, ret *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.left, ret)
	PostOrder(root.right, ret)
	*ret = append(*ret, root.value)
}
