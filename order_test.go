package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)

func TestOrderFromSlice(t *testing.T) {
	s := ds.NewSet()
	n := make([]int, 0)
	for i := 0; i < 30; i++ {
		s.Add(i)
	}
	for !s.IsEmpty() {
		n = append(n, s.Pop().(int))
	}
	
	o := ds.NewOrderFromSlice(n)
	fmt.Println(o)
}

func TestOrderSearch(t *testing.T) {
	o := ds.NewOrderFromSlice([]int{1,2,4,5})
	actual := o.Search(3)
	expected := -2
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(0)
	expected = 0
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	actual = o.Search(6)
	expected = -4
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}

func TestOrderAdd(t *testing.T) {
	o := ds.NewOrder()
	o.Add(1)
	actual := o[0]
	expected := 1
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
	o.Add(2)
	actual = o[1]
	expected = 2
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}

}
