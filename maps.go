// Helper functions for maps

package gomul

// Keys returns the keys of a map in an unordered slice of the key type
// If the key type is of type cmd.Ordered you may want to sort it with slices.Sort()
func Keys[KT comparable, VT any](m map[KT]VT) []KT {
	keys := []KT{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns the values of a map in an unordered slice of the value type
// If the value type is of type cmd.Ordered you may want to sort it with slices.Sort()
func Values[KT comparable, VT any](m map[KT]VT) []VT {
	values := []VT{}
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
