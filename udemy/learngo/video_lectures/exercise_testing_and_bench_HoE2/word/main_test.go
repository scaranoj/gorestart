// word identifies the quantity of words in strings
package word

import (
	"fmt"
	"maps"
	"testing"
)

func TestCount(t *testing.T) {
	s := "You can choose freewill"
	i := Count(s)
	if i != 4 {
		t.Error("expected 3, got", i)
	}
}

func TestUseCount(t *testing.T) {
	//unnecessary, the func returns one
	// m := make(map[string]int)
	//var m map[string]int
	s := "all the beautiful people" //provide a test string
	m := UseCount(s)                //call UseCount and pass it the test string and declare/assign result to a map
	// for k, v := range UseCount("all the beautiful people") {
	// if m != map[string]int{"all": 1, "the": 1, "beautiful": 1, "people": 1} {
	if !maps.Equal(m, map[string]int{"all": 1, "the": 1, "beautiful": 1, "people": 1}) {
		t.Error("expected map[string]int{all: 1, the: 1, beautiful: 1, people: 1}, got", m)
	}
}

// // UseCount counts how many times each word is used in the quote
// func UseCount(s string) map[string]int {
// 	xs := strings.Fields(s)
// 	m := make(map[string]int)
// 	for _, v := range xs {
// 		m[v]++
// 	}
// 	return m
// }

func ExampleCount() {
	s := "just say yes"
	fmt.Println(Count(s))
	//Output: 3
}

func BenchmarkCount(b *testing.B) {
	s := "if you choose not to decide you've still made a choice"
	for b.Loop() {
		Count(s)
	}
}

func BenchmarkUseCount(b *testing.B) {
	s := "the answers are coming neo"
	for b.Loop() {
		UseCount(s)
	}
}

// import "strings"

// // no need to write an example for this one
// // writing a test for this one is a bonus challenge; harder
// // UseCount counts how many times each word is used in the quote
// func UseCount(s string) map[string]int {
// 	xs := strings.Fields(s)
// 	m := make(map[string]int)
// 	for _, v := range xs {
// 		m[v]++
// 	}
// 	return m
// }

// // func Count(s string) int {
// // 	// write the code for this func
// // 	// counts how many words are in the quote
// // 	xs := strings.Fields(s)
// // 	c := 0
// // 	for range xs {
// // 		c++
// // 	}
// // 	return c
// // }

// // Count counts how many words are in the quote
// func Count(s string) int {
// 	xs := strings.Fields(s)
// 	return len(xs)

// }
