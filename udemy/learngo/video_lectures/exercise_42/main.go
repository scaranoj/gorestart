// Hands-on exercise #42
// ●
//  Using a COMPOSITE LITERAL:
// ● create an ARRAY which holds 5 VALUES of TYPE int
// ● assign VALUES to each index position.
// ● Range over the array and print the values out.
// ○ print out the VALUE and the TYPE
// https://go.dev/play/p/aghCIHXlIsB

package main

import "fmt"

func main() {

	x := [5]int{0, 1, 2, 3, 4}

	for i, v := range x {
		fmt.Printf("Value - %v\t Type - %T\n", v, i)
	}
}
