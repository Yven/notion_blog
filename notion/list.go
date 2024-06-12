package notion

import (
	"encoding/json"
	"io"
	"log"
	"strings"

	"notion_blog/filter"
	"notion_blog/lib/request"
)

type ListType string

const (
	DatabaseOrPage ListType = "page_or_database"
	Block          ListType = "block"
	PageList       ListType = "page"
)

// 不同的 ListType 返回对应的请求方法与路径
func (lt ListType) GetPath(id string) string {
	switch lt {
	case DatabaseOrPage:
		return "databases/" + id + "/query"
	case Block:
		return "blocks/" + id + "/children?page_size=100"
	case PageList:
		return "pages/" + id
	}

	return ""
}

type List struct {
	Object         string      `json:"object"`
	Results        []*Object   `json:"results"`
	NextCursor     string      `json:"next_cursor"`
	HasMore        bool        `json:"has_more"`
	Type           ListType    `json:"type"`
	PageOrDatabase interface{} `json:"page_or_database"`
	RequestId      string      `json:"request_id"`

	Property *Object

	OutputType OutputType
}

func NewNotionDb(id string, condition filter.Filter) *List {
	return NewNotion(request.Post, DatabaseOrPage, id, condition)
}
func NewNotionPage(id string, condition filter.Filter) *List {
	return NewNotion(request.Get, PageList, id, condition)
}
func NewNotionBlockChild(id string, condition filter.Filter) *List {
	return NewNotion(request.Get, Block, id, condition)
}

func NewNotion(method request.Method, listType ListType, id string, conditions filter.Filter) *List {
	// 构建请求 json
	var jsonData io.Reader
	if conditions == nil {
		jsonData = nil
	} else {
		jsonData = strings.NewReader(conditions.String())
	}

	// 通过需要获取的 list 类型和 id 获取请求方法和路径
	path := listType.GetPath(id)
	// fmt.Printf("Request: %s %s\n", method, path)
	// 发起请求
	res, err := request.Notion(method, path, jsonData)
	if err != nil {
		log.Printf("发送请求失败：%s", err)
		return nil
	}

	// 解析返回结果为 List 结构体
	prevList := NewList(res)

	// 获取下一页数据
	list := prevList.FetchNext(id)
	if list != nil {
		// 有下一页，将下一页的内容添加到 list 中，整合到同一个结构体中
		list.AppendContent(prevList.GetContent())
	} else {
		// 没有下一页
		list = prevList
	}

	// 如果请求的类型是一个 Block，则需要其获取所有子结构
	if listType == Block {
		for _, obj := range list.GetContent() {
			obj.FetchChild()
		}
	}

	return list
}

// 格式化数据
func NewList(data []byte) *List {
	response := &List{}
	err := json.Unmarshal(data, response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

// 获取下一页数据
func (list *List) FetchNext(id string) *List {
	if list.HasMore {
		return NewNotion(request.Get, list.Type, id, filter.NewNormalFilter("start_cursor", list.NextCursor))
	}

	return nil
}

func (list *List) GetContent() []*Object {
	return list.Results
}

func (list *List) AppendContent(objList []*Object) {
	list.Results = append(list.Results, objList...)
}

func (list *List) PageUpdate(body filter.Filter) {
	if _, err := request.Notion(request.Patch, PageList.GetPath(list.Property.Id), strings.NewReader(body.String())); err != nil {
		log.Panicf("更新页面属性失败：%s", err)
	}
}
