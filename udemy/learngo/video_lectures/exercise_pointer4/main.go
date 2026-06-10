package main

import "fmt"

func intDelta(id *int) {
	*id = 42
}

func sliceDelta(sd []int) {
	//sd = []int{6, 7, 8, 9, 10}
	sd[0] = 42 //remembering the slice syntax here to set a index/value pair
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33 //remembering the map syntax here to set a key/value pair
}

func main() {

	i := 69
	iPtr := &i
	fmt.Println(i)
	fmt.Println(iPtr)
	intDelta(iPtr)
	fmt.Println(i)

	xi := []int{1, 2, 3, 4, 5}
	sliceDelta(xi)
	fmt.Println(xi)

	m := map[string]int{"Hallo!": 16}
	mapDelta(m, "Bonjour!")
	fmt.Println(m)

}
