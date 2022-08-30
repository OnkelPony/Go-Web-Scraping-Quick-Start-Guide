package main

import (
	"bufio"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	resp, err := http.Get("https://web.archive.org/web/20170223095930/https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	println("Here is the free book of the day!")
	println("----------------------------------")
	rawText := doc.Find(`div.dotd-main-book-summary div:not(.eighteen-days-countdown-bar)`).Text()
	reader := bufio.NewReader(strings.NewReader(rawText))

	var line []byte
	for err == nil {
		line, _, err = reader.ReadLine()
		trimmedLine := strings.TrimSpace(string(line))
		if trimmedLine != "" {
			println(trimmedLine)
		}
	}
}
