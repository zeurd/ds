package ds_test

import (
	"testing"
	"github.com/zeurd/ds"
)

func TestNewHuffman(t *testing.T) {
	count, weights := ds.ReadHuffman("testdata/huffman_10_40_9_4")
	h := ds.NewHuffmanTree(weights)
	if h != count {
		t.Errorf("expected count: %d, actual: %d", count, h)
	}
}