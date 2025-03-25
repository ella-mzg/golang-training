package main

import (
	"fmt"
	"strconv"
)

type IPAddr []byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func joinIP(ip IPAddr) string {
	result := ""
	for i, b := range ip {
		if i > 0 {
			result += "."
		}
		result += strconv.Itoa(int(b))
	}
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// ip := IPAddr{1, 2, 3, 4}
	// fmt.Println(ip)

	// s := joinIP(ip)
	// fmt.Println(s)
}
