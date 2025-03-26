package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID       string `json:"userID"`
	Login    string `json:"userName"`
	Password string
}

var usersMap = make(map[string]User)

func handleUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if user, ok := usersMap[id]; ok {
		json.NewEncoder(w).Encode(user)
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func main() {
	data, err := os.ReadFile("server/users.json")
	if err != nil {
		log.Fatal("Reading error:", err)
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal("Parsing error:", err)
	}

	for _, user := range users {
		usersMap[user.ID] = user
	}

	http.HandleFunc("/", handleUser)
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
