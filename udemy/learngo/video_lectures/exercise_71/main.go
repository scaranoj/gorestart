package main

import "fmt"

// this is the higher order func
func printSquare(f func(int) int, a int) int {
	return f(a)
}

// this is the callback func
func square(n int) int {
	return n * n
}

func main() {
	ps := printSquare(square, 2)
	fmt.Println(ps)
}

// //This program calls another func, but not as an argument (the call is in the body of the caller `add`)

// func add(a, b int) int {
// 	return adder(a, b)

// }

// func adder(c, d int) int {

// 	return c + d
// }

// func main() {

// 	sum := add(2, 3)
// 	fmt.Println(sum)

// }

//* I'm trying to use an add and adder function. From main, you would call the add function with 2 ints. The add function would use another function called "adder" as a callback func (i.e. 'adder' would do the work of adding 2 ints together and returning a sum back to the calling func add. To do this, the return of add would be the signature of the adder func.
// func adder takes a func with signature `func (int, int) int` (like the add func) as its input paramter

//What if I created a func type? A: Yep, that's what's needed here.
//What if I assigned a func to a variable and passed that way? A: not necessary, we're using func signatures
