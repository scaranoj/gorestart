package main

import (
	"fmt"
)

func main() {

	type person struct {
		first  string
		last   string
		flavor []string
	}

	p1 := person{
		first:  "jason",
		last:   "johnson",
		flavor: []string{"strawberry", "vanilla", "chocolate"},
	}

	p2 := person{
		first:  "gary",
		last:   "larson",
		flavor: []string{"rocky road", "salted butterscotch", "peanut butter"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.flavor {
			fmt.Println(v.first, v.last, v2)
		}
	}

	// for _, v := range p1.flavor {
	// 	fmt.Println(p1.first, "likes", v)
	// }
	// for _, v := range p2.flavor {
	// 	fmt.Println(p2.first, "likes", v)
	// }

	// fmt.Println(p1)
	// fmt.Println(p2)

}

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
https://go.dev/play/p/-9q96CysQfG

1. "store the values of the struct type person into a MAP (key value) with the KEY being the last name. "Access" each value in the map (values in the map are accessed with `map[v]'). "Print out the values, ranging over the slice." So you'll want to create a for-range loop that goes over a slice you create that stores the values associated with the last name you're ranging over.
*/
