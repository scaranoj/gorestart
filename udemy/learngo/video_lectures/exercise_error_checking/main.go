//write code with errors before writing code w/o errors. Always check for errors

package main

import "fmt"

func main() {

	n, err := fmt.Println("Hello")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}
