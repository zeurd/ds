package ds_test

import (
	"testing"

	"github.com/zeurd/ds"
)

func TestCombi(t *testing.T) {
	c := ds.NewCombinations(25, 7)
	actual := c.Size()
	expected := 480700
	if actual != expected {
		t.Errorf("expected: %v. actual: %v", expected, actual)
	}
}
