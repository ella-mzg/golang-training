package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)

var urlList = flag.String("urls", "", "URLs list")

// go run main.go -urls "https://1000logos.net/wp-content/uploads/2016/10/Android-Logo-768x432.png,https://1000logos.net/wp-content/uploads/2016/10/Android-Logo-768x432.png,https://1000logos.net/wp-content/uploads/2017/06/Windows-Logo.png"

func hashFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		logger.Printf("Failed to open %s: %v", path, err)
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

func hashBytes(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

func downloadAndSave(url string) (string, []byte) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		logger.Printf("Failed to download %s: %v", url, err)
		return "", nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Printf("Failed to read response %s: %v", url, err)
		return "", nil
	}

	duration := time.Since(start)
	fmt.Printf("Downloaded %s in %v\n", url, duration)

	name := filepath.Base(url)
	localPath := filepath.Join(os.TempDir(), name)

	f, err := os.Create(localPath)
	if err != nil {
		logger.Printf("Error when writing %s: %v", localPath, err)
		return "", nil
	}
	defer f.Close()

	f.Write(data)

	return localPath, data
}

func main() {
	flag.Parse()

	hashes := make(map[string][]string)

	for _, path := range flag.Args() {
		hash := hashFile(path)
		if hash == nil {
			continue
		}
		key := string(hash)
		hashes[key] = append(hashes[key], path)
	}

	if *urlList != "" {
		urls := strings.Split(*urlList, ",")
		for _, url := range urls {
			path, data := downloadAndSave(url)
			if data == nil {
				continue
			}
			hash := hashBytes(data)
			key := string(hash)
			hashes[key] = append(hashes[key], fmt.Sprintf("%s (=> %s)", url, path))
		}
	}

	fmt.Println("Unique images :")
	for _, paths := range hashes {
		if len(paths) == 1 {
			fmt.Println(" -", paths[0])
		}
	}
}
