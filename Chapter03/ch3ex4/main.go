package main

import (
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"io"
)

func main() {
	// Set up the local disk cache
	storage := diskcache.New("./cache")
	cache := httpcache.NewTransport(storage)

	// Set this to true to inform us if the responses are being read from a cache
	cache.MarkCachedResponses = true
	cachedClient := cache.Client()

	// Make the initial request
	println("Caching: https://nakit.cz")
	resp, err := cachedClient.Get("https://nakit.cz")
	if err != nil {
		panic(err)
	}

	// httpcache requires you to read the body in order to cache the response
	io.ReadAll(resp.Body)
	resp.Body.Close()

	// Request index.html again
	println("Requesting: https://nakit.cz")
	resp, err = cachedClient.Get("https://nakit.cz")
	if err != nil {
		panic(err)
	}

	// Look for the flag added by httpcache to show the result is read from the cache
	_, ok := resp.Header["X-From-Cache"]
	if ok {
		println("Result was pulled from the cache!")
	}
}
