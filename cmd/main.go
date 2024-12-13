package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/wowTerrific/HTTPing/ping"
)

func sendPing(url string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- fmt.Sprintf("%s - %s", url, ping.Ping(url))
}

func main() {
	// TODO - delete this
	urls := []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"aslfkdjsf",
	}

	var wg sync.WaitGroup
	c := make(chan string)

	for _, url := range urls {
		wg.Add(1)
		go sendPing(url, c, &wg)

	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for resp := range c {
		if !strings.Contains(resp, "200 OK") {
			fmt.Printf("%s%s%s\r\n", "\033[31m", resp, "\033[0m")
		} else {
			fmt.Println(resp)
		}
	}
}
