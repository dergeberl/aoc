package utils

import "sort"

type runes []rune

func (s runes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s runes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s runes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(runes(r))
	return string(r)
}
