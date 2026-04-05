package main

import "fmt"

func main() {

	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}

	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Plymouth",
		model: "Fury",
		doors: 2,
		color: "Black",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "BYD",
		model: "Han",
		doors: 4,
		color: "Grey",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("The Fury is a model of car made by", v1.make)
	fmt.Println("The BYD Han engine is electric", v2.engine.electric)

}

/*

* Embed Struct Exercise:

  * Create a type engine struct, and include this field
    * electric bool
  * Create a type vehicle struct, and include these fields
    * engine
    * make
    * model
    * doors
    * color
  * Create two VALUES of TYPE vehicle
    * use a composite literal
  * Print out each of these values.
  * Print out a single field from each of these values.
*/
