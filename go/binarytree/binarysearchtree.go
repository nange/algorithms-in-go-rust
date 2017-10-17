package binarytree

type BinarySearchTree struct {
	data  int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func PreOrder(root *BinarySearchTree, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.data)
	PreOrder(root.left, ret)
	PreOrder(root.right, ret)
}

func InOrder(root *BinarySearchTree, ret *[]int) {
	if root == nil {
		return
	}
	InOrder(root.left, ret)
	*ret = append(*ret, root.data)
	InOrder(root.right, ret)
}

func PostOrder(root *BinarySearchTree, ret *[]int) {
	if root == nil {
		return
	}
	PostOrder(root.left, ret)
	PostOrder(root.right, ret)
	*ret = append(*ret, root.data)
}
