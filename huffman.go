package ds

import "fmt"

//HuffmanTree foo
type HuffmanTree struct {
	root *hnode
}

type huffman struct{}

type hnode struct {
	p *hnode      //parent
	z *hnode      //zero
	o *hnode      //one
	v interface{} //value
}

func newHnode(v interface{}) *hnode {
	return &hnode{v: v}
}

// NewHuffmanTree creates a  huffman tree from map of characters to their count
func NewHuffmanTree(weights map[rune]int) *HuffmanTree {
	kf := func(x interface{}) int { return x.(symbol).w }
	b := NewBinarySearchTree(true, kf)
	for k, v := range weights {
		b.Push(newSymbol(k, v, false))
	}
	h := huffman{}
	return h.build(b)
}

// Map returns a map of characters and their binary code
func (h *HuffmanTree) Map() map[rune]string {
	s := make(map[rune]string)
	h.inOrder(h.root, "", s)
	return s
}

func (h *HuffmanTree) String() string {
	return fmt.Sprintf("%v", h.Map())
}

func (h *HuffmanTree) inOrder(n *hnode, code string, s map[rune]string) {
	if n == nil {
		return
	}
	h.inOrder(n.z, code+"0", s)
	if n.v != nil {
		c := n.v.(symbol).s.(rune)
		s[c] = code
	}
	h.inOrder(n.o, code+"1", s)
}

type symbol struct {
	s    interface{} //symbol itsel
	w    int         //symbol weight
	meta bool        //meta is true if s is the result of a merge
}

func newSymbol(s interface{}, w int, meta bool) symbol {
	return symbol{s, w, meta}
}

func (h *huffman) build(b BinarySearchTree) *HuffmanTree {
	//base case
	if b.Len() <= 1 {
		return h.tree(b)
	}
	h.greedy(b)
	return h.build(b)
}

func (h *huffman) tree(b BinarySearchTree) *HuffmanTree {
	root := b.Min()
	s := root.(symbol)
	var hroot *hnode
	if s.meta {
		hroot = s.s.(*hnode)
	} else {
		hroot = newHnode(s)
	}
	return &HuffmanTree{hroot}
}

func (h *huffman) greedy(b BinarySearchTree) {
	k1, low1 := b.MinK()
	b.DeleteKV(k1, low1)
	k2, low2 := b.MinK()
	b.DeleteKV(k2, low2)

	metaS := h.merge(low1.(symbol), low2.(symbol))
	b.Push(metaS)
}

func (h *huffman) merge(s2, s1 symbol) symbol {
	n := newHnode(nil)
	var l *hnode
	if s1.meta {
		l = s1.s.(*hnode)
	} else {
		l = newHnode(s1)
	}
	var r *hnode
	if s2.meta {
		r = s2.s.(*hnode)
	} else {
		r = newHnode(s2)
	}
	n.z = l
	n.o = r
	return newSymbol(n, s1.w+s2.w, true)
}

//MinMax foo
func (h *HuffmanTree) MinMax() (int, int) {
	min := 1000_000
	max := -1
	m := h.Map()
	for _, v := range m {
		l := len(v)
		if l < min {
			min = l
		}
		if l > max {
			max = l
		}
	}
	return min, max
}
