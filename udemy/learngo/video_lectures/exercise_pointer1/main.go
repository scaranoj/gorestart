package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	i := 69
	p := &i         //'p' is a pointer to 'i'
	fmt.Println(p)  // 'p' is a pointer, so you should receive the ADDRESS that 'p' points to (not the value)
	fmt.Println(*p) // Dereferencing 'p' gives you the integer that 'p' points to
	*p = 21         // You can change the value that 'p' points to
	fmt.Println(i)  // The value of 'i' is now 21

}
