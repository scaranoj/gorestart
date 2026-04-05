package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "bob",
		friends: map[string]int{
			"freddy": 32,
			"eddy":   44,
			"carl":   55,
		},
		favDrinks: []string{"martinis", "snapple", "topo chico"},
	}

	fmt.Println(p1)

	for k, v := range p1.friends {
		fmt.Printf("friend name %v is %v years old\n", k, v)
	}

	for _, v := range p1.favDrinks {
		fmt.Println("Bob likes to drink:", v)
	}
}

/*
* Anonymous struct exercise instructions:
* Create and use an anonymous struct with these fields:
  * first string
  * friends map[string]int
  * favDrinks []string
* Print things.
* Solution: https://go.dev/play/p/l72ejdKEe3r
*/
