package interview

import (
	"context"
	"fmt"
	"time"
)

func makeRequest(url string) string {
	s := "make request url is: " + url
	return s
}

func handleRequest(ctx context.Context, urls []string) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	resultChan := make(chan string, len(urls))
	for i, url := range urls {
		go func(url string, i int) {
			if i == 1 {
				time.Sleep(time.Second * 6)
			}
			select {
			case <-ctx.Done():
				fmt.Println("cancel request to ", url)
			default:
				result := makeRequest(url)
				resultChan <- result
			}
		}(url, i)
	}
	for i := 0; i < len(urls); i++ {
		select {
		case res := <-resultChan:
			fmt.Println("get result", res)
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		}
	}
}

func Test() {
	urls := []string{"http://localhost", "https://localhost", "http://localhost:1111"}
	handleRequest(context.Background(), urls)
}
