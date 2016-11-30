/*
Fetches several web pages simultaneously using the
net/http package, and prints the URL of the biggest
home page (defined as the most bytes in the response)
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	size int
}

func main() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}
	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			results <- HomePageSize{
				URL:  url,
				size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.size > biggest.size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}

// The biggest home page: http://www.amazon.com
