package checkurl

import (
	"fmt"
	"net/http"
)

type reqResult struct {
	url   string
	staus string
}

// CheckURL checks url is up to go
func CheckURL() {
	results := make(map[string]string)
	c := make(chan reqResult)

	var urls = []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for range urls {
		result := <-c
		results[result.url] = result.staus
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan<- reqResult) {
	status := "OK"
	if res, err := http.Get(url); err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- reqResult{url: url, staus: status}
}
