//compare and contrast regular funcs from generic ones

package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(c, d float64) float64 {
	return c + d
}

func addG[T int | float64](e, f T) T {
	return e + f
}

// addNums is another option to add a type set to our generics instead of defining explicitly in the generics
type addNums interface {
	int | float64
}

func addT[T addNums](g, h T) T {
	return g + h
}

func main() {

	fmt.Println(addI(2, 2))

	fmt.Println(addF(4.5, 2.5))

	fmt.Println(addG(5, 5))
	fmt.Println(addG(4.7, 5.3))

	fmt.Println(addG(2.5, 2.5))

}
