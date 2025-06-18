package structure

import (
	"encoding/json"
	"log"
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
	Start    string `json:"start"`
	End      string `json:"end"`
	TimeZone string `json:"time_zone"`
}

type Object struct {
	Object         ObjectType `json:"object"`
	Id             string     `json:"id"`
	CreatedTime    string     `json:"created_time"`
	LastEditedTime string     `json:"last_edited_time"`
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
