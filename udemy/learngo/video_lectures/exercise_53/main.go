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

	for _, v := range p1.flavor {
		fmt.Println(p1.first, "likes", v)
	}
	for _, v := range p2.flavor {
		fmt.Println(p2.first, "likes", v)
	}

	fmt.Println(p1)
	fmt.Println(p2)

}

/*
Hands-on exercise #53 - struct with slice
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
https://go.dev/play/p/nLcea3bIb7h
*/
