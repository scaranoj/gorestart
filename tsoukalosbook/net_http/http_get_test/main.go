package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reponse: %T\n", resp)
}


