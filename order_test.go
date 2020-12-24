package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestOrderFromSlice(t *testing.T) {
	l := 10
	n := unorderedInts(l)
	o := ds.NewOrderFromSlice(n)
	if len(*o) != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, len(*o))
	}
	if !o.IsValid() {
		t.Errorf("Order is not valid: %v\n", o)
	}
	olit := ds.NewOrderFromSlice([]int{3,2,1})
	if len(*olit) != 3 {
		t.Errorf("Expected (literal) length: %d. Actual: %d\n", 3, len(*olit))
	}
	if !olit.IsValid() {
		t.Errorf("Order (literal) not valid: %v\n", olit)
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
	if len(n) != max {
		panic(n)
	}
	return n
}

func TestOrderSearchAbsent(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{1, 2, 4, 15, 28})
	actual := o.Search(20)
	expected := -5
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(0)
	expected = -1
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(60)
	expected = -6
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}
func TestOrderSearch(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{1, 2, 4, 15, 28})
	actual := o.Search(1)
	expected := 0
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(15)
	expected = 3
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(28)
	expected = 4
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
	if len(*o) != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, len(*o))
	}
}

func TestOrderDelete(t *testing.T) {
	l := 3000
	n := unorderedInts(l)
	o := ds.NewOrderFromSlice(n)
	for i, x := range n {
		o.Delete(x)
		if !o.IsValid() {
			t.Errorf("Order not valid after deleting %d: %v\n", x, o)
		}
		if len(*o) != l - (i+1) {
			t.Errorf("Expected length: %d. Actual: %d\n", l-(i+1), len(*o))
		}
	}
}
