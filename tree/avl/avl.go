package avl

import (
    "../../queue/linkedqueue"
    "../../assist"
)

type node struct {
	data interface{}
	left *node
	right *node
	height int
}

func newNode(data interface{}) *node {
	return &node {
		data: data,
		left: nil,
		right: nil,
		height: 1,
	}
}

type AVL struct {
	root *node
	size int
	comparator assist.Comparator
}

func New(comparator assist.Comparator) *AVL {
	return &AVL {
		root: nil,
		size: 0,
		comparator: comparator,
	}
}

func (avl *AVL) Add(v interface{}) {
	avl.root = avl.add(avl.root, v)
}

func (avl *AVL) add(nd *node, v interface{}) *node {
	if nd == nil {
		avl.size++
		return newNode(v)
	}

	cmp := assist.Compare(nd.data, v, avl.comparator)
	if cmp > 0 {
		nd.left = avl.add(nd.left, v)
	} else if cmp < 0 {
		nd.right = avl.add(nd.right, v)
	}

	nd.height = 1 + max(avl.getHeight(nd.left), avl.getHeight(nd.right))

	var balanceFactor int = avl.getBalanceFactor(nd)

	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(nd.left) > 0 {
		return avl.rightRotate(nd)	
	}

	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(nd.left) < 0 {
		return avl.leftRotate(nd)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(nd.left) < 0 {
		nd.left = avl.leftRotate(nd.left)
		return avl.rightRotate(nd)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(nd.left) > 0 {
		nd.right = avl.rightRotate(nd.right)
		return avl.leftRotate(nd)
	}

	return nd
}

func (avl *AVL) rightRotate(nd *node) *node {

	var x *node = nd.left
	var xRight *node =x.right

	x.right = nd
	nd.left = xRight

	nd.height = max(avl.getHeight(nd.left), avl.getHeight(nd.right)) + 1
	x.height = max(avl.getHeight(x.left), avl.getHeight(x.right)) + 1

	return x
}

func (avl *AVL) leftRotate(nd *node) *node {

	var x *node = nd.right
	var xLeft *node = x.left

	x.left = nd
	nd.right = xLeft

	nd.height = max(avl.getHeight(nd.left), avl.getHeight(nd.right)) + 1
	x.height = max(avl.getHeight(x.left), avl.getHeight(x.right)) + 1

	return x
}


func (avl *AVL) Remove(v interface{}) {
	avl.root = avl.remove(avl.root, v)
}

func (avl *AVL) remove(nd *node, v interface{}) *node {

	if nd == nil {
		return nil
	}

	var ret *node

	cmp := assist.Compare(v, nd.data, avl.comparator)
	if cmp > 0 {
		nd.right = avl.remove(nd.right, v)
		ret = nd
	} else if cmp < 0 {
		nd.left = avl.remove(nd.left, v)
		ret = nd
	} else {
		if nd.left == nil {
			var rightNode *node = nd.right
			nd.right = nil
			avl.size--
			ret = rightNode
		} else if nd.right == nil {
			var leftNode *node = nd.left
			nd.left = nil
			avl.size--
			ret = leftNode
		} else {

			var successor *node = avl.getMinNode(nd.right)
			successor.right = avl.remove(nd.right, successor.data) // 已经 size-- 了
			successor.left = nd.left

			nd.left = nil
			nd.right = nil

			ret = successor
		}
	}

	if ret == nil {
		return nil
	}

	ret.height = 1 + max(avl.getHeight(ret.left), avl.getHeight(ret.right))

	var balanceFactor int = avl.getBalanceFactor(ret)

	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(ret.left) > 0 {
		return avl.rightRotate(ret)	
	}

	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(ret.left) < 0 {
		return avl.leftRotate(ret)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(ret.left) < 0 {
		ret.left = avl.leftRotate(ret.left)
		return avl.rightRotate(ret)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(ret.left) > 0 {
		ret.right = avl.rightRotate(ret.right)
		return avl.leftRotate(ret)
	}	

	return ret
}

func (avl *AVL) RemoveMin() interface{} {

	var minValue interface{} = avl.GetMin()

	avl.root = avl.remove(avl.root, minValue)

	return minValue
}

func (avl *AVL) RemoveMax() interface{} {

	var maxValue interface{} = avl.GetMax()

	avl.root = avl.removeMax(avl.root)

	return maxValue
}

func (avl *AVL) removeMax(nd *node) *node {

	if nd.right == nil {
		var leftNode *node = nd.left
		nd.left = nil
		avl.size--
		return leftNode
	}

	nd.right = avl.removeMax(nd.right)

	return nd
}

func (avl *AVL) PreOrder() []interface{} {

	var ret []interface{}

	avl.preOrder(avl.root, &ret)

	return ret
}

func (avl *AVL) preOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用指针
	avl.preOrder(nd.left, ret)
	avl.preOrder(nd.right, ret)
}

func (avl *AVL) InOrder() []interface{} {

	var ret []interface{}

	avl.inOrder(avl.root, &ret)

	return ret
} 

func (avl *AVL) inOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	avl.inOrder(nd.left, ret)
	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用指针
	avl.inOrder(nd.right, ret)
}

func (avl *AVL) PostOrder() []interface{} {

	var ret []interface{}

	avl.postOrder(avl.root, &ret)

	return ret
} 

func (avl *AVL) postOrder(nd *node, ret *[]interface{}) {
	if nd == nil {
		return
	}

	avl.postOrder(nd.left, ret)
	avl.postOrder(nd.right, ret)
	*ret = append(*ret, nd.data) // 因为 append 可能使形参指向新的 切片，所以必须要用指针
}

func (avl *AVL) LevelOrder() []interface{} {

	var ret []interface{}

	if avl.root == nil {
		return ret
	}

	var queue *linkedqueue.Queue = linkedqueue.New()

	queue.EnQueue(avl.root)

	for !queue.Empty() {

		var nd *node = queue.Head().(*node)
		queue.DeQueue()

		ret = append(ret, nd.data)

		if nd.left != nil {
			queue.EnQueue(nd.left)
		}

		if nd.right != nil {
			queue.EnQueue(nd.right)
		}
	}

	return ret
}

func (avl *AVL) Contains(v interface{}) bool {
	return avl.contains(avl.root, v)
}

func (avl *AVL) contains(nd *node, v interface{}) bool {
	if nd == nil {
		return false
	}

	cmp := assist.Compare(nd.data, v, avl.comparator)
	if cmp == 0 {
		return true
	} else if cmp > 0 {
		return avl.contains(nd.left, v)
	} else {
		return avl.contains(nd.right, v)
	}
}

func (avl *AVL) GetMax() interface{} {
	if avl.root == nil {
		return nil
	}

	return avl.getMaxNode(avl.root).data
}

func (avl *AVL) getMaxNode(nd *node) *node {
	if nd.right == nil {
		return nd
	}

	return avl.getMaxNode(nd.right)
}

func (avl *AVL) GetMin() interface{} {
	if avl.root == nil {
		return nil
	}

	return avl.getMinNode(avl.root).data
}

func (avl *AVL) getMinNode(nd *node) *node {
	if nd.left == nil {
		return nd
	}

	return avl.getMinNode(nd.left)
}

func (avl *AVL) Size() int {
	return avl.size
}

func (avl *AVL) Empty() bool {
	return avl.size == 0
}

func (avl *AVL) isBST() bool {

	var keys []interface{} = avl.InOrder()

	for i := 1; i < len(keys); i++ {
		if assist.Compare(keys[i-1], keys[i], avl.comparator) > 0 {
			return false
		}
	}

	return true
}

func (avl *AVL) isBalance() bool {
	return avl.isBalanceRecursion(avl.root)
}

func (avl *AVL) isBalanceRecursion(nd *node) bool {
	if nd == nil {
		return true
	}

	var balanceFactor int = avl.getBalanceFactor(nd)

	if balanceFactor > 1 || balanceFactor < -1 {
		return false
	}

	return avl.isBalanceRecursion(nd.left) && avl.isBalanceRecursion(nd.right)
}

func (avl *AVL) getHeight(nd *node) int {
	if nd == nil {
		return 0
	}

	return nd.height
}

func (avl *AVL) getBalanceFactor(nd *node) int {
	if nd == nil {
		return 0
	}
	return avl.getHeight(nd.left) - avl.getHeight(nd.right)
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
} 

// TODO: 非递归版 前中后序遍历
// TODO: floor
// TODO: ceil
// TODO: rank
// TODO: select