package concurrency

type WebsiteChecker func(url string) bool

func WebsiteCheck(wc WebsiteChecker, url []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range url {
		results[url] = wc(url)
	}
	return results
}
