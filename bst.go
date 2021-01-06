package ds

import (
	"fmt"
)

// bst is a balanced binary search tree
type bst struct {
	r   *node //root node
	len int
	kf  func(interface{}) int
}

func newBst(kf func(interface{}) int) *bst {
	return &bst{
		kf: kf,
	}
}

// Push inserts in tree, using kf() to compute the key
func (b *bst) Push(x interface{}) {
	b.Insert(b.kf(x), x)
}

type node struct {
	k int         //key
	v interface{} //value
	p *node       //parent
	l *node       //left child
	r *node       //right child
	h int         //height
}

func (n *node) String() string {
	return fmt.Sprintf("%v v:%v", n.k, n.v)
}

func newNode(parent *node, key int, value interface{}) *node {
	return &node{
		p: parent,
		k: key,
		v: value,
		h: 1,
	}
}

func (n *node) isLeaf() bool {
	return n.r == nil && n.l == nil
}

func (b *bst) String() string {
	return fmt.Sprintf("%v", b.Slice())
}

// Len returns the number of elements in the tree
func (b *bst) Len() int {
	return b.len
}

// Root returns the root value of the tree
func (b *bst) Root() interface{} {
	return b.r.v
}

// Height returns the height of the tree
func (b *bst) Height() int {
	if b.r == nil {
		return 0
	}
	return b.r.h
}

// IsValid returns true if it b is a valid search tree
func (b *bst) IsValid() bool {
	return b.valid(b.r)
}

func (b *bst) valid(n *node) bool {
	if n == nil {
		return true
	}
	l := n.l
	r := n.l
	return b.avlProperty(l, r) && b.valid(l) && b.valid(r)
}

func (b *bst) avlProperty(l, r *node) bool {
	lh := b.height(l)
	rh := b.height(r)
	diff := lh - rh
	if diff < 0 {
		diff *= -1
	}
	return diff <= 1
}

// Search foo
func (b *bst) Search(key int) interface{} {
	return b.search(b.r, key).v
}

func (b *bst) search(n *node, key int) *node {
	if n == nil {
		return nil
	}
	if key == n.k {
		return n
	}
	// recurse right
	if key > n.k {
		return b.search(n.r, key)
	}
	// recurse left
	return b.search(n.l, key)
}

// Insert foo
func (b *bst) Insert(key int, value interface{}) bool {
	n := newNode(nil, key, value)
	if b.r == nil {
		b.r = n
		b.len++
		return true
	}
	done := b.insert(b.r, n)
	if done {
		b.len++
		b.rebalance(n.p)
	}
	return done
}

func (b *bst) insert(parent, node *node) bool {
	//no duplicate in the tree
	if parent.k == node.k {
		return false
	}
	parent.h++
	if node.k > parent.k {
		if parent.r == nil {
			node.p = parent
			parent.r = node
			return true
		}
		return b.insert(parent.r, node)
	}
	// parent.k > node.k
	if parent.l == nil {
		node.p = parent
		parent.l = node

		return true
	}
	return b.insert(parent.l, node)
}

// Min returns the min element in the tree
func (b *bst) Min() interface{} {
	return b.min(b.r).v
}

func (b *bst) min(n *node) *node {
	if n.l == nil {
		return n
	}
	return b.min(n.l)
}

// Max returns the max element in the tree
func (b *bst) Max() interface{} {
	return b.max(b.r).v
}

func (b *bst) max(n *node) *node {
	if n.r == nil {
		return n
	}
	return b.max(n.r)
}

// MinK returns the min element in the tree
func (b *bst) MinK() (int, interface{}) {
	m := b.min(b.r)
	return m.k, m.v
}

// MaxK returns the max element in the tree
func (b *bst) MaxK() (int, interface{}) {
	m := b.max(b.r)
	return m.k, m.v
}

// Slice returns a sorted slice
func (b *bst) Slice() []interface{} {
	s := make([]interface{}, 0)
	b.inOrder(b.r, &s)
	return s
}

// in order traversal: left - root - right
func (b *bst) inOrder(n *node, s *[]interface{}) {
	if n == nil {
		return
	}
	b.inOrder(n.l, s)
	*s = append(*s, n.v)
	b.inOrder(n.r, s)
}

func (b *bst) deleteChildLessNode(n *node) bool {
	p := n.p
	//deleting childless root
	if p == nil {
		b.r = nil
		return true
	}
	if p.r != nil && p.r.k == n.k {
		p.r = nil
		return true
	}
	if p.l != nil && p.l.k == n.k {
		p.l = nil
		return true
	}
	return false
}

