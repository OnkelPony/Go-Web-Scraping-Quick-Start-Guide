package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	doc.Find(`div.book-block-outer div[itemtype$="/Product"]`).
		Each(func(i int, e *goquery.Selection) {
			var title string
			var price float64

			title, _ = e.Attr("data-product-title")
			priceString, _ := e.Attr("data-product-price")
			price, err = strconv.ParseFloat(priceString, 64)
			if err != nil {
				println("Failed to parse price")
			}
			fmt.Printf("%54s ($%6.2f)\n", title, price)
		})
}
