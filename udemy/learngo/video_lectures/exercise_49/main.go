package main

import "fmt"

func main() {
	//declare and assign a map of string keys and value of a slice of strings (`[]strings`)
	am := map[string][]string{
		`bond_james`:       []string{`shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: []string{`intelligence`, `literature`, `computer science`},
		`no_dr`:            []string{`cats`, `ice cream`, `sunsets`},
	}

	//add bogus records to map
	am[`spongebob squarepants`] = []string{`starfish`, `ocean`, `yellow`}
	am[`Q`] = []string{`weapons`, `suits`, `camoflauge`}
	am[`John Wick`] = []string{`guns`, `cat`, `hotel`}

	fmt.Println("Here is the map\n", am)

	// original loop over map
	// for k, v := range am {
	// 	fmt.Printf("Key is: %v and value is %v\n", k, v)
	// }

	// for-range over slice
	for k, v := range am {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("The index of the slice is %v and the value is %v\n", i, v2)
		}
	}

}

// Example from Bodner

// teams := map[string][]string {
//     "Orcas": []string{"Fred", "Ralph", "Bijou"},
//     "Lions": []string{"Sarah", "Peter", "Billie"},
//     "Kittens": []string{"Waldo", "Raul", "Ze"},
// }
