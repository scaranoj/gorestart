package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}
func sliceDelta(sd []int) { //slices are already reference types, no `*` needed
	//sd = []int{6, 7, 8, 9}
	sd[0] = 55
}

func mapDelta(md map[string]int, s string) {
	md[s] = 20
}

func main() {

	a := 42
	fmt.Println(a)
	intDelta(&a) //delta func takes a pointer as an input parameter, so we give it one. No need here to create an extra variable to assign the address of `a` too
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	sliceDelta(xi) //slices are already reference types, don't have to pass with the address operator (`&`) NOT necessary
	fmt.Println(xi)

	m := make(map[string]int)
	m["Hola"] = 10
	mapDelta(m, "quatro")
	fmt.Println(m)

}
