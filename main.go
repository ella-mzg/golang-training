package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type error interface {
	Error() string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, something went wrong: %s", e.When, e.What)
}

func run() error {
	return MyError{
		When: time.Now(),
		What: "efeqzfrfrefrgqgrqrgzS",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("Erreur :", err)
	}
}
