package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestNewHuffman(t *testing.T) {
	weights := map[rune]int{'a': 60, 'b': 25, 'c': 10, 'd': 5}
	h := ds.NewHuffmanTree(weights)
	expected := map[rune]string{'a': "0", 'b': "10", 'c': "110", 'd': "111"}
	actual := h.Map()
	if len(expected) != len(actual) {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
	for k, v := range expected {
		if v != actual[k] {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}
	}
}

func TestHuffmanMaxLen(t *testing.T) {
	count, weights := ds.ReadHuffman("testdata/huffman_10_40_9_4")
	if count != len(weights) {
		t.Errorf("reading file failed. expected: %d, actual: %d", count, len(weights))
	}
	h := ds.NewHuffmanTree(weights)
	min, max := h.MinMax()
	if min != 4 {
		t.Errorf("min - expected : %v, actual: %v", 4, min)
	}
	if max != 9 {
		t.Errorf("max - expected : %v, actual: %v", 9, max)
	}
}
