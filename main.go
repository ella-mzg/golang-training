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

// func PrintIt(input interface{}) {
// 	fmt.Println(input)
// }

func PrintIt(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Println("This is an int :", v)
	case string:
		fmt.Println("This is a string :", v)
	default:
		fmt.Println("This is not an int nor a string :", v)
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println("Error :", err)
	}

	PrintIt(42)
	PrintIt("Hello world")
	PrintIt(3.14)
}
