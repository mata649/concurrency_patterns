package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	code     int
	url      string
	workerID int
}

func urlGenerator(urlList []string) <-chan string {
	urlStream := make(chan string)

	go func() {
		for _, url := range urlList {
			urlStream <- url
		}
		close(urlStream)
	}()

	return urlStream
}

func fetchUrl(workerID int, urlStream <-chan string, done chan bool) <-chan result {
	resultStream := make(chan result)
	go func() {
		defer close(resultStream)
		for url := range urlStream {
			select {
			case <-done:
				return
			default:
				resp, err := http.Head(url)
				if err != nil {
					resultStream <- result{code: 0, url: url, workerID: workerID}
					continue
				}
				resultStream <- result{code: resp.StatusCode, url: url, workerID: workerID}
			}
		}
	}()
	return resultStream
}

func mergeResultStream(channels []<-chan result, done chan bool) <-chan result {
	var wg sync.WaitGroup
	resultMerged := make(chan result)
	output := func(resultStream <-chan result) {
		defer wg.Done()
		for result := range resultStream {

			select {
			case resultMerged <- result:
			}
		}
	}

	wg.Add(len(channels))

	for _, channel := range channels {
		go output(channel)
	}
	go func() {
		wg.Wait()
		close(resultMerged)
	}()
	return resultMerged
}

func main() {
	start := time.Now()
	done := make(chan bool)
	defer close(done)
	nWorkers := 20
	urlList := []string{
		"http://www.youtube.com",
		"http://www.facebook.com",
		"http://www.baidu.com",
		"http://www.yahoo.com",
		"http://www.amazon.com",
		"http://www.wikipedia.org",
		"http://www.qq.com",
		"http://www.google.co.in",
		"http://www.twitter.com",
		"http://www.live.com",
		"http://www.taobao.com",
		"http://www.bing.com",
		"http://www.instagram.com",
		"http://www.weibo.com",
		"http://www.sina.com.cn",
		"http://www.linkedin.com",
		"http://www.yahoo.co.jp",
		"http://www.msn.com",
		"http://www.vk.com",
		"http://www.google.de",
	}
	urlStream := urlGenerator(urlList)
	var resultChans []<-chan result
	for i := 0; i < nWorkers; i++ {
		resultChans = append(resultChans, fetchUrl(i, urlStream, done))

	}
	results := mergeResultStream(resultChans, done)

	for result := range results {

		fmt.Printf("WorkerID: %v, url: %v, code: %v \n", result.workerID, result.url, result.code)

	}
	fmt.Println(time.Since(start))
}
