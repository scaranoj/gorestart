// Create two functions to change a field in a struct called "first" of TYPE string:
// One function will use VALUE SEMANTICS
// 	this function will return some VALUE of some TYPE
// The other function will use POINTER SEMANTICS
// 	this function will return nothing
// don't use methods

package main

import "fmt"

type dog struct {
	first string
}

// got this right
func nameChange(d dog, newName string) dog {
	d.first = newName
	return d
}

func nameChangePtr(d *dog, newName string) dog {
	d.first = newName
	return *d
}

func main() {

	// if mutating struct field, need define the initial assignment value first
	d1 := dog{
		first: "fido",
	}

	//can print starting value just to show contrast in output
	fmt.Println(d1.first)

	//could condense into one line though seems idiomatic to assign the return first seperately before using it
	d := nameChange(d1, "charlie")
	fmt.Println(d.first)

	d2 := dog{
		first: "crackers",
	}
	fmt.Println(d2.first)

	e := nameChangePtr(&d2, "butters")
	fmt.Println(e.first)

}
