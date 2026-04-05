package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	Person
	ltk   bool
	first string
}

func main() {

	sa1 := secretAgent{
		Person: Person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk:   true,
		first: "Ethan Hawk",
	}

	p1 := Person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(sa1)
	fmt.Println(sa1.first, p1.first)
	fmt.Println(sa1.ltk, sa1.Person.first)
}
