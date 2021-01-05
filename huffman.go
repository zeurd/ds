package ds

// NewHuffmanTree creates a  huffman tree from map of characters to their count
func NewHuffmanTree(weights map[interface{}]int) int {
	kf := func(x interface{}) int {return x.(symbol).w}
	b := NewBinarySearchTree(false, kf)
	for k, v := range weights{
		b.Push(symbol{k, v})
	}
	return b.Len()
}

type symbol struct {
	s interface{} //symbol itsel
	w int         //symbol weight
}
