package util

import "testing"

func TestSortList(t *testing.T) {
	t.Fatal(SortList([]map[string]int{map[string]int{"a": 1, "b": 4}, map[string]int{"a": 12, "b": 43}, map[string]int{"a": 31, "b": 14}}, "b", true))
}
