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
	o := ds.NewOrder()
	o.Add(1)
	o.Add(2)
	o.Add(4)
	o.Add(5)
	fmt.Println(o)
	actual := o.Search(5)
	expected := 3
	if expected != actual {
		t.Errorf("expected: %v; actual: %v\n", expected, actual)
	}
}
