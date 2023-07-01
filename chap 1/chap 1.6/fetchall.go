// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	topsites := make([]string, 0)
	file, err := os.Open("top-1m.txt")
	input := bufio.NewScanner(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
	for i := 1; i <= 1000; i++ {
		input.Scan()
		topsites = append(topsites, "http://"+input.Text())
	}
	file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range topsites {
		go fetch(url, ch) // start a goroutine
	}
	counter := 1
	for range topsites {
		fmt.Println(counter, <-ch) // receive from channel ch
		counter++
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
