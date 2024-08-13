<div align="center">

# Notion Blog
[![Build Status](https://img.shields.io/badge/notoin_blog-1.0-69cafd)](https://github.com/Yven/notion_blog)

调用Notion API获取内容，导入到Typecho的文章数据库中，或导出页面为`.md`文件

</div>

## 特性
- [x] 调用Notion API获取页面内容
- [x] 写入到Typecho数据库中
- [x] 输出页面为Markdown
- [x] 使用Docker运行
- [x] 使用命令行调用
- [ ] 输出页面为HTML
- [ ] notion 库

## 使用
### 1.配置参数

1. 复制 `.env` 文件
    ```shell
    cp env.example .env
    ```
2. 参考[官方文档](https://developers.notion.com/docs/create-a-notion-integration)创建 integration 并获取 Token，然后填入下面的 `NOTION_KEY` 中
3. 在你的文章 Database 中点击 share 然后 copy link 获取链接中的 `database_id`，如此形式：`https://www.notion.so/{name}/{database_id}?v={view_id}`
    ```
    # 填入Notion Integration Token
    NOTION_KEY=
    # 填入database id
    NOTION_DB_ID=
    ```
4. 填写你的 Typecho 数据库配置（如果需要，或导出为 md 文件）
    ```
    # 以下为Typecho数据库的配置
    BLOG_DB_HOST=
    BLOG_DB_USER=
    BLOG_DB_PASSWORD=
    BLOG_DB_NAME=
    BLOG_DB_CHARSET=
    ```
5. 默认将在每天早上 09:41 运行一次，可以在 `./crontab` 文件中添加或修改定时任务

### 2.运行

**使用 Docker**

```shell
docker build . --tag notion_blog 

docker run -d --name notion_blog notion_blog
```

**使用 Docker Compose**

```shell
# 如果要连接到其他 docker-compose 项目中的数据库，注意对应修改 network 名称
docker-compose up -d
```

**直接运行**

```shell
# 编译
go build main.go runner
# 添加权限
chmod +X ./runner

# 参数：
# -l 日志目录
# -c 配置文件
# 运行
./runner -l ./logs -c .env
```

## 开发

1. 在 `./plugin` 目录中新建子目录
2. 实现 `./notion/adapter.go` 中定义的接口 `ListWriter`
3. 在 `main.go` 中调用


## TO DO
### Must
- [x] 重写 filter 调用方法
- [x] 写入到数据库可能出现的重复录入问题
- [x] 录入完成修改 notion 中的状态
- [ ] request 方法，是否需要创建一个 struct，避免每次请求都需要读取配置
- [ ] 图片处理方式
- [ ] plugin 开发方式
- [ ] 输出到文件，改为一个文章一个文件
- [ ] 没有日志就不创建日志输出文件
- [ ] filter 使用方法
- [ ] notion 改造为 go extension
### Maybe
- [ ] 重写 Log
- [ ] 注释文档补全
- [ ] 配置文件
- [ ] list/database 的 CURD 方法
- [ ] Debug 模式输出参考 json 结构

## License
The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
