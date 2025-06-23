<div align="center">

# Notion Blog
[![Build Status](https://img.shields.io/badge/notoin_blog-1.0-69cafd)](https://github.com/Yven/notion_blog)

调用Notion API获取内容，导出页面为`.md`

</div>

## 特性
- [x] 调用Notion API获取页面内容
- [x] 输出页面为Markdown
- [ ] 输出页面为HTML

## 使用
### 1.配置参数

1. 安装:
    ```shell
    go get github.com/Yven/notion_blog
    ```
2. 参考[官方文档](https://developers.notion.com/docs/create-a-notion-integration)创建 integration 并获取 Token
3. 在你的 Database 中点击 share 然后 copy link 获取链接中的 `database_id`，如此形式：`https://www.notion.so/{name}/{database_id}?v={view_id}`
4. 使用:
    ```golang
    key := "Your_Integration_Token"
	DbId := "databse_id"

	notion := NewClient(key)
	list, err := notion.NewDb(DbId).Query(QueryDatabase{
		Filter: filter.Status("Status").Equal("Done"),
	})
	if err != nil {
		log.Fatal(err)
	}


	for _, pageItem := range list.GetContent() {
        // do something...
    }
    ```
5. 查看`client/notion_test.go`详细用法

## 注意
1. 图片直接获取 notion 的链接，并非长期有效的链接，需要后续自行处理

## TO DO
### Must
- [x] 重写 filter 调用方法
- [x] 写入到数据库可能出现的重复录入问题
- [x] 录入完成修改 notion 中的状态
- [x] filter 使用方法
- [x] notion 改造为 go extension
### Maybe
- [ ] 注释文档补全
- [ ] list/database 的 CURD 方法

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
