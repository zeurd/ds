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

func (n *node) String() string {
	s := fmt.Sprintf("[%d:%v]", n.k, n.v)
	if n.p != nil {
		s += fmt.Sprintf(" (p: %d)", n.p.k)
	}
	if n.l != nil {
		s += fmt.Sprintf(" (lc: %d)", n.l.k)
	}
	if n.r != nil {
		s += fmt.Sprintf(" (rc: %d)", n.r.k)
	}
	return s
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
	l int
}

func (b *Bst) String() string {
	// if b.r == nil {
	// 	return "[]"
	// }
	// if b.r.isLeaf() {
	// 	return fmt.Sprintf("[%v]", b.r.v)
	// }
	// s := "[ "
	// q :=
	// q = append(q, b.r.v)
	// i := 0
	// for {
	// 	nodeCount := len(q)
	// 	if nodeCount == 0 {
	// 		break
	// 	}
	// 	node := q[i]
	// 	s += fmt.Sprintf("%v ", node)

	// }
	cmp := func(x interface{}) int {return x.(nL).l}
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

// InsertFoo foo
func (b *Bst) InsertFoo(key int, value interface{}) {
	// fmt.Printf("insert %d\n", key)
	if b.r == nil {
		b.r = newNode(nil, key, value)
		return
	}
	parent, left := b.searchParent(b.r, key)
	// fmt.Printf("parent: %v\n", parent)
	n := newNode(parent, key, value)
	if left {
		parent.l = n
	} else {
		parent.r = n
	}
	b.rebalance(n)
}

// Insert foo
func (b *Bst) Insert(key int, value interface{}) bool {
	n := newNode(nil, key, value)
	if b.r == nil {
		b.r = n
		return true
	}
	inserted := b.insert(b.r, n)
	if inserted {
		b.rebalance(n)
	}
	return inserted
}

func (b *Bst) insert(parent, node *node) bool {
	// no duplicates
	if parent.k == node.k {
		return false
	}
	if parent.k < node.k {
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
	node := nL{n.v, level}
	q.Add(node)
	b.inLevels(n.lgit, level*2, q)
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

func (b *Bst) rebalance(n *node) {
	p := n.p
	fmt.Printf("Rebalance. n: %v p: %v\n", n, p)
	if n.l != nil && n.r != nil {
		if n.l.h > n.r.h+1 {
			b.rebalanceRight(n)
		}
		if n.r.h > n.l.h+1 {
			b.rebalanceLeft(n)
		}
	}
	if p != nil {
		b.rebalance(p)
	}
}

func (b *Bst) rebalanceRight(n *node) {
	fmt.Println("rebalance right")
	//bad case, m, problematic grandchild
	m := n.l
	if m != nil && m.r.h > m.l.h {
		b.rotateLeft(m)
	}
	b.rotateRight(n)
}

func (b *Bst) rebalanceLeft(n *node) {
	fmt.Println("rebalance left")
	m := n.r
	if m != nil && m.l.h > m.r.h {
		b.rotateRight(m)
	}
	b.rotateLeft(n)
}

func (b *Bst) rotateRight(n *node) {
	fmt.Println("rotate right")
	p := n.p
	if p.k > n.k {
		panic("rotateLeft: " + n.String())
	}
	n.p = p.p
	n.r = p
	B := n.r
	if B != nil {
		n.r = nil
		p.l = B
		b.adjustHeight(B)
	}
	b.adjustHeight(p)
	b.adjustHeight(n)

}

func (b *Bst) rotateLeft(n *node) {
	fmt.Println("rotate left")
	p := n.p
	if p.k > n.k {
		panic("rotateLeft: " + n.String())
	}
	// n keeps p's parent
	n.p = p.p
	n.l = p
	// rewire b
	B := n.l
	if B != nil {
		n.l = nil
		//check panic: p.r  should be n before that ?
		p.r = B
		b.adjustHeight(B)
	}
	b.adjustHeight(p)
	b.adjustHeight(n)
}

func (b *Bst) adjustHeight(n *node) {
	n.h = 1 + b.maxH(n.l, n.r)
}

func (b *Bst) maxH(n1 *node, n2 *node) int {
	if n1 == nil && n2 == nil {
		return 0
	}
	if n1 == nil {
		return n2.h
	}
	if n2 == nil {
		return n1.h
	}
	if n1.h > n2.h {
		return n1.h
	}
	return n2.h
}
