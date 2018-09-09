package bst

import (
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

func (bst *BST) RemoveMin() interface{} {

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

func (bst *BST) RemoveMax() interface{} {

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

func (bst *BST) PreOrder(ret *[]interface{}) {

	bst.preOrder(bst.root, ret)
}

func (bst *BST) preOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用切片
	bst.preOrder(nd.left, ret)
	bst.preOrder(nd.right, ret)
}

func (bst *BST) InOrder(ret *[]interface{}) {

	bst.inOrder(bst.root, ret)
} 

func (bst *BST) inOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	bst.inOrder(nd.left, ret)
	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用切片
	bst.inOrder(nd.right, ret)
}

func (bst *BST) PostOrder(ret *[]interface{}) {
	bst.postOrder(bst.root, ret)
} 

func (bst *BST) postOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	bst.postOrder(nd.left, ret)
	bst.postOrder(nd.right, ret)
	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用切片
}

func (bst *BST) LevelOrder(ret *[]interface{}) {

	if bst.root == nil {
		return
	}

	var queue *linkedqueue.Queue = linkedqueue.New()

	queue.EnQueue(bst.root)

	for !queue.Empty() {

		var nd *node = queue.Head().(*node)
		queue.DeQueue()

		*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用切片

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

// 	var stringBuilder []rune

// 	bst.rstring(bst.root, &stringBuilder)

// 	return string(stringBuilder)
// }

// func (bst *BST) rstring(nd *node, sb *[]rune) {

// 	if nd == nil {
// 		return
// 	}

// 	*sb = append(*sb, fmt.Sprintnd.data)
// }

// TODO: 非递归版 前中后序遍历
// TODO: floor
// TODO: ceil
// TODO: rank
// TODO: select