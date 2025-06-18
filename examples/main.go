package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/Yven/notion_blog/client"
	"github.com/Yven/notion_blog/filter"
	"github.com/Yven/notion_blog/plugin/file"
	"github.com/Yven/notion_blog/plugin/typecho"
	"github.com/Yven/notion_blog/structure"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	var logFilePath string
	var configFile string

	app := cli.NewApp()
	app.Name = "NotionBlog"
	app.Usage = "Request Notion API to build blog"
	app.Version = "1.0"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "log, l",
			Usage:       "log file path",
			Destination: &logFilePath,
		},
		cli.StringFlag{
			Name:        "conf, c",
			Usage:       "config file path",
			Destination: &configFile,
		},
	}

	app.Action = func(c *cli.Context) error {
		// 默认当前程序执行路径下的 env 文件
		if len(configFile) == 0 {
			configFile = ".env"
		}
		// 配置文件不存在
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			log.Panicf("配置文件不存在：%s", err)
		}
		// 读取 env 文件配置
		err := godotenv.Load(configFile)
		if err != nil {
			log.Panicf("配置文件加载失败: %s", err)
		}

		// 全局捕获 panic
		defer func() {
			if err := recover(); err != nil {
				log.Println("【致命错误】")
				log.Panicln(err)
			}
		}()

		// 设置 log 格式
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		// 设置 log 输出流
		log.SetOutput(setLog(logFilePath))

		// 获取文章格式配置
		outputStructType := os.Getenv("BLOG_ARTICLE_TYPE")
		if len(outputStructType) == 0 {
			log.Panicln("文章格式设置不能为空")
		}

		// 设置输出 Blog 模式
		var outputHandle structure.ListWriter = nil
		outputType := os.Getenv("BLOG_OUTPUT_TYPE")
		if outputType == "file" {
			outputPath := os.Getenv("NOTION_PAGE_PATH")
			outputHandle = file.NewFile(outputPath, outputStructType)
		} else if outputType == "typecho" {
			outputHandle = typecho.NewDb(typecho.DbOptions{
				Host:     os.Getenv("BLOG_DB_HOST"),
				Port:     os.Getenv("BLOG_DB_PORT"),
				User:     os.Getenv("BLOG_DB_USER"),
				Passwd:   os.Getenv("BLOG_DB_PASSWORD"),
				Database: os.Getenv("BLOG_DB_NAME"),
				Charset:  os.Getenv("BLOG_DB_CHARSET"),
			})
		} else {
			log.Panicln("输出模式设置错误")
		}
		if outputHandle == nil {
			log.Panicln("输出模式设置错误")
		}

		// 读取 Notion 数据库 ID
		databaseId := os.Getenv("NOTION_DB_ID")
		if len(databaseId) == 0 {
			log.Panicln("请配置 Blog 数据库 ID")
		}

		// 设置请求 Notion 数据库筛选条件
		// condition := filter.Status("Status").Equal("waiting").
		// 	And(filter.MultiSelect("Tag").Contain("test"))
		// 发起请求
		notion := client.NewClient(os.Getenv("NOTION_API_KEY"))
		list, err := notion.NewDb(databaseId).Query(client.QueryDatabase{
			Filter: filter.Status("Status").Equal("edit"),
		})
		// list := client.NewDb(databaseId, &structure.DatabaseList{
		// 	FilterParam: condition,
		// })

		if list == nil {
			return nil
		}

		// 遍历结果
		for _, page := range list.GetContent() {
			// 获取每页的具体内容数据
			// pageContent := page.Fetch(client)
			pageContent, err := notion.NewBlock(page.Id).Children(page, client.BaseQuery{})
			if err != nil {
				log.Panicln(err)
			}
			// 设置输出格式
			pageContent.SetOutputType(outputStructType)
			// 输出
			outputHandle.Writer(pageContent)
			// 修改文章属性为已完成
			notion.NewPage(page.Id).Update(client.UpdatePage{
				Properties: filter.Status("Status").Set("name", "publish"),
			})

			// data, _ := json.Marshal(blockObj)
			// lib.WriteFile("output-block-rever.json", data)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}

// 设置输出日志
func setLog(logFilePath string) io.Writer {
	// 默认输出到标准输出流
	var output io.Writer = os.Stdout
	// 设置了日志输出目录则尝试创建日志文件
	if len(logFilePath) != 0 {
		if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
			// 目录不存在尝试创建目录
			err := os.Mkdir(logFilePath, 0777)
			if err != nil {
				log.Printf("创建日志目录失败：%s", err)
				log.Println("...将以标准模式输出日志")
				return output
			}
		}

		// 输出日志设置
		logFilename := fmt.Sprintf("%s/%s.log", logFilePath, time.Now().Format("2006-01-02"))
		logFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			log.Println("日志文件创建失败")
			log.Println("...将以标准模式输出日志")
			return output
		} else {
			// 同时输出到标准输出和文件中
			output = io.MultiWriter(logFile, os.Stdout)
		}
	}

	return output
}
