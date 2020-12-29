package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestBst(t *testing.T) {
	b := ds.NewBst()
	b.Insert(3, "three")
	b.Insert(5, "five")
	// b.Insert(2, "two")
	// b.Insert(1, "one")
	b.Insert(6, "six")

	s3 := b.Search(3)
	s5 := b.Search(5)
	// s2 := b.Search(2)
	// s1 := b.Search(1)
	// s6 := b.Search(6)
	if s3 != "three" {
		t.Errorf("three expected but found: %v", s3)
	}
	if s5 != "five" {
		t.Errorf("five expected but found: %v", s5)
	}
	// if s2 != "two" {
	// 	t.Errorf("two expected but found: %v", s2)
	// }
	// if s1 != "one" {
	// 	t.Errorf("one expected but found: %v", s1)
	// }
	// if s6 != "six" {
	// 	t.Errorf("six expected but found: %v", s6)
	// }

}
