package ZuoYeBang

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

/*
	from 作业帮
		GO 并发请求三个URL，之后将请求信息写入文件 用什么方式
*/
var (
	mutex sync.Mutex
)

func mainProcess(urls []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, file, &wg)
	}
	wg.Wait()
	return nil
}

func fetchUrl(url string, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %s\n", url, err)
		return
	}

	mutex.Lock()         // 加锁
	defer mutex.Unlock() // 解锁

	if _, err := file.WriteString(fmt.Sprintf("Response from %s: %s\n", url, body)); err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
	}
}

/*
	扩展：
		GO 并发请求三个URL，同时用context控制2秒超时，请求后写入同一个文件，用什么方式
*/
func mainProcess1(urls []string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	for _, url := range urls {
		wg.Add(1)
		go fetchUrlWithContext(ctx, url, file, &wg)
	}
	wg.Wait()
	return nil
}

func fetchUrlWithContext(ctx context.Context, url string, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()
	select {
	case <-ctx.Done():
		fmt.Printf("Request to %s canceled due to timeout\n", url)
		return
	default:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Error reading response body from %s: %s\n", url, err)
			return
		}

		mutex.Lock()         // 加锁
		defer mutex.Unlock() // 解锁

		if _, err := file.WriteString(fmt.Sprintf("Response from %s: %s\n", url, body)); err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
		}
	}
}
