package notion

import (
	"encoding/json"
	"log"
	"time"
)

type ObjectType string
type IdLink map[string]string

func (obj *IdLink) GetId() string {
	return (*obj)[(*obj)["type"]]
}

const (
	ObjectPage     ObjectType = "page"
	ObjectBlock    ObjectType = "block"
	ObjectDatabase ObjectType = "database"
)

type IdMap struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

type Date struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	TimeZone string    `json:"time_zone"`
}

type Object struct {
	Object         ObjectType `json:"object"`
	Id             string     `json:"id"`
	CreatedTime    time.Time  `json:"created_time"`
	LastEditedTime time.Time  `json:"last_edited_time"`
	CreatedBy      IdMap      `json:"created_by"`
	LastEditedBy   IdMap      `json:"last_edited_by"`
	Parent         IdLink     `json:"parent"`
	Archived       bool       `json:"archived"`
	InTrash        bool       `json:"in_trash"`

	PropertyObject
	BlockObject
}

func NewObject(data []byte) *Object {
	response := &Object{}
	err := json.Unmarshal(data, response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

// 获取 Block 的内容
func (obj *Object) Fetch() *List {
	list := NewNotionBlockChild(obj.Id, nil)
	list.Property = obj

	return list
}

// 递归获取所有层级的子结构内容
func (obj *Object) FetchChild() {
	if !*obj.HasChildren {
		return
	}

	child := NewNotionBlockChild(obj.Id, nil)
	obj.AppendChild(child.GetContent())

	for _, childObj := range child.GetContent() {
		childObj.FetchChild()
	}
}
