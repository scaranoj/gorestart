package main

import "fmt"

//anon func are usually found within other funcs, but can also be assigned to a variable at package level scope and then called with that variable
//though careful when doing this as they are mutable, and package level scope is usually considered for immutable items

var add = func(a, b int) int { return a + b }

func main() {

	func() {
		fmt.Println("I'm an anonymous function!")
	}()

	printFunction := func() {
		fmt.Println("I'm an anonymous function assigned to a variable!")
	}

	printFunction() // Call the function using the variable name

	//calling and assigning the add func at package level scope to a var to print
	x := add(2, 2)
	fmt.Println(x)

	//declaring and assigning an anon func within another func (i.e. main())
	subtract := func(a, b int) int { return a - b }
	s := subtract(5, 2)
	fmt.Println(s)

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i is : ", i)
		}
	}()
}
