package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("foo ran")
}

func main() {
	//use log.Fatal(), which will automatically call os.Exit()
	// notice the defer doesn't run
	defer foo()
	_, err := os.Open("fatal.txt")
	if err != nil {
		log.Fatalln(err)
	}

}

//Fatalln is the same as Println() except it also calls os.Exit(1)
