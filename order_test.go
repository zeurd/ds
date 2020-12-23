package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestOrderFromSlice(t *testing.T) {
	n := unorderedInts(30)
	o := ds.NewOrderFromSlice(n)
	if !o.IsValid() {
		t.Errorf("Order is not valid: %v\n", o)
	}
}

func unorderedInts(max int) []int {
	s := ds.NewSet()
	n := make([]int, 0)
	for i := 0; i < max; i++ {
		s.Add(i)
	}
	for !s.IsEmpty() {
		n = append(n, s.Pop().(int))
	}
	return n
}

func TestOrderSearch(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{1, 2, 4, 15, 28})
	actual := o.Search(20)
	expected := -4
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(0)
	expected = 0
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(60)
	expected = -5
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestOrderAdd(t *testing.T) {
	l := 30000
	n := unorderedInts(l)
	o := ds.NewOrder()
	for _, x := range n {
		o.Add(x)
	}
	if !o.IsValid() {
		t.Errorf("Order is not valid: %v\n", o)
	}
	if len(o) != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, len(o))
	}
}

