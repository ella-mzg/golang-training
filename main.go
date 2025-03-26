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

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// func handleUser(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	if user, ok := usersMap[id]; ok {
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 	}
// }

func handleUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if user, ok := usersMap[id]; ok {
		respondJSON(w, http.StatusOK, user)
	} else {
		respondError(w, http.StatusNotFound, "User not found")
	}
}

// func handleUser(w http.ResponseWriter, r *http.Request) {
// 	id := r.FormValue("id")

// 	// w.Header().Set(
// 	// 	"Content-Type",
// 	// 	"application/json; charset=utf-8",
// 	// )

// 	// if user, ok := usersMap[id]; ok {
// 	// 	w.WriteHeader(http.StatusOK)
// 	// 	json.NewEncoder(w).Encode(user)
// 	// } else {
// 	// 	w.WriteHeader(http.StatusNotFound)
// 	// 	http.Error(w, "User not found", http.StatusNotFound)
// 	// }

// 	if user, ok := usersMap[id]; ok {
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		w.WriteHeader(http.StatusOK)

// 		userJson, err := json.Marshal(user)
// 		if err != nil {
// 			http.Error(w, "Error", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write(userJson)
// 	}
// }

func main() {
	data, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatal("Reading error :", err)
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		log.Fatal("Parsing error :", err)
	}

	for _, user := range users {
		usersMap[user.ID] = user
	}

	http.HandleFunc("/user", handleUser)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
