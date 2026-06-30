package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	//define json data
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":28}]`
	//unmarshal takes a []byte so do a byte slice conversion
	sb := []byte(s)
	//declare variable of type person to write into
	var people []person
	//call Unnmarshal
	err := json.Unmarshal(sb, &people)
	//check for err
	if err != nil {
		fmt.Println(err)
	}
	//print original json
	fmt.Println(string(sb))
	//print all of the unmarshalled data
	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nindex at", i)
		fmt.Println("\nthe value is", v.First, v.Last, v.Age)
	}
}

// unmarshaling json
// what does the func need? ([]byte for the input and a any pointer to output to)
// create struct (use jason-to-go site to help get the type definition you'll need)
// create json data (just use output data from marshal exercise)
// create destination variable
// call json.Unmarshal
// print result
//example code https://go.dev/play/p/O9q0DmpzsZ
