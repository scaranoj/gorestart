package main

import "fmt"

func main() {

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "I'm 008"}
	xx := [][]string{xs1, xs2}
	fmt.Println(xx)

	for i, v := range xx {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}

/*
* Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
~~~
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
~~~
* Range over the records, then range over the data in each record.
* code: https://go.dev/play/p/2jxGMyXiZrE
*/
