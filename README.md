<div align="center">

# Notion blog API
调用Notion API获取内容，导入到Typecho的文章数据库中，或导出页面为`.md`文件

</div>

## 特性
- [x] 调用Notion API获取页面内容
- [x] 写入到Typecho数据库中
- [x] 输出页面为Markdown
- [x] 使用Docker运行
- [x] 使用命令行调用
- [ ] 输出页面为HTML

## 安装

## 使用
### 1.复制并修改配置文件
```shell
cp env.example .env
```

1. 参考[官方文档](https://developers.notion.com/docs/create-a-notion-integration)创建integration并获取Token，然后填入key中
2. 在你的文章Database中点击share然后copy link获取链接中的database_id，如此形式：`https://www.notion.so/{name}/{database_id}?v={view_id}`
3. 填写你的Typecho数据库配置

```
# 填入Notion Integration Token
NOTION_KEY=
# 填入database id
NOTION_DB_ID=

# 以下为Typecho数据库的配置
BLOG_DB_HOST=
BLOG_DB_USER=
BLOG_DB_PASSWORD=
BLOG_DB_NAME=
BLOG_DB_CHARSET=
```

### 2.运行


**使用Docker构建项目**

```shell
docker build . --tag notion_blog 

docker run -d --name notion_blog notion_blog
```

## TO DO
### Must
- [x] 重写 filter 调用方法
- [x] 写入到数据库可能出现的重复录入问题
- [ ] 录入完成修改 notion 中的状态
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
