package concurrency

import "time"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}

func MockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func SlowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
