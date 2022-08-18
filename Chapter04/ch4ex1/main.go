package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.packtpub.com/")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringBody := string(data)

	numLinks := strings.Count(stringBody, "<a href=")
	fmt.Printf("Packt Publishing homepage has %d links!\n", numLinks)
}
