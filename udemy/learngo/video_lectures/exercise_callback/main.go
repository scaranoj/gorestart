package main

import "fmt"

func main() {

	fmt.Printf("Type: %T\n", add)
	fmt.Printf("Type: %T\n", subtract)
	fmt.Printf("Type: %T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)

}

// has to have same sig as func that is passed in
// the third input parm is a anon func with the type `func(int, int) int` ..and then the final int in the sig is the return of doMath

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b

}

func subtract(a int, b int) int {
	return a - b
}
