package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string) result {
	resp, err := http.Head(url)
	if err != nil {
		return result{url: url, code: 0}
	}
	return result{url: url, code: resp.StatusCode}
}

func main() {
	start := time.Now()
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
	var results []result

	for _, url := range urlList {
		results = append(results, fetch(url))
	}

	for _, result := range results {
		fmt.Printf("WorkerID: %v, url: %v, code: %v \n", result.workerID, result.url, result.code)
	}
	fmt.Println(time.Since(start))
}
