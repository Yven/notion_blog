package lib

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func WriteFile(filename string, data []byte) {
	// 写入文件
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Truncate(0)
	if err != nil {
		log.Fatal("清空文件内容失败:", err)
	}

	if _, err := f.Write(data); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filename string) []byte {
	// 文件存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal("文件不存在")
	}
	if path.Ext(filename) != ".json" {
		log.Fatal(filename + "不是 json 文件!")
	}

	buffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return buffer
}

func Contain[T comparable](arr []T, target T) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}

	return false
}

func DownloadFile(fileUrl string) ([]byte, error) {
	var Url *url.URL
	Url, err := url.Parse(fileUrl)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	file, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return file, nil
}
