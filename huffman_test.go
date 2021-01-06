package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestNewHuffman(t *testing.T) {
	weights := map[rune]int{'a': 60, 'b': 25, 'c': 10, 'd': 5}
	h := ds.NewHuffmanTree(weights)
	expected := map[rune][]uint8{'a': {0}, 'b': {1, 0}, 'c': {1, 1, 0}, 'd': {1, 1, 1}}
	actual := h.Map()
	if len(expected) != len(actual) {
		t.Errorf("actual: %v, expected: %v", actual, expected)
	}
	for k, actualbits := range expected {
		for i, bit := range actualbits {
			if bit != actual[k][i] {
				t.Errorf("actual: %v, expected: %v", actual, expected)
			}

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

func TestHuffmanEncoder(t *testing.T) {
	text := "aaaaaaaaaaaabbbbbccd"
	e := ds.Encoder(text)
	actual := e.Encode("abcd")
	expected := []uint8{0, 1, 0, 1, 1, 0, 1, 1, 1}
	for i, b := range expected {
		if actual[i] != b {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}
	}
	e.Decode(expected)
}
