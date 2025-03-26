package main

import (
	"fmt"
	"sync"
)

func myFunction(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("It's over")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go myFunction(&wg)

	fmt.Println("End of program")

	wg.Wait()
}

// go tool dist list
