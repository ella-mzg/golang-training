package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)

func hashFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		logger.Printf("Error when opening file %s: %v", path, err)
		return nil
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	hasher := sha256.New()

	buf := make([]byte, 4096)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			hasher.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}

	return hasher.Sum(nil)
}

func main() {
	if len(os.Args) < 2 {
		logger.Println("No file to analyze")
		return
	}

	hashes := make(map[string][]string)

	for _, path := range os.Args[1:] {
		hash := hashFile(path)
		if hash == nil {
			continue
		}
		key := string(hash)
		hashes[key] = append(hashes[key], path)
	}

	fmt.Println("Unique files :")
	for _, paths := range hashes {
		if len(paths) == 1 {
			fmt.Println(" -", paths[0])
		}
	}
}
