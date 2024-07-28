package gomul

import (
	"slices"
	"testing"
)

func TestSet(t *testing.T) {
	s := []int{2, 2, 3, 1, 4, 8, 5, 6, 0, 7, 1, 1, 5, 9, 2}
	u := Set(s)
	slices.Sort(u)
	r := slices.Compare(u, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	if r != 0 {
		t.Errorf("The resulting slice not a set with unique items, %v", u)
	}
}
