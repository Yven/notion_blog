package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Method string

const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Delete Method = "DELETE"
	Patch  Method = "PATCH"
)

func Request(method Method, apiURL *url.URL, header map[string]string, data io.Reader) (rs []byte, err error) {
	count := 0
	for {
		count += 1
		req, err := http.NewRequest(string(method), apiURL.String(), data)
		if err != nil {
			return []byte{}, fmt.Errorf("创建请求失败: %s", err)
		}
		if method == Post || method == Patch {
			req.Header.Set("Content-Type", "application/json")
		}

		for key, val := range header {
			req.Header.Set(key, val)
		}

		client := http.Client{}
		client.Timeout = 30 * time.Second
		resp, err := client.Do(req)
		if err != nil {
			return []byte{}, fmt.Errorf("请求失败: %s", err)
		}

		if resp.StatusCode != 200 {
			if resp.StatusCode == 429 {
				after, err := strconv.Atoi(resp.Header.Get("Retry-After"))
				if err == nil {
					time.Sleep(time.Duration(after) * time.Second)
					continue
				} else {
					return []byte{}, fmt.Errorf("Retry-After 参数错误[%s]: %s", resp.Header.Get("Retry-After"), err)
				}
			} else {
				res, err := io.ReadAll(resp.Body)
				log.Println(string(res))
				log.Println(method, apiURL.String())
				return []byte{}, fmt.Errorf("未知状态码%d：%s", resp.StatusCode, err)
			}
		}

		if resp.ContentLength != 0 || count > 3 {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		}
	}
}
