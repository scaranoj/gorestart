// package word provides custom functions that identify the quantity of words in strings
package word

import "strings"

// UseCount counts how many times each word is used in the quote
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// func Count(s string) int {
// 	// write the code for this func
// 	// counts how many words are in the quote
// 	xs := strings.Fields(s)
// 	c := 0
// 	for range xs {
// 		c++
// 	}
// 	return c
// }

// Count counts how many words are in the quote
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)

}
