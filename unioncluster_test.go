package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestClusters(t *testing.T) {
	c := ds.NewClusters()
	graph, m, _, _ := ds.ReadVE("testdata/cluster_1_8_21", true)
	for node := range graph.Nodes() {
		c.Add(node)
	}
	if m != c.Count() {
		t.Errorf("wrong count of clusters: %d vs %d\n", m, c.Count())
	}
	c.Union(1, 2)
	//c.Union(1, 3)
	if m-1 != c.Count() {
		t.Errorf("wrong count of clusters after 2 unions: %d vs %d\n", m-2, c.Count())
	}
	if !c.Connected(2,1) {
		t.Errorf("2 and 3 not connected: %d, %d,\n%v\n", c.Find(2), c.Find(3), c)
	}
}
