// utilities for Set (group of unique items)
// Based on https://github.com/fatih/set

package gomul

// Set returns a slice of unique unordered items.
func Set[S ~[]E, E comparable](s S) S {
	m := make(map[E]struct{}, 0)
	for _, v := range s {
		m[v] = struct{}{}
	}
	result := make([]E, 0)
	for k := range m {
		result = append(result, k)
	}
	return result
}
