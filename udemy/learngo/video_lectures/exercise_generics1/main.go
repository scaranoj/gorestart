package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(c, d float64) float64 {
	return c + d
}

func addT[T int | float64](e, f T) T {
	return e + f
}

func main() {

	fmt.Println(addI(2, 2))
	fmt.Println(addF(2.5, 2.5))
	fmt.Println(addT(2, 2))
	fmt.Println(addT(2.5, 2.5))
}
