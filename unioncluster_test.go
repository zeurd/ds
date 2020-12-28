package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestClusters(t *testing.T) {
	c := ds.NewClusters()
	graph, m, _, _ := ds.ReadVE("testdata/cluster_1_8_21", true)
	var v interface{}
	for node := range graph.Nodes() {
		added := c.Add(node)
		v = node
		if !added {
			t.Errorf("%d was not added\n", v)
		}
		added = c.Add(v)
		if added {
			t.Errorf("%d was added twice \n", v)
		}
	}
	if m != c.Count() {
		t.Errorf("wrong count of clusters: %d vs %d\n", m, c.Count())
	}
	c.Union(1, 2)
	c.Union(1, 3)
	if m-2 != c.Count() {
		t.Errorf("wrong count of clusters after 2 unions: %d vs %d\n", m-2, c.Count())
	}
	if !c.Connected(2, 1) {
		t.Errorf("2 and 3 not connected: %d, %d,\n%v\n", c.Find(2), c.Find(3), c)
	}
	c.Union(7, 8)
	c.Union(8, 6)
	c.Union(3, 4)
	c.Union(2, 5)
	if c.Count() != 2 {
		t.Errorf("should have 2 clusters: %v\n", c)
	}
	c.Union(6, 5)
	if c.Count() != 1 {
		t.Errorf("should have 1 clusters: %v\n", c)
	}
	s := c.Find(v)
	for node := range graph.Nodes() {
		if s != c.Find(node) {
			t.Errorf("%d and %d are not in the same cluster: %d and %d\n%v", v, node, s, c.Find(node), c)
		}
	}
}
