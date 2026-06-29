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

	// basket := FruitBasket{
	// 	Name:    "Variety",
	// 	Fruit:   []string{"orange", "apples", "dates", "pineapple"},
	// 	ID:      1234,
	// 	private: "bottom",
	// 	Created: time.Now(),
	// }

	jsonData := []byte(`
{	
	"Name": "Variety",
    "Fruit": [
        "orange",
        "apples",
        "dates",
        "pineapple"
    ],
    "ref": 1234,
    "Created": "2026-06-15T11:07:42.394800055-06:00"
}	
	`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(basket.Name, basket.Fruit, basket.ID)
	fmt.Println(basket.Created)

}

//jsonData, err := json.MarshalIndent(basket, "", "    ")
