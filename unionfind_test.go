package ds_test

import (
	"fmt"
	"testing"

	"github.com/zeurd/ds"
)

func TestUF(t *testing.T) {
	u := ds.NewUnionFind()
	u.Add(1)
	u.Add(2)
	u.Add(3)
	if u.Rank(1) != 0 || u.Rank(1) != u.Rank(2) || u.Rank(2) != u.Rank(3) {
		t.Errorf("one added element does not have rank 0: r1(%d), r2(%d), r3(%d)", u.Rank(1), u.Rank(2), u.Rank(3))
	}
	u.Union(1, 2)
	s1 := u.Find(1)
	s2 := u.Find(2)
	s3 := u.Find(3)
	fmt.Println("after")
	u.Find(1)

	if s1 != s2 {
		t.Errorf("union failed (s1 != s2):  s1 (%d), s2 (%d)\n", s1, s2)
	}
	if !u.Connected(1, 2) {
		t.Errorf("union failed (1 and 2 not connected): s1 (%d), s2 (%d)\n", s1, s2)
	}
	if u.Connected(s3, s1) || u.Connected(s2, s3) {
		t.Errorf("s3 (%d) connected to s1 (%d) or s2 (%d)", s3, s1, s2)
	}

	r1 := u.Rank(1)
	r2 := u.Rank(2)
	if r1-r2 != 1 && r2-r1 != 1 {
		t.Errorf("r1 or r2 should be have been incremented: r1(%d), r2(%d)", r1, r2)
	}
}

func TestUFUnion(t *testing.T) {
	u := ds.NewUnionFind()
	for i := 1; i <= 9; i++ {
		added := u.Add(i)
		if !added {
			t.Errorf("%d was not added\n", i)
		}
		added = u.Add(i)
		if added {
			t.Errorf("%d was added twice \n", i)
		}
	}
	u.Union(1, 5)
	s1 := u.Find(1)
	s5 := u.Find(5)
	if s1 != s5 {
		t.Errorf("union failed (s1 != s2):  s1 (%d), s2 (%d)\n", s1, s5)
	}
	if !u.Connected(1, 5) {
		t.Errorf("union failed (1 and 2 not connected): s1 (%d), s2 (%d)\n", s1, s5)
	}
	u.Union(1, 9)
	s1 = u.Find(1)
	s9 := u.Find(9)
	merged := u.Union(9, 5)
	if merged {
		t.Errorf("5 and 9 should already be in same group")
	}
	if s1 != s9 {
		t.Errorf("1 and 9 not merged but in groups: %d and %d\n", s1, s9)
	}
}
