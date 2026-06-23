package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Jenny",
		Last:  "Moneypenny",
		Age:   28,
	}

	fmt.Println(p1, p2)

	people := []Person{p1, p2}

	sb, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(string(sb))
	fmt.Println(string(sb))

}
