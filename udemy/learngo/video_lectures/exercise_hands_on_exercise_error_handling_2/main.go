package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Error: Couldn't Marshal, terminating:", err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
// fmt.errorf is fine, we don't want to exit the program, just pass the error up to the callin main
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return bs, fmt.Errorf("There was an error into JSON %s", err)
	}
	return bs, nil
}
