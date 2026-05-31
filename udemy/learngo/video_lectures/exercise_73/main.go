package main

import "fmt"

// wrapper func
func logger(f func()) {
	fmt.Println("Before wrapper is called")
	f()
	fmt.Println("After wrapper is called")
}

func main() {
	wrappedFunc := func() {
		fmt.Println("wrapped func running now!")
	}
	logger(wrappedFunc)
}
