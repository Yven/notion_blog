package file

import (
	"fmt"
	"log"
	"os"
	"time"

	"notion_blog/notion"
)

type File struct {
	handler *os.File
}

func (f *File) Writer(list *notion.List) error {
	_, err := f.handler.Write([]byte(list.Output()))
	return err
}

// 设置输出文件
func NewFile(outputPath string, fileType string) *File {
	// 输出 Blog 文件模式设置
	if len(outputPath) != 0 {
		var err error
		if _, err = os.Stat(outputPath); os.IsNotExist(err) {
			err := os.Mkdir(outputPath, 0777)
			if err != nil {
				log.Panicln("创建输出目录失败：%s", err)
				return nil
			}
		}

		outputName := fmt.Sprintf("%s/%s.%s", outputPath, time.Now().Format("200601021504"), fileType)
		f, err := os.OpenFile(outputName, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			log.Panicln("创建输出文件失败：%s", err)
			return nil
		}
		err = f.Truncate(0)
		if err != nil {
			log.Panicln("文件操作失败：%s", err)
			return nil
		}

		return &File{
			handler: f,
		}
	}

	return nil
}
