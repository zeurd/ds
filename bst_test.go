package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestBstBasic(t *testing.T) {
	b := ds.NewBst()
	b.Insert(3, "three")
	b.Insert(5, "five")
	b.Insert(2, "two")
	b.Insert(1, "one")
	b.Insert(6, "six")

	s3 := b.Search(3)
	s5 := b.Search(5)
	s2 := b.Search(2)
	s1 := b.Search(1)
	s6 := b.Search(6)
	if s3 != "three" {
		t.Errorf("three expected but found: %v", s3)
	}
	if s5 != "five" {
		t.Errorf("five expected but found: %v", s5)
	}
	if s2 != "two" {
		t.Errorf("two expected but found: %v", s2)
	}
	if s1 != "one" {
		t.Errorf("one expected but found: %v", s1)
	}
	if s6 != "six" {
		t.Errorf("six expected but found: %v", s6)
	}

	if b.Min() != "one" {
		t.Errorf("one was the expected min but found: %v", b.Min())
	}
	if b.Max() != "six" {
		t.Errorf("one was the expected max but found: %v", b.Max())
	}

	b.Insert(4, "four")

	expectedOrder := []interface{}{"one", "two", "three", "four", "five", "six"}
	actual := b.Slice()
	if len(expectedOrder) != len(actual) {
		t.Errorf("Expected order (%v) has different length as actual (%v)", expectedOrder, actual)
	}
	for i, e := range expectedOrder {
		if e != actual[i] {
			t.Errorf("Expected order (%v) not equal actual (%v)", expectedOrder, actual)
		}
	}
	if !b.IsValid() {
		t.Errorf("BST not valid: %v", b)
	}
}

func TestBstPredecessor(t *testing.T) {
	b := ds.NewBst()
	for i := 0; i <= 10; i++ {
		b.Insert(i, i)
	}
	for i := 1; i <= 10; i++ {
		expected := i - 1
		p := b.Predecessor(i).(int)
		if p != expected {
			t.Errorf("expected predecessor: %d but got: %d", expected, p)
		}
	}
	if !b.IsValid() {
		t.Errorf("BST not valid: %v", b)
	}
	if b.Height() != 10 {
		t.Errorf("Expected height: 10; actual: %d\n%v", b.Height(), b.Slice())
	}
}

func TestBstDelete(t *testing.T) {
	b := ds.NewBstWithRoot(3, 3)
	b.Insert(1, 1)
	b.Insert(5, 5)
	b.Insert(2, 2)
	b.Insert(4, 4)

	b.Delete(3)
	// newRoot := b.Root().(int)
	// if newRoot != 2 {
	// 	t.Errorf("expected new root 2, but got: %d", newRoot)
	// }
	if !b.IsValid() {
		t.Errorf("BST not valid: %v", b)
	}
}

func TestBstRightLeft(t *testing.T) {
	b := ds.NewBst()
	xs := []int{0, 5, 1}
	for _, x := range xs {
		b.Insert(x, x)
	}
	if b.Min() != 0 {
		t.Errorf("expected min: %d but found %d", 0, b.Min())
	}
	if b.Max() != 5 {
		t.Errorf("expected max: %d but found %d", 5, b.Max())
	}
	if b.Len() != len(xs) {
		t.Errorf("expected len: %d but found %d", len(xs), b.Len())
	}
}

func TestBstLeftRight(t *testing.T) {
	b := ds.NewBst()
	xs := []int{5, 0, 1}
	for _, x := range xs {
		b.Insert(x, x)
	}
	if b.Min() != 0 {
		t.Errorf("expected min: %d but found %d", 0, b.Min())
	}
	if b.Max() != 5 {
		t.Errorf("expected max: %d but found %d", 5, b.Max())
	}
	if b.Len() != len(xs) {
		t.Errorf("expected len: %d but found %d", len(xs), b.Len())
	}
}

func TestBstFoo(t *testing.T) {
	b := ds.NewBst()
	max := 5
	xs := []int{0, 1, 2, 3, 4} //unorderedInts(max)
	for _, x := range xs {
		b.Insert(x, x)
	}
	if b.Min() != 0 {
		t.Errorf("expected min: %d but found %d", 0, b.Min())
	}
	if b.Max() != max-1 {
		t.Errorf("expected max: %d but found %d", max-1, b.Max())
	}
	if b.Len() != len(xs) {
		t.Errorf("expected len: %d but found %d", max, b.Len())
	}

	if !b.IsValid() {
		t.Errorf("BST not valid: %v", b)
	}
	//fmt.Println(b)
}
