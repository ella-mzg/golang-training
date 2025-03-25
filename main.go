package main

import (
	"encoding/json"
	"fmt"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string `json:"colors,omitempty"`
}
type Animal struct {
	Name  string
	Order string
}

func main() {
	// // Colors
	// group := ColorGroup{
	// 	ID:     1,
	// 	Name:   "Reds",
	// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	// }
	// b, err := json.Marshal(group)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Println(string(b))

	// Animals
	var jsonBlob = []byte(
		`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll", "Order": "Dasyuromorphia"}
	]`)
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		fmt.Println("error:", err)
	}
	// fmt.Println("%+v", animals)
	fmt.Printf("%v", animals)
}
