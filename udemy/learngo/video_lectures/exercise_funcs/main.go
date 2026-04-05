package main

import "fmt"

func main() {

	hello()
	hola("Yayson")
	s1 := aloha("Mr. Jason")
	fmt.Println(s1)
	ha, age := helloAge("Jason", 35)
	fmt.Println(ha, age)

}

//the syntax of a function in Go is:
//func (receiver) identifier (parameters) (returns) {...}

//write a simple function that takes no args and has no returns and call it

func hello() {
	fmt.Println("Hello, my name is Jason.")
}

//write a func that takes one argument but doesn't return anything

func hola(s string) {
	fmt.Println("Hola, ", s)
}

//write a function that takes one arg and returns one arg

func aloha(s string) string {
	return fmt.Sprint("Aloha, ", s)
}

func helloAge(name string, age int) (string, int) {
	return fmt.Sprint(name, " is"), age
}
