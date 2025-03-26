package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Response struct {
	addr     string
	respText string
	err      error
}

func callServer(addr string, ch chan Response) {
	resp, err := http.Get(addr)
	if err != nil {
		ch <- Response{addr: addr, respText: "", err: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ch <- Response{
			addr:     addr,
			respText: "",
			err:      errors.New("HTTP Error: " + strconv.Itoa(resp.StatusCode)),
		}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Response{addr: addr, respText: "", err: err}
		return
	}

	ch <- Response{addr: addr, respText: string(body), err: nil}
}

func main() {
	ch1 := make(chan Response)
	ch2 := make(chan Response)

	go callServer("http://localhost:8000/?id=1", ch1)
	go callServer("http://localhost:8000/?id=2", ch2)

	select {
	case res := <-ch1:
		fmt.Println("Fastest endpoint:", res.addr)
		if res.err != nil {
			fmt.Println("Error:", res.err)
		} else {
			fmt.Println("Response:", res.respText)
		}
	case res := <-ch2:
		fmt.Println("Fastest endpoint:", res.addr)
		if res.err != nil {
			fmt.Println("Error:", res.err)
		} else {
			fmt.Println("Response:", res.respText)
		}
	}

	// res1 := <-ch1
	// res2 := <-ch2

	// for _, res := range []Response{res1, res2} {
	// 	fmt.Println("🔗 URL:", res.addr)
	// 	if res.err != nil {
	// 		fmt.Println("Error:", res.err)
	// 	} else {
	// 		fmt.Println("Response:", res.respText)
	// 	}
	// 	fmt.Println()
	// }
}
