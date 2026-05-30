package main

import "fmt"

//closure

func main() {

	a := 0
	f := func() int {
		a++
		return a
	}
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
