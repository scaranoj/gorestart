package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	//this was the original attempt
	//forgot that you can pass in slice operators as args to an append assignment
	// y := x[:3]
	// fmt.Println(y)

	// z := x[6:]
	// fmt.Println(z)

	// x = append(y, z...)
	// fmt.Println(x)

}

/* To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
follow these steps:
● start with this slice
○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:
○ [42, 43, 44, 48, 49, 50, 51]
https://go.dev/play/p/DWR7qRa-VIR
*/
