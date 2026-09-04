package main

import "fmt"

func main() {
	// notice that the values returned are last in, first out
	// this is the behavior of using multiple defer statements in a func
	lifo()

}

func lifo() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}