// for n with one single child
func (b *bst) spliceOut(n *node) bool {
	if n.l == nil && n.r == nil {
		panic("splice out with no child")
	}
	if n.l != nil && n.r != nil {
		panic("splice out with 2 children")
	}

	var child *node
	if n.l != nil {
		child = n.l
	} else {
		child = n.r
	}

	//promote child
	parent := n.p
	child.p = parent
	if parent == nil {
		b.r = child
		return true
	}

	nIsLeft := parent.l != nil && parent.l.k == n.k
	if nIsLeft {
		parent.l = child
	} else {
		parent.r = child
	}
	return true
}

// Delete foo
func (b *bst) Delete(key int) bool {
	n := b.search(b.r, key)
	if n == nil {
		return false
	}
	b.len--

	// 2 children => after swapping with predecessor, there is 1 or 0 child
	if n.l != nil && n.r != nil {
		m := b.predecessor(n)
		n, m = b.swap(n, m)
	}
	var deleted bool
	// 0 child
	if n.l == nil && n.r == nil {
		deleted = b.deleteChildLessNode(n)
	} else {
		// 1 child
		deleted = b.spliceOut(n)
	}
	if n.p != nil {
		b.rebalance(n.p)
	}
	return deleted
}

func (b *bst) DeleteKV(key int, x interface{}) bool {
	panic("uniomplemented")
}

func (b *bst) swap(n1, n2 *node) (*node, *node) {
	k1 := n1.k
	k2 := n2.k
	v1 := n1.v
	v2 := n2.v

	n1.k = k2
	n1.v = v2

	n2.k = k1
	n2.v = v1

	return n2, n1

}

// Predecessor returns the predecessor of the given key
func (b *bst) Predecessor(key int) interface{} {
	n := b.search(b.r, key)
	p := b.predecessor(n)
	return p.v
}

func (b *bst) predecessor(n *node) *node {
	// case 1: left non-empty, return max key in left sub-tree
	if n.l != nil {
		return b.max(n.l) //follow right side in left subtree
	}
	// case 2: follow parent until parent.k < n.k
	parent := n.p
	for parent != nil && parent.k > n.k {
		parent = parent.p
	}
	return parent
}

// if left subtree is bigger, returns a positive number; negative if right is bigger
func (b *bst) getBalance(n *node) int {
	if n == nil {
		return 0
	}
	return b.height(n.l) - b.height(n.r)
}

func (b *bst) rebalance(n *node) {
	//get balance at the parent of the new node
	balance := b.getBalance(n)

	// rebalance left side
	if balance > 1 {
		//bad
		m := n.l
		if m != nil && b.height(m.r) > b.height(m.l) {
			b.rotateLeft(m)
		}
		b.rotateRight(n)
	}
	//rebalance right side
	if balance < -1 {
		// bad case, need to make an extra rotation
		m := n.r
		if m != nil && b.height(m.l) > b.height(m.r) {
			b.rotateRight(m)
		}
		b.rotateLeft(n)
	}
	b.adjustHeight(n)
	if n.p != nil {
		b.rebalance(n.p)
	}
}

func (b *bst) adjustHeight(n *node) {
	if n == nil {
		return
	}
	h1 := b.height(n.l)
	h2 := b.height(n.r)
	if h1 > h2 {
		n.h = h1 + 1
	} else {
		n.h = h2 + 1
	}
}
func (b *bst) height(n *node) int {
	if n == nil {
		return 0
	}
	return n.h
}

func (b *bst) setParent(parent, y *node) {
	y.p = parent
	if parent == nil {
		b.r = y
		return
	}
	if parent.k > y.k {
		parent.l = y
	} else {
		parent.r = y
	}
}

func (b *bst) rotateLeft(x *node) {
	// 0. y is right child of p (right is the heavier side)
	y := x.r
	// 0. z is left child of y, that will need to be rewired
	z := y.l
	parent := x.p

	// 1. x becomes left child of y (x is smaller by def)
	y.l = x
	// 2. y is new parent and keeps p's parent
	b.setParent(parent, y)
	x.p = y
	// 3. z to be rewired to right child of x
	x.r = z
	if z != nil {
		z.p = x
	}
	// //update heights
	b.adjustHeight(x)
	b.adjustHeight(y)
	b.adjustHeight(parent)
}

func (b *bst) rotateRight(x *node) {
	//0.
	y := x.l
	z := y.r
	parent := x.p
	// 1. x becomes right child of y (x is bigger by defintion)
	y.r = x
	// 2. y keeps x's parents, x takes y as parent
	b.setParent(parent, y)
	x.p = y
	// 3. right child of y becomes left child of x
	x.l = z
	if z != nil {
		z.p = x
	}

	// //update heights
	b.adjustHeight(x)
	b.adjustHeight(y)
	b.adjustHeight(parent)
}

func (b *bst) replaceValue(k int, v interface{}) {
	n := b.search(b.r, k)
	n.v = v
}
