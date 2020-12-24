package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)

func TestReadGraph(t *testing.T) {
	G := ds.ReadGraph("testdata/graph")
	expectedLen := 200
	expectedCost := 18834238
	expectedEdges := 3734
	if expectedLen != G.Len() {
		t.Errorf("expected len: %d, actual len: %d", expectedLen, G.Len())
	}
	e, c := G.EdgesCost()
	if expectedCost != c {
		t.Errorf("expected cost: %d, actual cost: %d", expectedCost, c)
	}
	if expectedEdges != e {
		t.Errorf("expected number of edges: %d, actual: %d", expectedEdges, e)
	}
}
func TestShortestPath(t *testing.T) {
	G := ds.ReadGraph("testdata/graph")
	expectedLen := "2599,2610,2947,2052,2367,2399,2029,2442,2505,3068,"
	s := ""
	for _, e := range []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197} {
		d, _ := G.ShortestPath(1, e)
		s += fmt.Sprintf("%d,", d)
	}
	if expectedLen != s {
		t.Errorf("expected len: %v\nactual len: %v\n", expectedLen, s)
	}
}

// func TestReadVE(t *testing.T) {
// 	G := ds.ReadVE("testdata/ve")
// 	expectedLen := 20
// 	expectedCost := -4635
// 	expectedEdges := 10
// 	edges, cost := G.EdgesCost()
// 	len := G.Len()
// 	if len != expectedLen {
// 		t.Errorf("expectedLen: %d ; actual len: %d\n", expectedLen, len)
// 	}
// 	if edges != expectedEdges {
// 		t.Errorf("expectedEdges: %d ; actual edgest: %d\n", expectedEdges, edges)
// 	}
// 	if cost != expectedCost {
// 		t.Errorf("expectedCost: %d ; actual cost: %d\n", expectedCost, cost)
// 	}
// }
// func TestMST(t *testing.T) {
// 	G := ds.ReadVE("testdata/e")
// 	expectedCost := -10519
// 	expectedEdges := 1
// 	mst, cost := G.MST()
// 	if cost != expectedCost {
// 		t.Errorf("expectedCost: %d ; actual cost: %d\n", expectedCost, cost)
// 	}
// 	edges, cost2 := mst.EdgesCost()
// 	if edges != expectedEdges {
// 		t.Errorf("expectedEdges: %d ; actual edgest: %d\n", expectedEdges, edges)
// 	}
// 	if cost != expectedCost {
// 		t.Errorf("expectedCost: %d ; actual cost2: %d\n", expectedCost, cost2)
// 	}
// }
