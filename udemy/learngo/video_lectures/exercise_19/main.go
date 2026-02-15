package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Hello World!")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
