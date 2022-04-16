package main

import (
	"log"
	"time"
)

func generateThrottled(data string, bufferLimit int, clearInterval time.Duration) <-chan string {
	out := make(chan string, bufferLimit)

	go func() {
		for {
			for i := 0; i < bufferLimit; i++ {
				out <- data
			}
			time.Sleep(time.Duration(clearInterval))
		}
	}()
	return out
}

func main() {

	out := generateThrottled("foo", 2, time.Second)
	for f := range out {
		log.Println(f)
	}
}
