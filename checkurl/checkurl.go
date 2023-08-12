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

func checkURL() {
	var results = make(map[string]string)

	for _, url := range urls {
		fmt.Println("Checking url:", url)
		result := "OK"
		if err := hitURL(url); err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	if res, err := http.Get(url); err != nil || res.StatusCode >= 400 {
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}
