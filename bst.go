package ds

import "fmt"

// Bst is a balanced binary search tree
type Bst struct {
	r *node //root node
}

// NewBst returns a new Bst
func NewBst() *Bst {
	return &Bst{}
}

type node struct {
	k int         //key
	v interface{} //value
	p *node       //parent
	l *node       //left child
	r *node       //right child
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

func newNode(parent *node, key int, value interface{}) *node {
	return &node{
		p: parent,
		k: key,
		v: value,
	}
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
func (b *Bst) Insert(key int, value interface{}) {
	// fmt.Printf("insert %d\n", key)
	if b.r == nil {
		b.r = newNode(nil, key, value)
		return
	}
	parent, left := b.searchParent(b.r, key)
	// fmt.Printf("parent: %v\n", parent)
	child := newNode(parent, key, value)
	if left {
		parent.l = child
	} else {
		parent.r = child
	}
}

// Min returns the min element in the tree
func (b *Bst) Min() interface{} {
	min, _ := b.searchParent(b.r, -1<<32)
	return min.v
}

// Max returns the max element in the tree
func (b *Bst) Max() interface{} {
	min, _ := b.searchParent(b.r, 1<<32-1)
	return min.v
}

// keeps track of parent, to know where to insert
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
	s := make([]interface{},0)
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
	b.inOrder(n.r,s)
}
