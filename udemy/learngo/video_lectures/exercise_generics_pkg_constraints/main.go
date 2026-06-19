//compare and contrast regular funcs from generic ones

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

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
	// go standard library packages have a "Constraint" package that you can use.
	// by indicating both here using the type set syntax, the addNums interface can include 4 different types in this
	// example, as well as their underlying types (i.e. any alias to the base types specified in the packages)
	constraints.Integer | constraints.Float
}

func addT[T addNums](g, h T) T {
	return g + h
}

type aliasInt int

func main() {

	fmt.Println(addI(2, 2))

	fmt.Println(addF(4.5, 2.5))

	fmt.Println(addG(5, 5))
	fmt.Println(addG(4.7, 5.3))

	fmt.Println(addG(2.5, 2.5))

	var n aliasInt = 42
	fmt.Println(addT(n, 58))

}
