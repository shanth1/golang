package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"net/http"
	"os"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при создании файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Fprintf(file, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}
	fmt.Println("Status code:", resp.Status)

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: copy %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
