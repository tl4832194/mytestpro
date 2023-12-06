package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 获取单个url
func GetUrl() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("b: %s\n", b)
	}
}

// 获取多个url
func GetUrlAll(url string, ch chan<- string) {
	start := time.Now()
	// 获取访问url地址的响应信息(即curl url的结果)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbyes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}
	// time.Since()函数用于计算从...到现在的时间
	secs := time.Since(start).Seconds()
	// 将信息传给ch通道，方便在main函数中接收并使用
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbyes, url)
}
