package checkurl

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

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

// CheckURL checks url is up to go
func CheckURL() {
	var results = make(map[string]string)

	c := make(chan error)

	// start channel
	for _, url := range urls {
		go hitURL(url, c)
	}

	// get channel
	for _, url := range urls {
		result := "OK"
		if <-c == errRequestFailed {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan error) {
	if res, err := http.Get(url); err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		c <- errRequestFailed
	}
	c <- nil
}
