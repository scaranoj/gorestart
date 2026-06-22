package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	ID      int64 `json:"ref"`
	private string
	Created time.Time
}

func main() {

	basket := FruitBasket{
		Name:    "Variety",
		Fruit:   []string{"orange", "apples", "dates", "pineapple"},
		ID:      1234,
		private: "bottom",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(basket, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println((string)(jsonData))

}

//jsonData, err := json.MarshalIndent(basket, "", "    ")
