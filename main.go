package main

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	respText string
	err      error
}

func callServer(addr string, ch chan Response) {
	resp, err := http.Get(addr)
	if err != nil {
		ch <- Response{respText: "", err: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response{respText: "", err: err}
		return
	}

	ch <- Response{respText: string(body), err: nil}
}

func main() {
	ch := make(chan Response)

	go callServer("https://github.com/ella-mzg", ch)

	result := <-ch

	if result.err != nil {
		fmt.Println("Error :", result.err)
	} else {
		fmt.Println("Response :", result.respText)
	}
}
