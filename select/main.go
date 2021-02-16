package main

import (
	"fmt"
	"net/http"
	"time"
)

const limitOfTenSecond = 10 * time.Second

func measureResponseTime(URL string) time.Duration {
	start := time.Now()
	http.Get(URL)
	return time.Since(start)
}

// Runner compares two URLs and selects the fastest one
func Runner(a, b string) (winner string, erro error) {
	return Config(a, b, limitOfTenSecond)
}

// Config is a config to runner
func Config(a, b string, timeLimit time.Duration) (winner string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeLimit):
		return "", fmt.Errorf("timeout of requisition 10 seconds")
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()

	return ch
}

func main() {

}
