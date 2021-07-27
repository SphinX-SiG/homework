package slices

import "sort"

func SortStringsInPlace(s []string) {
	sort.Strings(s)
}

func SortStringsPure(s []string) []string {
	n := make([]string, len(s))
	copy(n, s)
	SortStringsInPlace(n)
	return n
}

type User struct {
	FirstName string
	LastName  string
}

func SortUsersPure(s []User) []User {
	n := make([]User, len(s))
	copy(n, s)

	sort.Slice(n, func(i, j int) bool {
		if n[i].FirstName == n[j].FirstName {
			return n[i].LastName < n[j].LastName
		}
		return n[i].FirstName < n[j].FirstName
	})
	return n
}

func SortUsersPureDesc(s []User) []User {
	n := make([]User, len(s))
	copy(n, s)

	sort.Slice(n, func(i, j int) bool {
		if n[i].FirstName == n[j].FirstName {
			return n[i].LastName > n[j].LastName
		}
		return n[i].FirstName > n[j].FirstName
	})
	return n
}
