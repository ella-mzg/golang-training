package main

import (
	"fmt"
	"sync"
)

type Response struct {
	respText string
	err      error
}

func run(c chan Response) {
	resp := Response{
		respText: "Hello from goroutine!",
		err:      nil,
	}
	c <- resp
}

func myFunction(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("It's over")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go myFunction(&wg)

	ch := make(chan Response)

	go run(ch)

	result := <-ch
	fmt.Println("Received:", result.respText)

	wg.Wait()
}
