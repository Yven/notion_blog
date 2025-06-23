package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/Yven/notion_blog/lib"
	"github.com/Yven/notion_blog/structure"
)

type Endpoints struct {
	Id        string
	Client    Client
	Endponint EndpointType
}

type EndpointType string

const (
	TypeDatabases EndpointType = "databases"
	TypeBlocks    EndpointType = "blocks"
	TypePages     EndpointType = "pages"
)

func (t *Endpoints) GetPath(query ...any) string {
	if len(query) > 0 {
		if _, ok := query[0].(bool); ok && query[0].(bool) {
			return string(t.Endponint)
		} else {
			return string(t.Endponint) + "/" + t.Id + "/" + query[0].(string)
		}
	} else {
		return string(t.Endponint) + "/" + t.Id
	}
}

type Client interface {
	NewDb(id string) *Databases
	NewPage(id string) *Pages
	NewBlock(id string) *Blocks
	NewNotion(method lib.Method, path string, query url.Values, param any) (list *structure.List, err error)
}

type client struct {
	Host    string
	Prefix  string
	Version string
	Key     string
}

func NewClient(key string) *client {
	return &client{
		Host:    "api.notion.com",
		Prefix:  "/v1/",
		Version: "2022-06-28",
		Key:     key,
	}
}

func (c *client) NewDb(id string) *Databases {
	return &Databases{
		Endpoints: Endpoints{
			Id:        id,
			Client:    c,
			Endponint: TypeDatabases,
		},
	}
}

func (c *client) NewPage(id string) *Pages {
	return &Pages{
		Endpoints: Endpoints{
			Id:        id,
			Client:    c,
			Endponint: TypePages,
		},
	}
}

func (c *client) NewBlock(id string) *Blocks {
	return &Blocks{
		Endpoints: Endpoints{
			Id:        id,
			Client:    c,
			Endponint: TypeBlocks,
		},
	}
}

func (c *client) NewNotion(method lib.Method, path string, query url.Values, param any) (list *structure.List, err error) {
	// 检查参数是否可以被转换为JSON
	if _, paramErr := json.Marshal(param); paramErr != nil {
		return nil, errors.New("请求参数无法转换为JSON格式: " + paramErr.Error())
	}

	var apiURL *url.URL
	// fmt.Println("路由参数:", method, path, query, param)
	apiURL = &url.URL{
		Scheme:   "https",
		Host:     c.Host,
		Path:     c.Prefix + path,
		RawQuery: query.Encode(),
	}

	header := make(map[string]string)
	header["Notion-Version"] = c.Version
	header["Authorization"] = "Bearer " + c.Key

	body, _ := json.Marshal(param)
	// 发起请求
	res, err := lib.Request(method, apiURL, header, strings.NewReader(string(body)))
	if err != nil {
		return nil, errors.New("发送请求失败: " + fmt.Sprintf("%s", err))
	}

	// 解析返回结果为 List 结构体
	list = &structure.List{}
	err = json.Unmarshal(res, list)
	if err != nil {
		return nil, errors.New("解析返回结果失败: " + fmt.Sprintf("%s", err))
	}

	if list.HasMore {
		if method == lib.Get {
			query.Add("start_cursor", list.NextCursor)
		} else {
			// 判断 params 是否实现了 ListParams 接口
			if _, ok := param.(ListParams); !ok {
				// 获取下一页数据
				param.(ListParams).Next(list.NextCursor)
			}
		}

		nextList, err := c.NewNotion(method, path, query, param)
		if err != nil {
			return nil, err
		}
		if nextList != nil {
			// 有下一页，将下一页的内容添加到 list 中，整合到同一个结构体中
			list.AppendContent(nextList.GetContent())
		}
	}

	return list, nil
}
