package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)
func intHeap() ds.Heap {
	h := ds.NewHeap()
	h.Push(1)
	h.Push(2)
	h.Push(50)
	h.Push(3)
	h.Push(4)
	h.Push(51)
	h.Push(52)
	h.Push(5)
	h.Push(6)
	h.Push(7)
	h.Push(8)
	return h
}
func TestHeapInsert(t *testing.T) {
	h := intHeap()
	expected := "[1 2 50 3 4 51 52 5 6 7 8]"
	actual := h.String()
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestHeapUpdate(t *testing.T) {
	h := intHeap()
	h.Update(4, 2)
	h.Update(6, 0)
	h.Update(5, 1)
	h.Update(51, 8)
	expected := "[[6] [1 2] [4 5 50 52] [3 7 8 51]]"
	actual := h.String()
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestHeapDelete(t *testing.T) {
	h := intHeap()
	h.Delete(2)
	h.Delete(5)
	h.Delete(8)
	h.Delete(50)
	expected := "[[1] [3 7] [6 4 51 52]]"
	actual := h.String()
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestHeapPop(t *testing.T) {
	h := ds.NewHeap()
	s := ds.NewSet()
	n := 30
	for i := 0; i <= n; i++ {
		s.Add(i)
	}
	for i := 0; i <= n; i++ {
		h.Push(s.Pop())
	}
	for i := 0; i <= n; i++ {
		x := h.Pop()
		fmt.Println(x)
	}
	expected := "[]"
	actual := h.String()
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestHeapSubtree(t *testing.T) {
	n := 30
	h := ds.NewHeap()
	for i := 0; i <= n; i++ {
		h.Push(i)
	}
	actual := make([]int, 0)
	expected := "[30 22 30 18 22 26 30 16 18 20 22 24 26 28 30 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]"
	for i := 0; i <= n; i++ {
		s := h.Subtree(i)
		ts := len(s) - 1
		i := len(s[ts]) - 1
		actual = append(actual, s[ts][i].(int))
	}
	if fmt.Sprintf("%v", actual) != expected {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestHeapEdge(t *testing.T) {
	g := ds.ReadVE("testdata/ve")
	VX := ds.NewHeap()
	for edge := range g.Edges() {
		e := edge.(ds.Edge)
		VX.Insert(e.To(), e.Weight())
	}
}