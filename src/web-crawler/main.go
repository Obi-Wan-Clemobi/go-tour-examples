package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(depth)

	lookupTable := make(map[string]bool)

	_Crawl(url, depth, fetcher, &waitGroup, lookupTable)

	waitGroup.Wait()
	return
}

// Fetch URLs in parallel
// Don't fetch the same URL twice
func _Crawl(url string, depth int, fetcher Fetcher, waitGroup *sync.WaitGroup, lookupTable map[string]bool) {
	if depth <= 0 {
		return
	}

	if !lookupTable[url] {
		body, urls, err := fetcher.Fetch(url)
		lookupTable[url] = true

		if err != nil {
			fmt.Println(err)
			return
		}

		defer waitGroup.Done()

		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go _Crawl(u, depth-1, fetcher, waitGroup, lookupTable)
		}
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
