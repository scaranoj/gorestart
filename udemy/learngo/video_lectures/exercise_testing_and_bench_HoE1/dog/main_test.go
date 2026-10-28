package dog

import (
	"fmt"
	"testing"
)

// create test for Years
func TestYears(t *testing.T) {
	y := Years(2)
	if y != 14 {
		t.Error("Expected 14, got", y)
	}
}

//create test for YearsTwo

func TestYearsTwo(t *testing.T) {
	tc := YearsTwo(5)
	if tc != 35 {
		t.Error("Expected 35, got", tc)
	}
}

//create example functions

func ExampleYears() {
	fmt.Println(Years(3))
	// Output: 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	//Output: 28
}

// func BenchmarkYears(b *testing.B) {
// 	const y = 5
// 	for i := 0; i < b.N; i++ {
// 		Years(y)
// 	}
//}

func BenchmarkYears(b *testing.B) {
	const y = 10
	for b.Loop() { //new "b.Loop() method" https://go.dev/blog/testing-b-loop
		Years(y)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	// 	YearsTwo(b.N)
	// }
	const y = 10
	for b.Loop() { //new "b.Loop() method" https://go.dev/blog/testing-b-loop
		YearsTwo(y)
	}
}

/***

// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

***/
