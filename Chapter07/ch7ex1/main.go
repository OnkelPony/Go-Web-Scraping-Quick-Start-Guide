package main

import (
	"fmt"
	"time"
)

func startTicker() {
	ticks := 0
	for {
		fmt.Println(ticks)
		ticks++
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	println("Starting ticker")
	go startTicker()
	time.Sleep(1 * time.Second)
}
