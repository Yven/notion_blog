package filter

import (
	"encoding/json"
	"log"
)

type FilterObjectType string

const (
	ObjectTypeCheckbox    FilterObjectType = "checkbox"
	ObjectTypeDate        FilterObjectType = "date"
	ObjectTypeFiles       FilterObjectType = "files"
	ObjectTypeFormula     FilterObjectType = "formula"
	ObjectTypeMultiSelect FilterObjectType = "multi_select"
	ObjectTypeNumber      FilterObjectType = "number"
	ObjectTypePeople      FilterObjectType = "people"
	ObjectTypePhoneNumber FilterObjectType = "phone_number"
	ObjectTypeRelation    FilterObjectType = "relation"
	ObjectTypeRichText    FilterObjectType = "rich_text"
	ObjectTypeSelect      FilterObjectType = "select"
	ObjectTypeStatus      FilterObjectType = "status"
	ObjectTypeTimestamp   FilterObjectType = "timestamp"
	ObjectTypeID          FilterObjectType = "ID"
)

type MetaDataValue map[FilterConditionType]interface{}
type MetaData map[string]interface{}

func Checkbox(name string) *MetaData {
	return ObjectCommon(ObjectTypeCheckbox, name)
}
func Date(name string) *MetaData {
	return ObjectCommon(ObjectTypeDate, name)
}
func Files(name string) *MetaData {
	return ObjectCommon(ObjectTypeFiles, name)
}
func Formula(name string) *MetaData {
	return ObjectCommon(ObjectTypeFormula, name)
}
func MultiSelect(name string) *MetaData {
	return ObjectCommon(ObjectTypeMultiSelect, name)
}
func Number(name string) *MetaData {
	return ObjectCommon(ObjectTypeNumber, name)
}
func People(name string) *MetaData {
	return ObjectCommon(ObjectTypePeople, name)
}
func PhoneNumber(name string) *MetaData {
	return ObjectCommon(ObjectTypePhoneNumber, name)
}
func Relation(name string) *MetaData {
	return ObjectCommon(ObjectTypeRelation, name)
}
func RichText(name string) *MetaData {
	return ObjectCommon(ObjectTypeRichText, name)
}
func Select(name string) *MetaData {
	return ObjectCommon(ObjectTypeSelect, name)
}
func Status(name string) *MetaData {
	return ObjectCommon(ObjectTypeStatus, name)
}
func Timestamp(name string) *MetaData {
	return ObjectCommon(ObjectTypeTimestamp, name)
}
func ID(name string) *MetaData {
	return ObjectCommon(ObjectTypeID, name)
}

func ObjectCommon(objType FilterObjectType, name string) *MetaData {
	metaData := make(MetaData, 2)
	metaData["property"] = name
	valueInside := make(MetaDataValue)
	metaData[string(objType)] = valueInside

	return &metaData
}

func NewNormalFilter(key string, val interface{}) *MetaData {
	metaData := make(MetaData)
	metaData[key] = val
	return &metaData
}

func (meta *MetaData) SetFilter() *MetaData {
	newMeta := make(MetaData)
	newMeta["filter"] = meta

	return &newMeta
}

func (meta *MetaData) String() string {
	jsonData, err := json.Marshal(meta)
	if err != nil {
		log.Fatal("序列化出错,错误原因: ", err)
	}

	return string(jsonData)
}

func (meta *MetaData) StringIndent() string {
	jsonData, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		log.Fatal("序列化出错,错误原因: ", err)
	}

	return string(jsonData)
}
