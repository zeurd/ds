package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestOrderFromSlice(t *testing.T) {
	l := 10000
	n := unorderedInts(l)
	o := ds.NewOrderFromSlice(n)
	if o.Len() != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, o.Len())
	}
	if !o.IsValid() {
		t.Errorf("Order is not valid: %v\n", o)
	}
	olit := ds.NewOrderFromSlice([]int{3, 2, 1})
	if olit.Len() != 3 {
		t.Errorf("Expected (literal) length: %d. Actual: %d\n", 3, olit.Len())
	}
	if !olit.IsValid() {
		t.Errorf("Order (literal) not valid: %v\n", olit)
	}
}

func TestOrderDuplicates(t *testing.T) {
	l := 1000
	n := make([]int, l)
	for i := 0; i < l; i++ {
		n[i] = i % 10
	}
	o := ds.NewOrderFromSlice(n)
	if o.Len() != l {
		t.Errorf("Expected (literal) length: %d. Actual: %d\n", l, o.Len())
	}
	if !o.IsValid() {
		t.Errorf("Order (literal) not valid: %v\n", o)
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
	if o.Len() != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, o.Len())
	}
}

func TestOrderDelete(t *testing.T) {
	l := 10
	xs := unorderedInts(l)
	o := ds.NewOrderFromSlice(xs)
	for i, x := range xs {
		o.Delete(x)
		if !o.IsValid() {
			t.Errorf("Order not valid after deleting %d: %v\n", x, o)
		}
		if o.Len() != l-(i+1) {
			t.Errorf("Expected length: %d. Actual: %d\n", l-(i+1), o.Len())
		}
	}
}

func TestNewOrderFromInts(t *testing.T) {
	o := ds.NewOrderFromInts(2, 4, 7, 3, 0, 8, 10, 11, 7, 13, 9)
	if !o.IsValid() {
		t.Errorf("Order not valid: %v\n", o)
	}
	if o.Len() != 11 {
		t.Errorf("Expected length: %d. Actual: %d\n", 11, o.Len())
	}
	if o.Max() != 13 {
		t.Errorf("Expected max: %d. Actual: %d\n", 13, o.Max())
	}
	if o.Min() != 0 {
		t.Errorf("Expected max: %d. Actual: %d\n", 0, o.Min())
	}
}

func TestInMap(t *testing.T) {
	m := make(map[int]*ds.Order)
	m[0] = ds.NewOrderFromInts(1)
	if k, ok := m[0]; ok {
		k.Add(0)
	}
	o := m[0]
	if o.Len() != 2 {
		t.Errorf("Expected length: %d. Actual: %d\n", 2, o.Len())
	}
	if o.Max() != 1 {
		t.Errorf("Expected max: %d. Actual: %d\n", 1, o.Max())
	}
	if o.Min() != 0 {
		t.Errorf("Expected max: %d. Actual: %d\n", 0, o.Min())
	}
	if k, ok := m[1]; ok {
		k.Add(0)
	} else {
		m[1] = ds.NewOrderFromInts(5)
	}
	o = m[1]
	if o.Len() != 1 {
		t.Errorf("Expected length: %d. Actual: %d\n", 1, o.Len())
	}
	if o.Max() != 5 {
		t.Errorf("Expected max: %d. Actual: %d\n", 5, o.Max())
	}
	if o.Min() != 5 {
		t.Errorf("Expected max: %d. Actual: %d\n", 5, o.Min())
	}
}

func TestOrderSearchDuplicatesAbsent(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{0, 2, 4, 0, 2, 4, 0, 2, 4})
	actual := o.Search(-1)
	expected := -1
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(3)
	expected = -7
	if expected != actual {
		t.Errorf("search 3 -> expected: %v; actual: %v\n", expected, actual)
	}
}

func TestOrderSearchDuplicates(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{0, 2, 4, 0, 2, 4, 0, 2, 4})
	expected := 4
	actual := o.Get(o.Search(expected))
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	expected = 0
	actual = o.Get(o.Search(expected))
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	expected = 2
	actual = o.Get(o.Search(expected))
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestOrderAddDeleteDuplicates(t *testing.T) {
	o := ds.NewOrder()
	l := 1000
	n := make([]int, l)
	for i := 0; i < l; i++ {
		n[i] = i % 333
	}
	for _, x := range n {
		o.Add(x)
	}
	if o.Len() != l {
		t.Errorf("Expected length: %d. Actual: %d\n", l, o.Len())
	}
	if !o.IsValid() {
		t.Errorf("Order not valid: %v\n", o)
	}
	for i, x := range n {
		o.Delete(x)
		if !o.IsValid() {
			t.Errorf("Order not valid after deleting %d: %v\n", x, o)
		}
		if o.Len() != l-(i+1) {
			t.Errorf("Expected length: %d. Actual: %d\n", l-(i+1), o.Len())
		}
	}
}

//TODO: test cases for not !IsValid and delete in empty order
