package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"regexp"
	"sync"
)

func crawlWeb(url string) string {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		fmt.Print(err)
	}

	body := resp.Body()
	html := string(body)

	fmt.Println("crawl successful")
	return html
	//fmt.Print(html)
}

func extractLinks(html string) []string {
	pattern := `href=["'](.*?)["']`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(html, -1)

	return matches
}

func Crawl() {
	url := "https://gobyexample.com/"

	html := crawlWeb(url)

	links := extractLinks(html)

	var mutex sync.Mutex

	//for _, link := range links {
	//	fmt.Println(link)
	//}

	//ch := make(chan string)

	var wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mutex.Lock()

				if len(links) == 0 {
					mutex.Unlock()
					return
				}

				link := links[0]
				links = links[1:]
				mutex.Unlock()

				fmt.Println(link)
				//crawlWeb(link)
			}
		}()
	}

	//for _, link := range links {
	//	ch <- link
	//}
	wg.Wait()
}
