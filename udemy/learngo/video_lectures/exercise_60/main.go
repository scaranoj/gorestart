package main

import "fmt"

func main() {
	defer fmt.Println("!")
	defer fmt.Println("World")
	fmt.Println("Hello")
}

/*
* “defer” multiple functions in main
  * show that a deferred func runs after the func containing it exits.
  * determine the order in which the multiple defer funcs run
  * https://go.dev/play/p/tbvSX8qy6TT
*/
