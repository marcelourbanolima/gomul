package gomul

import (
	"slices"
	"testing"
)

func TestKeysAndValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"e": 5,
		"b": 2,
		"d": 4,
	}

	k := Keys(m)
	slices.Sort(k)
	ck := slices.Compare(k, []string{"a", "b", "c", "d", "e"})
	if ck != 0 {
		t.Errorf("The key slices are not equal: %v", k)
	}

	v := Values(m)
	slices.Sort(v)
	cv := slices.Compare(v, []int{1, 2, 3, 4, 5})
	if cv != 0 {
		t.Errorf("The value slices are not equal: %v", v)
	}
}
