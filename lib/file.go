package lib

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/gabriel-vasile/mimetype"
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

func DownloadFile(fileUrl string) string {
	var Url *url.URL
	Url, err := url.Parse(fileUrl)
	if err != nil {
		log.Fatalf("解析url错误:%v", err)
	}

	arr := []string{"jpg", "jpeg", "png", "svg", "bmp", "gif", "tif", "tiff", "heic"}
	ext := path.Ext(Url.Path)
	if !Contain(arr, ext[1:]) {
		return fileUrl
	}

	res, err := http.Get(Url.String())
	defer res.Body.Close()
	file, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("下载失败:%v", err)
	}

	mtype := mimetype.Detect(file)

	return fmt.Sprintf("data:%s;base64,%s", mtype.String(), base64.StdEncoding.EncodeToString(file))
}
