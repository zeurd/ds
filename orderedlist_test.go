package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestOrderedList(t *testing.T) {
	comp := func(s interface{}) int { return len(s.(string)) }
	o := ds.NewOrderedList(comp)
	o.AddAll("bb", "ccc", "ccc", "a", "eeeee", "dddd")
	if o.Len() != 6 {
		t.Errorf("not the expected length")
	}
	if o.Min() != "a" {
		t.Errorf("not the expected min")
	}
	if o.Max() != "eeeee" {
		t.Errorf("not the expected max")
	}
}

func TestOrderedListDelete(t *testing.T) {
	comp := func(s interface{}) int { return len(s.(string)) }
	o := ds.NewOrderedList(comp)
	o.AddAll("a", "bb", "ccc", "dddd", "eeeee", "ffffff")
	i := o.Delete("")
	if i != -1 {
		t.Errorf("delete absent failed")
	}
	i = o.Delete("ccc")
	if i != 2 || o.Len() != 5 {
		t.Errorf("delete failed")
	}
}

func  TestOrderedListEdges(t *testing.T) {
	comp := func(e interface{}) int {return e.(ds.Edge).Weight()}
	o := ds.NewOrderedList(comp)
	graph, _, _, _ := ds.ReadVE("testdata/cluster_1_8_21", true)
	edges := graph.Edges()
	for e := range edges {
		o.Add(e)
	}
	if !o.IsValid() {
		t.Errorf("ordered list of edges not valid")
	}
	if o.Min().(ds.Edge).Weight() != 1 {
		t.Errorf("smallest edge not first in order: %v", o)
	}
}
