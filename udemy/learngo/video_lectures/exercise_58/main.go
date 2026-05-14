package main

import "fmt"

func main() {

	f := foo()
	fmt.Println(f)

	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 24, "Bar ran"
}

/*

package main

v1

import "fmt"

func foo() int {
	return 5
}

func bar() (int, string) {
	return 10, "bar ran"

}

func main() {

	fmt.Println(foo())
	fmt.Println(bar())

}

/*Hands on exercise
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
https://go.dev/play/p/Jeh_wripOEe
*/
