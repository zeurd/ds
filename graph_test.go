package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)

func TestReadGraph(t *testing.T) {
	G := ds.ReadGraph("testdata/graph")
	expectedLen := 200
	if expectedLen != len(G) {
		t.Errorf("expected len: %d, actual len: %d", expectedLen, len(G))
	}
}
func TestShortestPath(t *testing.T) {
	G := ds.ReadGraph("testdata/graph")
	expectedLen := "2599,2610,2947,2052,2367,2399,2029,2442,2505,3068,"
	s := ""
	for _, e := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		d, _:= G.ShortestPath(1, e)
		s += fmt.Sprintf("%d,", d)
	}

	if expectedLen != s {
		t.Errorf("expected len: %v\nactual len: %v\n", expectedLen, s)
	}
}
