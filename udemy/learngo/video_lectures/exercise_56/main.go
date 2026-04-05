package main

import "fmt"

func main() {

	// type person struct {
	// 	first     string
	// 	friends   map[string]int
	// 	favDrinks []string
	// }

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Horace",
		friends: map[string]int{
			"Todd":   32,
			"Bill":   26,
			"Vinnie": 41,
			"Joey":   38,
		},
		favDrinks: []string{"whiskey", "G&T", "Mezcal"},
	}
	fmt.Println(p1)

}

// first string
// 		friends map[string]int{"Todd", "Bill", "Vinnie", "Joey"}
// 		favDrinks []

/*
* Anonymous struct exercise instructions:
* Create and use an anonymous struct with these fields:
  * first string
  * friends map[string]int
  * favDrinks []string
* Print things.
* Solution: https://go.dev/play/p/l72ejdKEe3r
*/
