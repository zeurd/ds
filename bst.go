package ds

import (
	"fmt"
)

// Bst is a balanced binary search tree
type Bst struct {
	r *node //root node
}

// NewBst returns a new Bst
func NewBst() *Bst {
	return &Bst{}
}

// NewBstWithRoot foo
func NewBstWithRoot(key int, value interface{}) *Bst {
	return &Bst{
		newNode(nil, key, value),
	}
}

type node struct {
	k int         //key
	v interface{} //value
	p *node       //parent
	l *node       //left child
	r *node       //right child
	h int         //height
}

// meta, return int (the number of children), bool (has left child), bool (is left child)
func (n *node) meta() (int, bool, bool) {
	c := 0
	hasLeft := false
	isLeft := false
	if n.l != nil {
		c++
		hasLeft = true
	}
	if n.r != nil {
		c++
	}
	if n.p != nil && n.p.l == n {
		isLeft = true
	}
	return c, hasLeft, isLeft
}

func (n *node) replaceChild(r *node, left bool) {
	if left {
		n.l = r
	} else {
		n.r = r
	}
}

func (n *node) getOneChild(left bool) *node {
	if left {
		return n.l
	}
	return n.r
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

type nL struct {
	x interface{}
	h int
	l int
}

func (b *Bst) String() string {
	cmp := func(x interface{}) int { return x.(nL).l }
	q := NewOrderedList(cmp)
	b.inLevels(b.r, 1, q)
	return q.String()

}

// Root returns the root value of the tree
func (b *Bst) Root() interface{} {
	return b.r.v
}

// Height returns the height of the tree
func (b *Bst) Height() int {
	return b.r.h
}

// IsValid returns true if it b is a valid search tree
func (b *Bst) IsValid() bool {
	return b.valid(b.r)
}

func (b *Bst) valid(n *node) bool {
	if n == nil {
		return true
	}
	l := n.l
	r := n.l
	return b.avlProperty(l, r) && b.valid(l) && b.valid(r)
}

func (b *Bst) avlProperty(l, r *node) bool {
	lh := 0
	rh := 0
	if l != nil {
		lh = l.h
	}
	if r != nil {
		rh = r.h
	}
	diff := lh - rh
	if diff < 0 {
		diff *= -1
	}
	return diff <= 1
}

// Search foo
func (b *Bst) Search(key int) interface{} {
	return b.search(b.r, key).v
}

func (b *Bst) search(n *node, key int) *node {
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
func (b *Bst) Insert(key int, value interface{}) bool {
	n := newNode(nil, key, value)
	if b.r == nil {
		b.r = n
		return true
	}
	done := b.insert(b.r, n)
	n.h = b.max(n.r, n.l) + 1
	if done {
		b.rebalance(n)
	}
	return done
}

func (b *Bst) insert(parent, node *node) bool {
	// no duplicates
	if parent.k == node.k {
		return false
	}
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
func (b *Bst) Min() interface{} {
	min, _ := b.searchParent(b.r, -1<<32)
	return min.v
}

// Max returns the max element in the tree
func (b *Bst) Max() interface{} {
	max, _ := b.searchParent(b.r, 1<<32-1)
	return max.v
}

// returns the parent where to insert
func (b *Bst) searchParent(p *node, key int) (*node, bool) {
	if p.k == key {
		panic("duplicate")
	}
	// left
	if key < p.k {
		if p.l == nil {
			return p, true
		}
		return b.searchParent(p.l, key)
	}
	// right
	if p.r == nil {
		return p, false
	}
	return b.searchParent(p.r, key)
}

// Slice returns a sorted slice
func (b *Bst) Slice() []interface{} {
	s := make([]interface{}, 0)
	b.inOrder(b.r, &s)
	return s
}

// in order traversal: right - root - left
func (b *Bst) inOrder(n *node, s *[]interface{}) {
	if n == nil {
		return
	}
	b.inOrder(n.l, s)
	*s = append(*s, n.v)
	b.inOrder(n.r, s)
}

// in levels traversal
func (b *Bst) inLevels(n *node, level int, q *OrderedList) {
	if n == nil {
		return
	}
	node := nL{n.v, n.h, level}
	q.Add(node)
	b.inLevels(n.l, level*2, q)
	b.inLevels(n.r, level*2+1, q)
}

// Delete foo
func (b *Bst) Delete(key int) {
	n := b.search(b.r, key)
	nC, hasLeft, isLeft := n.meta()
	parent := n.p

	// case 0: no child -> delete
	if nC == 0 {
		parent.replaceChild(nil, isLeft)
	} else if nC == 1 {
		//case 1: one child -> splice out
		parent.replaceChild(n.getOneChild(hasLeft), isLeft)
	} else {
		//case 2: 2 children
		pred := b.predecessor(n) // by defintion, predecessor is the last right child in its tree (is Right and has no left)
		b.swap(n, pred)          //n is now pred
		potentialLeft := pred.getOneChild(true)
		pred.p.replaceChild(potentialLeft, false)
	}
	b.rebalance(n)
}

// Predecessor returns the predecessor of the given key
func (b *Bst) Predecessor(key int) interface{} {
	n := b.search(b.r, key)
	p := b.predecessor(n)
	return p.v
}

func (b *Bst) predecessor(n *node) *node {
	// case 1: left non-empty, return max key in left sub-tree
	if n.l != nil {
		maxL, _ := b.searchParent(n.l, 1<<32-1) //follow right side in left subtree
		return maxL
	}
	// case 2: follow parent until parent.k < n.k
	parent := n.p
	for parent != nil && parent.k > n.k {
		parent = parent.p
	}
	return parent
}

func (b *Bst) swap(n1, n2 *node) {
	n1.k, n2.k = n2.k, n1.k
	n1.v, n2.v = n2.v, n1.v
}

// if left subtree is bigger, returns a positive number; negative if right is bigger
func (b *Bst) getBalance(n *node) int {
	if n == nil {
		return 0
	}
	return b.height(n.l) - b.height(n.r)
}

func (b *Bst) height(n *node) int {
	if n == nil {
		return 0
	}
	return n.h
}

func (b *Bst) max(n1, n2 *node) int {
	h1 := b.height(n1)
	h2 := b.height(n2)
	if h1 > h2 {
		return h1
	}
	return h2
}

func (b *Bst) rebalance(n *node) {

	//get balance at the parent of the new node
	balance := b.getBalance(n)
	fmt.Printf("Rebalance. n: %v balance: %v\n", n.v, balance)

	// left is bigger; rebalance left
	if balance > 1 {
		fmt.Println("\tleft")
		if n.k > n.l.k {
			//left right case
			fmt.Println("\tleft right")
			b.rotateLeft(n.l)
		}
		b.rotateRight(n)
	}
	//right is bigger; rebalance right
	if balance < -1 {
		fmt.Println("\tright")
		// right left case
		// if n.k < n.r.k {
		// 	fmt.Println("\tright left")
		// 	b.rotateRight(n.r)
		// }
		b.rotateLeft(n)

	}
	n.h = 1 + b.max(n.r, n.l)
	if n.p != nil {
		b.rebalance(n.p)
	}
}

func (b *Bst) rotateLeft(x *node) {
	y := x.r
	T2 := y.l

	//rotation
	y.l = x
	x.r = T2
	//update heights
	x.h = 1 + b.max(x.l, x.r)
	y.h = 1 + b.max(y.l, y.r)

	x.l = y
}

func (b *Bst) rotateRight(y *node) {
	x := y.l
	T2 := x.r
	x.r = y
	y.l = T2
	y.h = 1 + b.max(y.l, y.r)

	y.r = x
}
