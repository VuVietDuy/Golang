package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"regexp"
)

func CrawlUrl(link string) string {
	client := resty.New()

	resp, err := client.R().Get(link)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode() != 200 {
		fmt.Println("status", resp.Status())
	}

	res := string(resp.Body())
	fmt.Println(res)
	return res
}

func findLink(html string) []string {
	//links := make([]string, 0)
	pattern := `href="([^"]*)"`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(html, -1)
	fmt.Println(matches)
	return matches
}

func Crawl() {
	html := CrawlUrl("https://gobyexample.com/")
	findLink(html)
}
