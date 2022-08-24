package internal

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	value interface{}
}

func (bt *BinaryTree) SetLeft(l *BinaryTree) {
	bt.left = l
}

func (bt *BinaryTree) GetLeft() *BinaryTree {
	return bt.left
}

func (bt *BinaryTree) SetRight(l *BinaryTree) {
	bt.right = l
}

func (bt *BinaryTree) GetRight() *BinaryTree {
	return bt.right
}

func (bt *BinaryTree) SetValue(v interface{}) {
	bt.value = v
}

func (bt *BinaryTree) GetValue() interface{} {
	return bt.value
}
