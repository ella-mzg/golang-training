package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Login    string `json:"userName"`
	Password string
}

func main() {
	u := User{
		Login:    "Pierre",
		Password: "pass012", // lowercase 'p'assword = ignored during unmarshaling
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error serializing:", err)
		return
	}

	fmt.Println(string(jsonData))
	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println(users)
}
