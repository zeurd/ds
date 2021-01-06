package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)

func TestNewHuffman(t *testing.T) {
	weights := map[interface{}]int{"a": 60, "b": 25, "c": 10, "d": 5}
	h := ds.NewHuffmanTree(weights)
	fmt.Println(h)
}
