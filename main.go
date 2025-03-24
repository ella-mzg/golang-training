package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func readFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error when opening file:", err)
	}

	info, _ := f.Stat()
	size := info.Size()
	buffer := make([]byte, size)
	f.Read(buffer)
	return buffer
}

func hashFile(path string) []byte {
	data := readFile(path)
	sum := sha256.Sum256(data)
	return sum[:]
}

func main() {
	file1 := os.Args[1]
	file2 := os.Args[2]

	hash1 := hashFile(file1)
	fmt.Println(hash1)
	hash2 := hashFile(file2)
	fmt.Println(hash2)

	if string(hash1) == string(hash2) {
		fmt.Println("Identiques") // 1 et 3 identiques, 2 unique
	} else {
		fmt.Println("Différents")
	}
}
