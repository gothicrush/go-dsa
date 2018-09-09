package bst

import (
    //"../../stack/linkedstack"
    "../../queue/linkedqueue"
    "../../assist"
)

type node struct {
	data interface{}
	left *node
	right *node
}

func newNode(data interface{}) *node {
	return &node {
		data: data,
		left: nil,
		right: nil,
	}
}

type BST struct {
	root *node
	size int
	comparator assist.Comparator
}

func New(comparator assist.Comparator) *BST {
	return &BST {
		root: nil,
		size: 0,
		comparator: comparator,
	}
}

func (bst *BST) Add(v interface{}) {
	bst.root = bst.add(bst.root, v)
}

func (bst *BST) add(nd *node, v interface{}) *node {
	if nd == nil {
		bst.size++
		return newNode(v)
	}

	cmp := assist.Compare(nd.data, v, bst.comparator)
	if cmp > 0 {
		nd.left = bst.add(nd.left, v)
	} else if cmp < 0 {
		nd.right = bst.add(nd.right, v)
	}

	return nd
}

func (bst *BST) Remove(v interface{}) {
	bst.root = bst.remove(bst.root, v)
}

func (bst *BST) remove(nd *node, v interface{}) *node {

	if nd == nil {
		return nil
	}

	if nd.left == nil {
		var rightNode *node = nd.right
		nd.right = nil
		bst.size--
		return rightNode
	}

	if nd.right == nil {
		var leftNode *node = nd.left
		nd.left = nil
		bst.size--
		return leftNode
	}

	var successor *node = bst.getMinNode(nd.right)
	successor.right = bst.removeMin(nd.right) // 已经 size-- 了
	successor.left = nd.left

	nd.left = nil
	nd.right = nil

	return successor
}

func (bst *BST) RemoveMin(v interface{}) interface{} {

	var minValue interface{} = bst.GetMin()

	bst.root = bst.removeMin(bst.root)

	return minValue
}

func (bst *BST) removeMin(nd *node) *node {

	if nd.left == nil {
		var rightNode *node = nd.right
		nd.right = nil
		bst.size--
		return rightNode
	}

	nd.left = bst.removeMin(nd.left)

	return nd
}

func (bst *BST) RemoveMax(v interface{}) interface{} {

	var maxValue interface{} = bst.GetMax()

	bst.root = bst.removeMax(bst.root)

	return maxValue
}

func (bst *BST) removeMax(nd *node) *node {

	if nd.right == nil {
		var leftNode *node = nd.left
		nd.left = nil
		bst.size--
		return leftNode
	}

	nd.right = bst.removeMax(nd.right)

	return nd
}

func (bst *BST) PreOrder(visitor assist.Visitor) {
	bst.preOrder(bst.root, visitor)
}

func (bst *BST) preOrder(nd *node, visitor assist.Visitor) {
	if nd == nil {
		return
	}

	visitor(nd)
	bst.preOrder(nd.left, visitor)
	bst.preOrder(nd.right, visitor )
}

func (bst *BST) InOrder(visitor assist.Visitor) {
	bst.inOrder(bst.root, visitor)
} 

func (bst *BST) inOrder(nd *node, visitor assist.Visitor) {
	if nd == nil {
		return
	}

	bst.inOrder(nd.left, visitor)
	visitor(nd)
	bst.inOrder(nd.right, visitor)
}

func (bst *BST) PostOrder(visitor assist.Visitor) {
	bst.postOrder(bst.root, visitor)
} 

func (bst *BST) postOrder(nd *node, visitor assist.Visitor) {
	if nd == nil {
		return
	}

	bst.postOrder(nd.left, visitor)
	bst.postOrder(nd.right, visitor)
	visitor(nd)
}

func (bst *BST) LevelOrder(visitor assist.Visitor) {

	if bst.root == nil {
		return
	}

	var queue *linkedqueue.Queue = linkedqueue.New()

	queue.EnQueue(bst.root)

	for !bst.Empty() {

		var nd *node = queue.Head().(*node)
		queue.DeQueue()

		visitor(nd)

		if nd.left != nil {
			queue.EnQueue(nd.left)
		}

		if nd.right != nil {
			queue.EnQueue(nd.right)
		}
	}
}

func (bst *BST) Contains(v interface{}) bool {
	return bst.contains(bst.root, v)
}

func (bst *BST) contains(nd *node, v interface{}) bool {
	if nd == nil {
		return false
	}

	cmp := assist.Compare(nd.data, v, bst.comparator)
	if cmp == 0 {
		return true
	} else if cmp > 0 {
		return bst.contains(nd.left, v)
	} else {
		return bst.contains(nd.right, v)
	}
}

func (bst *BST) GetMax() interface{} {
	if bst.root == nil {
		return nil
	}

	return bst.getMaxNode(bst.root).data
}

func (bst *BST) getMaxNode(nd *node) *node {
	if nd.right == nil {
		return nd
	}

	return bst.getMaxNode(nd.right)
}

func (bst *BST) GetMin() interface{} {
	if bst.root == nil {
		return nil
	}

	return bst.getMinNode(bst.root).data
}

func (bst *BST) getMinNode(nd *node) *node {
	if nd.left == nil {
		return nd
	}

	return bst.getMinNode(nd.left)
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) Empty() bool {
	return bst.size == 0
}

// func (bst *BST) String() string {

// }

// TODO: 非递归版 前中后序遍历
// TODO: floor
// TODO: ceil
// TODO: rank
// TODO: select