package main

import (
	"fmt"
)

func main() {
	const name, age = "kim", 22

	fmt.Printf("%v is %v years old. \t and the type is %T and %T", name, age, name, age)
}
