package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)
	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.amazonfsdafsa.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || (resp.StatusCode >= 400) {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}
