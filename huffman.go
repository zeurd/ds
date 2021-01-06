package ds

import (
	"bytes"
	"fmt"
)

//HuffmanTree foo
type HuffmanTree struct {
	root    *hnode
	encoder map[rune][]uint8
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
	kf := func(x interface{}) int { return x.(*symbol).w }
	b := NewBinarySearchTree(true, kf)
	for k, v := range weights {
		b.Push(newSymbol(k, v, false))
	}
	h := huffman{}
	t := h.build(b)
	t.mapping()
	return t
}

//Encoder foo
func Encoder(s string) *HuffmanTree {
	m := make(map[rune]int)
	for _, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r]++
		}
	}
	return NewHuffmanTree(m)
}

//Decode foo
func (h *HuffmanTree) Decode(bits []uint8) string {
	var b bytes.Buffer
	n := h.root
	for _, bit := range bits {
		n = h.next(n, bit)
		if n.v != nil {
			char := n.v.(*symbol).s.(rune)
			b.WriteRune(char)
			n = h.root
		}
	}
	return b.String()
}

func (h *HuffmanTree) next(n *hnode, b uint8) *hnode {
	if b == 0 {
		return n.z
	}
	return n.o
}

// Encode foo
func (h *HuffmanTree) Encode(s string) []uint8 {
	encoded := make([]uint8, 0)
	for _, r := range s {
		encoded = append(encoded, h.encoder[r]...)
	}
	return encoded
}

// map returns a map of characters and their binary code
func (h *HuffmanTree) mapping() {
	s := make(map[rune][]uint8)
	h.inOrder(h.root, []uint8{}, s)
	h.encoder = s
}

// Map returns a map of characters and their binary code
func (h *HuffmanTree) Map() map[rune][]uint8 {
	return h.encoder
}

func (h *HuffmanTree) String() string {
	return fmt.Sprintf("%v", h.Map())
}

func (h *HuffmanTree) inOrder(n *hnode, code []uint8, s map[rune][]uint8) {
	if n == nil {
		return
	}
	c0 := make([]uint8, len(code)+1)
	c1 := make([]uint8, len(code)+1)
	copy(c0, code)
	copy(c1, code)
	c0[len(c0)-1] = 0
	c1[len(c1)-1] = 1
	h.inOrder(n.z, c0, s)
	if n.v != nil {
		c := n.v.(*symbol).s.(rune)
		s[c] = code
	}
	h.inOrder(n.o, c1, s)
}

type symbol struct {
	s    interface{} //symbol itsel
	w    int         //symbol weight
	meta bool        //meta is true if s is the result of a merge
}

func newSymbol(s interface{}, w int, meta bool) *symbol {
	return &symbol{s, w, meta}
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
	s := root.(*symbol)
	var hroot *hnode
	if s.meta {
		hroot = s.s.(*hnode)
	} else {
		hroot = newHnode(s)
	}
	return &HuffmanTree{root: hroot}
}

func (h *huffman) greedy(b BinarySearchTree) {
	k1, low1 := b.MinK()
	b.DeleteKV(k1, low1)
	k2, low2 := b.MinK()
	b.DeleteKV(k2, low2)

	metaS := h.merge(low1.(*symbol), low2.(*symbol))
	b.Push(metaS)
}

func (h *huffman) merge(s2, s1 *symbol) *symbol {
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
