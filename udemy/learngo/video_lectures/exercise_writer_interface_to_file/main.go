package main

import (
	"log"
	"os"
)

func main() {

	// this is a function, just note the intellisense callout/popup desc doesn't show a receiver type
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)

	}
	defer f.Close()

	s := []byte("Hello Gophers")
	// this is a method, just note the intellisense callout/popup desc DOES show a receiver type
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}

}
