package maps

import (
	"sort"
	"strings"
)

func SortedValues(m map[int]string) string {
	keys := make([]int, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var res []string
	for _, n := range keys {
		res = append(res, m[n])
	}
	return strings.Join(res, "")
}
