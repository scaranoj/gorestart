package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//notice defer will run before the func exits. Can also use a recover here if want
	defer foo()
	_, err := os.Open("test.txt")
	if err != nil {
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("Foo ran")
}

//log.Panicln() is equivalent to fmt.Println() followed by a call to panic()
