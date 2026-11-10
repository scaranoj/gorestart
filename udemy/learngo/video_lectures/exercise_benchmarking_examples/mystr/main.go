package mystr

import "strings"

// we want to concat the slice we created (the one we created using the strings.Split func)
// create a named func called cat that will use a for range to add each element of the
// slice to a new string, then just print the string

func Cat(xs []string) string {
	s := ""
	for _, v := range xs {
		s += v
		s += " "
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
