package main

import (
	"fmt"
	"strings"
)

func main() {

	const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

	xs := strings.Split(s, " ")

	//its a slice so prob want to range over it

	for _, v := range xs {
		fmt.Println(v)
	}

	//fmt.Println(xs)

	//fmt.Printf("\n%s\n", mystr.Cat(xs))

	//this will run since it's already in the std lib but might need to be in my local package
	// for the sake of benchmark comparison? Try creating a simple wrapper for join in the mystr pkg
	j := strings.Join(xs, " ")
	fmt.Println(j)

	// this is a more efficient method
	// how to know? Add these to a _test.go file, run a `go test -bench`
	fmt.Println(strings.Join(xs, " "))

}
