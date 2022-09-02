package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://web.archive.org/web/20170223095930/https://www.packtpub.com/latest-releases")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Here are the latest releases!")
	println("-----------------------------")
	time.Sleep(1 * time.Second)
	doc.Find(`div.book-block-outer div[itemtype$="/Product"] a`).
		Each(func(i int, e *goquery.Selection) {
			var title, description, author, price string
			link, _ := e.Attr("href")
			//fmt.Printf("Link: %s", link)
			link = "https://web.archive.org" + link

			respPage, err := http.Get(link)
			if err != nil {
				panic(err)
			}
			bookPage, err := goquery.NewDocumentFromReader(respPage.Body)
			if err != nil {
				panic(err)
			}
			title = bookPage.Find("div.book-top-block-info h1").Text()
			description = strings.TrimSpace(bookPage.Find("div.book-top-block-info div.book-top-block-info-one-liner").Text())
			price = strings.TrimSpace(bookPage.Find("div.book-top-pricing-block div.book-top-pricing-price div.book-top-pricing-main-book-price").Text())
			authorNodes := bookPage.Find("div.book-top-block-info div.book-top-block-info-authors")
			if len(authorNodes.Nodes) < 1 {
				return
			}
			author = strings.TrimSpace(authorNodes.Nodes[0].FirstChild.Data)
			fmt.Printf("%s\nby: %s\n%s\n%s\n---------------------\n\n", title, author, price, description)
			time.Sleep(1 * time.Second)
		})
}
