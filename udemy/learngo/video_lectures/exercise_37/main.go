package main

import "fmt"

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
		"Q":          52,
	}
	age := m["James"]
	fmt.Println("The age of Bond\t", age)
	ageq := m["Q"]
	fmt.Println("The age of Q\t", ageq)

	// comma ok idiom ex from effective go:
	// if seconds, ok := timeZone[tz]; ok {
	//     return seconds
	// }

	// This was my original version of the comma ok idiom, which works, though the instructor was able to make more concise
	// v, ok := m["Q"]
	// if ok == true {
	// 	fmt.Println("Q is in the map", v, ok)
	// }

	//instructor's version
	if v, ok := m["Q"]; ok {
		fmt.Println("Q is in the map", v, ok)
	}

	for k, v := range m {
		fmt.Printf("key %v value is %v\n", k, v)
	}

}
