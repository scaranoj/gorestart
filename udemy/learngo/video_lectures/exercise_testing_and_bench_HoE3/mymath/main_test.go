package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	// original single sanity test
	// xi := []int{1, 2, 3, 4}
	type test struct {
		input  []int
		output float64
	}

	tests := []test{
		{[]int{10, 11, 12, 13, 14}, 12},
		{[]int{20, 21, 22, 23, 24}, 22},
		{[]int{30, 31, 32, 33, 34}, 32},
		{[]int{40, 41, 42, 43, 44}, 42},
		{[]int{1, 2, 100, 200, 300}, 100.66666666666667},
	}

	for _, v := range tests {
		r := CenteredAvg(v.input)
		if r != v.output {
			t.Errorf("Got: %v, Want: %v", r, v.output)
		}

	}

}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5, 10000}
	fmt.Println(CenteredAvg(xi))
	//Output: 3.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for b.Loop() { //remember, don't pass anything into b.Loop
		xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000} //create the slice before the loop
		CenteredAvg(xi)                              //call it with something
	}
}

// Hints:

// 1. **Understand what the function does**: Before you can test it, you need to understand what it *returns*. Try tracing through your test case `[1, 2, 3, 4]` step by step:
//    - What does `sort.Ints()` do? //sorts a slice of ints in increasing order
//    - What does `xi[1:(len(xi) - 1)]` extract? A: it extracts a slice of a slice, starting at index `1` to it's length - 1
//    - What's the average of what's left? A: he's dividing the conversion result of float64 by the len of the resulting slice

//    Manually calculate the expected result, then verify it with your `fmt.Println()` approach.

// 2. **Table-driven test structure**: Think about what a table-driven test looks like in Go. You'll likely need:
//    - A slice of test cases (structs)
//    - Each test case containing an input and an expected output
//    - A loop that iterates through test cases
//    - An assertion inside the loop

// 3. **Float64 comparison gotcha**: When comparing floating-point numbers, direct equality (`==`) can be problematic due to floating-point precision issues. What are some ways to safely compare two `float64` values? (Hint: should they be exactly equal, or is being "close enough" acceptable?)

// 4. **Edge cases**: What happens if you pass a very small slice? A slice with one element? These might reveal whether the function handles them gracefully or panics.

// Use the **`-v` (verbose) flag**. By default, `go test` suppresses output from `fmt.Println` and only shows pass/fail. With `-v`, you see everything.

// **In VS Code:**

// 1. **Integrated Terminal** (easiest):
//    - Open the integrated terminal: `Ctrl + `` (backtick) on Windows/Linux, or `Cmd + `` on Mac
//    - Run: `go test -v`
//    - You'll see all your `fmt.Println()` output alongside the test results

//    Fixes Summarized:

// * Defined a struct for each test case, with an input slice and an expected output.
// * Created a slice of those structs as the test table, with one input/expected-result pair per case.
// * Iterated over the table with for _, v := range tests; v is the current case, so its fields are accessed as v.input and v.output.
// * Called CenteredAvg(v.input) and stored the returned value in r.
// * Compared r with the expected value, reporting a test failure with t.Errorf if they differ. This replaced printing the result, which wouldn’t fail the test.
// * Fixed the struct’s closing brace and updated the loop to range over tests rather than the old name xi.
