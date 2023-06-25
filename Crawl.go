package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

func CrawlUrl(link string) {
	client := resty.New()

	resp, err := client.R().Get(link)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status", resp.Status())
	res := string(resp.Body())
	fmt.Println(res)

}

func Crawl() {
	CrawlUrl("https://github.com/VuVietDuy")
}
