package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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

func Notion(method Method, path string, data io.Reader) (rs []byte, err error) {
	var apiURL string
	if apiHost := os.Getenv("NOTION_API"); len(apiHost) != 0 {
		apiURL = os.Getenv("NOTION_API") + path
	} else {
		return []byte{}, fmt.Errorf("Notion API 配置为空")
	}
	header := make(map[string]string)
	if apiVersion := os.Getenv("NOTION_VERSION"); len(apiVersion) != 0 {
		header["Notion-Version"] = apiVersion
	} else {
		return []byte{}, fmt.Errorf("Notion Version 配置为空")
	}
	if apiKey := os.Getenv("NOTION_KEY"); len(apiKey) != 0 {
		header["Authorization"] = "Bearer " + apiKey
	} else {
		return []byte{}, fmt.Errorf("Notion Key 配置为空")
	}

	return Request(method, apiURL, path, header, data)
}

func Request(method Method, apiURL string, path string, header map[string]string, data io.Reader) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		return []byte{}, fmt.Errorf("解析url错误: %s", err)
	}
	count := 0
	for {
		count += 1
		req, _ := http.NewRequest(string(method), Url.String(), data)
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
				res, err := ioutil.ReadAll(resp.Body)
				log.Println(string(res))
				log.Println(method, Url.String())
				return []byte{}, fmt.Errorf("未知状态码%d：%s", resp.StatusCode, err)
			}
		}

		if resp.ContentLength != 0 || count > 3 {
			defer resp.Body.Close()
			return ioutil.ReadAll(resp.Body)
		}
	}
}
