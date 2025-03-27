package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type Response struct {
	addr     string
	respText string
	err      error
}

func callServer(addr string, ch chan Response, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(addr)
	if err != nil {
		ch <- Response{addr: addr, err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- Response{
			addr: addr,
			err:  errors.New("HTTP error: " + strconv.Itoa(resp.StatusCode)),
		}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response{addr: addr, err: err}
		return
	}

	ch <- Response{addr: addr, respText: string(body)}
}

func main() {
	urls := []string{
		"http://localhost:8000/?id=1",
		"http://localhost:8000/?id=2",
		"http://localhost:8000/?id=3",
	}

	var wg sync.WaitGroup
	ch := make(chan Response, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go callServer(url, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for res := range ch {
		fmt.Println("URL:", res.addr)
		if res.err != nil {
			fmt.Println("Error:", res.err)
		} else {
			fmt.Println("Response:", res.respText)
		}
	}
}
