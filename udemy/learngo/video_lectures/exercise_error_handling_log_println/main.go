package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.Open("myfile.txt")
	if err != nil {
		//same as fmt.Println() but also provides a timestamp
		log.Println("Error happened", err)
	}
	defer f.Close()
}
