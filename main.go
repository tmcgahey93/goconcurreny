package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchURL(url string) string {
	sleepTime := time.Duration(rand.Intn(10000)) * time.Millisecond
	time.Sleep(sleepTime)
	return fmt.Sprintf("Data from %s, sleepTime %v", url, sleepTime)
}

func main() {
	//rand.Seed(time.Now().UnixNano())

	urls := []string{
		"url1",
		"url2",
		"url3",
		"url4",
		"url5",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			data := fetchURL(u)
			results <- data
		}(url)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}
}
