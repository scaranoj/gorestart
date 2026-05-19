package main

import "fmt"

func main() {

	bob := person{
		first: "bob",
		age:   32,
	}

	bob.speak("bob", 32)

}

type person struct {
	first string
	age   int
}

func (p person) speak(f string, a int) {
	fmt.Printf("Hello %v, you are %v years old", f, a)
}

/*
Create a user defined struct with
* the identifier “person”
* the fields:
* ■ first
* ■ age
* attach a method to type person with
  * the identifier “speak”
  * the method should have the person say their name and age
* create a value of type person
* call the method from the value of type person


*/
