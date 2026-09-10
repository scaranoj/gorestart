package main

import "fmt"

func main() {
	div60()
}

func div60() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Caught panic:", v)
		}
	}()

	for _, val := range []int{10, 50, 0, 100} {
		// Natural panic: division by zero. No need to call panic manually
		result := 60 / val
		fmt.Println(result)
	}
}
