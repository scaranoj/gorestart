package main

import "fmt"

func main() {

	a := 34
	fmt.Println(a)
	b := &a
	fmt.Println(b)
	fmt.Println(intDelta(5))
	fmt.Println(intDeltaP(b))
}

func intDelta(a int) int {
	return a - 1
}

func intDeltaP(p *int) int {
	return *p - 1 //initially forgot (when trying to recall from memory) to dereference `p` so it can be used, but then remembered when linter gave warning
}
