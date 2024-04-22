package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func downloadPage(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	fileName := path.Base(u.Path)
	if fileName == "" || fileName == "/" {
		fileName = "Download.html"
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: go run wget.go <url>")
		os.Exit(1)
	}

	urlStr := args[0]
	err := downloadPage(urlStr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Downloaded:", urlStr)
}
