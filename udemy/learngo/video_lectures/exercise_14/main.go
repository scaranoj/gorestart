package main

import "fmt"

var movie string = "The Big Lebowski"

const age int = 42

func main() {
	charName1 := "\"The Dude\""
	charName2 := "Walter"
	fmt.Printf("%v from the movie %s is %d years old. %s's age is unknown.", charName1, movie, age, charName2)

}
