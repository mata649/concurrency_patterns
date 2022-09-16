package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("URL: %v, Status: %v", url, 0)

	}

	fmt.Println("URL: %v, Status: %v", url, resp.Status)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

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

	for _, url := range urlList {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
