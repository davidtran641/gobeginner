package concurrency

import (
	"time"
)

type WebsiteChecker func(string) bool

type result struct {
	url string
	res bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url: u, res: wc(u)}
		}(url)
	}

	timeout := time.After(100 * time.Microsecond)

	for i := 0; i < len(urls); i++ {
		select {
		case result := <-resultChannel:
			results[result.url] = result.res
		case <-timeout:
			return results
		}
	}

	return results
}
