// 图解参考: http://www.cnblogs.com/MrListening/p/5782752.html
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

func (n *Node) delete(parent bool) {
	if parent && n.parent != nil {
		if n.parent.left == n {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	}
	n.left = nil
	n.right = nil
	n.parent = nil
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

// 分四种情况处理:
// 1. 要删除的节点无左右孩子
// 2. 要删除的节点只有左孩子
// 3. 要删除的节点只有右孩子
// 4. 要删除的节点有左、右孩子
func (bst *BSTree) Delete(i int) bool {
	var parent *Node
	h := bst.head
	n := &Node{value: i}
	for h != nil {
		switch n.Compare(h) {
		case -1:
			parent = h
			h = h.left
		case 1:
			parent = h
			h = h.right
		case 0:
			isleftleaf := false
			if h != bst.head && parent.left == h { // 当前节点不是root节点
				isleftleaf = true
			}
			if h.left == nil && h.right == nil {
				if h == bst.head {
					bst.head = nil
					bst.size--
					return true
				}
				h.delete(true)
			}

			if h.left != nil && h.right == nil {
				if h == bst.head {
					bst.head = h.left
					bst.size--
					return true
				}
				if isleftleaf {
					parent.left = h.left
				} else {
					parent.right = h.left
				}
				h.delete(false)
			}

			if h.right != nil && h.left == nil {
				if h == bst.head {
					bst.head = h.right
					bst.size--
					return true
				}
				if isleftleaf {
					parent.left = h.right
				} else {
					parent.right = h.right
				}
				h.delete(false)
			}

			if h.left != nil && h.right != nil {
				leftmost := findLeftmost(h.right)
				h.value, leftmost.value = leftmost.value, h.value
				leftmost.delete(true)
			}
			bst.size--
			return true
		}
	}
	return false
}

func (bst *BSTree) Size() int {
	return bst.size
}

func findLeftmost(n *Node) *Node {
	if n == nil {
		return nil
	}
	ret := n
	for ret.left != nil {
		ret = ret.left
	}
	return ret
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
